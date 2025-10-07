package lua

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	lua "github.com/yuin/gopher-lua"
	"github.com/zMoooooritz/go-let-loose/pkg/rcon"
)

var rconInstance *rcon.Rcon

var activePlugins = make(map[string]chan bool)

func InitLua(r *rcon.Rcon) {
	rconInstance = r

	startAllPlugins()
}

func DeinitLua() {
	for _, stop := range activePlugins {
		stop <- true
	}
	activePlugins = make(map[string]chan bool)
}

func StartPlugin(pluginName string) {
	files, err := os.ReadDir(getPluginsPath())
	if err != nil {
		fmt.Println("Error reading plugins:", err)
		return
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".lua" && plugName(file.Name()) == pluginName {
			stop := make(chan bool)

			activePlugins[pluginName] = stop

			fmt.Println("Loading plugin:", pluginName)
			pluginPath := filepath.Join(getPluginsPath(), file.Name())
			go runLuaPlugin(pluginPath, stop)
		}
	}
}

func StopPlugin(pluginName string) {
	if stop, ok := activePlugins[pluginName]; ok {
		stop <- true
		delete(activePlugins, pluginName)
	}
}

func plugName(path string) string {
	base := filepath.Base(path)
	ext := filepath.Ext(base)
	return strings.TrimSuffix(base, ext)
}

func getPluginsPath() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(filepath.Dir(filepath.Dir(filename))), "plugins")
}

func GetRconInstance() *rcon.Rcon {
	return rconInstance
}

func startAllPlugins() {
	files, err := os.ReadDir(getPluginsPath())
	if err != nil {
		fmt.Println("Error reading plugins:", err)
		return
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".lua" {
			StartPlugin(plugName(file.Name()))
		}
	}
}

func runLuaPlugin(filePath string, stop chan bool) {
	L := lua.NewState()
	defer L.Close()

	RegisterBindings(L)
	RegisterEvents(L)
	RegisterHelp(L)

	// Register the exit function before loading and running running the Lua file
	exit := make(chan bool)
	L.SetGlobal("exit", L.NewFunction(func(L *lua.LState) int {
		exit <- true
		return 0
	}))

	pluginName := plugName(filePath)

	go func() {
		err := L.DoFile(filePath)
		if err != nil {
			fmt.Println("Error loading plugin:", err)
			return
		}

		fmt.Println("Initializing plugin:", pluginName)
		initFunc := L.GetGlobal("Init")
		if initFunc.Type() == lua.LTFunction {
			if err := L.CallByParam(lua.P{
				Fn:      initFunc,
				NRet:    0,
				Protect: true,
			}); err != nil {
				fmt.Println("Error calling Init():", err)
				return
			}
		} else {
			fmt.Println("No Init() function found in Plugin", pluginName)
		}

		fmt.Println("Running plugin:", pluginName)
		runFunc := L.GetGlobal("Run")
		if runFunc.Type() == lua.LTFunction {
			if err := L.CallByParam(lua.P{
				Fn:      runFunc,
				NRet:    0,
				Protect: true,
			}); err != nil {
				fmt.Println("Error calling Run():", err)
				return
			}
		} else {
			fmt.Println("No Run() function found in Plugin", pluginName, "... exiting")
			return
		}
	}()

	for {
		select {
		case <-exit:
			fmt.Println("Stopping plugin:", pluginName)
			UnregisterEvents()
			return
		case <-stop:
			fmt.Println("Stopping plugin:", pluginName)
			runFunc := L.GetGlobal("Stop")
			if runFunc.Type() == lua.LTFunction {
				if err := L.CallByParam(lua.P{
					Fn:      runFunc,
					NRet:    0,
					Protect: true,
				}); err != nil {
					return
				}
			}
			UnregisterEvents()
			return
		}
	}
}
