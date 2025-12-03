package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"

	"github.com/zMoooooritz/go-let-loose/pkg/hll"
	"github.com/zMoooooritz/go-let-loose/pkg/logger"
	"github.com/zMoooooritz/go-let-loose/pkg/rcon"
)

const workerCount = 1

// getAPIStructNames extracts all struct type names from the API package
func getAPIStructNames(apiPackagePath string) (map[string]bool, error) {
	structNames := make(map[string]bool)

	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, apiPackagePath, nil, 0)
	if err != nil {
		return nil, err
	}

	for _, pkg := range pkgs {
		for _, file := range pkg.Files {
			ast.Inspect(file, func(n ast.Node) bool {
				if typeSpec, ok := n.(*ast.TypeSpec); ok {
					if _, isStruct := typeSpec.Type.(*ast.StructType); isStruct {
						structNames[strings.ToLower(typeSpec.Name.Name)] = true
					}
				}
				return true
			})
		}
	}

	return structNames, nil
}

func main() {
	logger.DefaultLogger()

	logger.Info("Starting go-let-loose-api...")

	var cfg rcon.ServerConfig

	flag.StringVar(&cfg.Host, "host", "", "hostname of server")
	flag.StringVar(&cfg.Port, "port", "", "port on the server")
	flag.StringVar(&cfg.Password, "password", "", "password of the rcon")
	flag.Parse()

	rcn, err := rcon.NewRcon(cfg, workerCount, rcon.WithVerification())
	if err != nil {
		logger.Fatal(err)
		os.Exit(0)
	}

	cmds, err := rcn.GetCommands()
	if err != nil {
		logger.Fatal(err)
		os.Exit(0)
	}

	cmdDetails := []hll.CommandDetails{}
	for _, c := range cmds {
		if c.ClientSupported {
			commandDetails, err := rcn.GetCommandDetails(c.ID)
			if err == nil {
				cmdDetails = append(cmdDetails, commandDetails)
			} else {
				cmdDetails = append(cmdDetails, hll.CommandDetails{
					Name:        c.ID,
					Text:        c.Name,
					Description: c.Name,
				})
			}
		} else {
			cmdDetails = append(cmdDetails, hll.CommandDetails{
				Name:        c.ID,
				Text:        c.Name,
				Description: c.Name,
			})
		}
	}

	// Get the API package path
	apiPackagePath := filepath.Join("internal", "socket", "api")
	apiStructs, err := getAPIStructNames(apiPackagePath)
	if err != nil {
		logger.Warn("Failed to parse API package: " + err.Error())
	} else {
		logger.Info("Found " + fmt.Sprintf("%d", len(apiStructs)) + " struct types in API package")

		// Check each command
		missingStructs := []string{}
		for _, cmd := range cmdDetails {
			// Convert command name to PascalCase to match Go struct naming
			structName := strings.ToLower(cmd.Name)
			if !apiStructs[structName] {
				missingStructs = append(missingStructs, structName)
			}
		}

		if len(missingStructs) > 0 {
			logger.Warn("Commands without corresponding API structs:")
			for _, missing := range missingStructs {
				logger.Warn("  - " + missing)
			}
		} else {
			logger.Info("All commands have corresponding API structs")
		}
	}

	f, err := os.Create("hll_api.json")
	if err != nil {
		logger.Fatal(err)
		os.Exit(1)
	}

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	if err := enc.Encode(cmdDetails); err != nil {
		logger.Fatal(err)
		_ = f.Close()
		os.Exit(1)
	}
	_ = f.Close()

	rcn.Close()
}
