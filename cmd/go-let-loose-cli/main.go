package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/zMoooooritz/go-let-loose/internal/util"
	"github.com/zMoooooritz/go-let-loose/pkg/config"
	"github.com/zMoooooritz/go-let-loose/pkg/logger"
	"github.com/zMoooooritz/go-let-loose/pkg/rcon"
)

var (
	// Set via ldflags when building
	Version   = ""
	CommitSHA = ""

	logsCmd = "showlog"

	otherCommands = []string{
		"clear",
	}

	normalCommands = []string{
		"help",
		"get name",
		"get slots",
		"get gamemode",
		"get gamestate",
		"get maxqueuedplayers",
		"setmaxqueuedplayers",
		"get numvipslots",
		"setnumvipslots",
		"say",
		"broadcast",
		"get map",
		"rotlist",
		"rotadd",
		"rotdel",
		"map",
		"gamelayout",
		"querymapshuffle",
		"togglemapshuffle",
		"listcurrentmapsequence",
		"playerinfo",
		"adminadd",
		"admindel",
		"vipadd",
		"vipdel",
		"message",
		"punish",
		"switchteamondeath",
		"switchteamnow",
		"kick",
		"tempban",
		"pardontempban",
		"permaban",
		"pardonpermaban",
		"get idletime",
		"get highping",
		"get teamswitchcooldown",
		"get autobalanceenabled",
		"get votekickenabled",
		"get votekickthreshold",
		"get profanity",
		"setkickidletime",
		"sethighping",
		"setteamswitchcooldown",
		"setautobalanceenabled",
		"setautobalancethreshold",
		"setvotekickenabled",
		"setvotekickthreshold",
		"resetvotekickthreshold",
	}

	indexedListCommands = []string{
		"get mapsforrotation",
		"get players",
		"get playerids",
		"get adminids",
		"get admingrups",
		"get vipids",
		"get tempbans",
		"get permabans",
		"get profanities",
		"get objectiverow_0",
		"get objectiverow_1",
		"get objectiverow_2",
		"get objectiverow_3",
		"get objectiverow_4",
		"banprofanity",
		"unbanprofanity",
	}

	unindexedListCommands = []string{
		logsCmd,
	}

	allCommands = slices.Concat(normalCommands, indexedListCommands, unindexedListCommands, otherCommands)
)

type model struct {
	textInput textinput.Model
	viewport  viewport.Model
	history   []string
	config    rcon.ServerConfig
	rcon      *rcon.Rcon
	connected bool
	width     int
	height    int
}

func initialModel(cfg rcon.ServerConfig) model {
	ti := textinput.New()
	ti.Focus()
	ti.Prompt = " > "
	ti.Placeholder = ""
	ti.ShowSuggestions = true

	ti.PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	ti.TextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("227"))

	vp := viewport.New(0, 0)
	vp.KeyMap = viewport.KeyMap{
		Up:   key.NewBinding(key.WithKeys("up")),
		Down: key.NewBinding(key.WithKeys("down")),
	}

	if cfg.Host == "" && ti.Placeholder == "" {
		ti.Placeholder = "Enter host..."
	}
	if cfg.Port == "" && ti.Placeholder == "" {
		ti.Placeholder = "Enter port..."
	}
	if cfg.Password == "" && ti.Placeholder == "" {
		ti.Placeholder = "Enter password..."
		ti.EchoMode = textinput.EchoPassword
	}

	var history []string
	var rcn *rcon.Rcon
	var connected = false
	if ti.Placeholder == "" {
		var err error
		ti.Placeholder = "Enter command..."
		rcn, err = rcon.NewRcon(cfg, 1)
		if err != nil {
			history = append(history, "Unable to establish connection to the server\n")
			rcn = nil
		} else {
			name, err := rcn.GetServerName()
			msg := "Successfully connected"
			if err == nil {
				msg += " to " + name
			}
			history = append(history, msg+"\n")
			connected = true
			ti.SetSuggestions(allCommands)
		}
		vp.SetContent(strings.Join(history, "\n"))
	}

	return model{
		textInput: ti,
		viewport:  vp,
		history:   history,
		config:    cfg,
		rcon:      rcn,
		connected: connected,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if m.config.Host == "" {
				if m.textInput.Value() != "" {
					m.config.Host = m.textInput.Value()
					m.textInput.Placeholder = "Enter port..."
					m.history = append(m.history, fmt.Sprintf(" Host: %s", m.config.Host))
				}
			} else if m.config.Port == "" {
				if m.textInput.Value() != "" {
					_, err := strconv.Atoi(m.textInput.Value())
					if err != nil {
						m.textInput.Placeholder = "Invalid port, try again..."
					} else {
						m.config.Port = m.textInput.Value()
						m.textInput.Placeholder = "Enter password..."
						m.textInput.EchoMode = textinput.EchoPassword
						m.history = append(m.history, fmt.Sprintf(" Port: %s", m.config.Port))
					}
				}
			} else if m.config.Password == "" {
				if m.textInput.Value() != "" {
					m.config.Password = m.textInput.Value()
					m.textInput.Placeholder = "Enter command..."
					m.textInput.EchoMode = textinput.EchoNormal
					m.history = append(m.history, fmt.Sprintf(" Pass: %s\n", strings.Repeat("*", len(m.config.Password))))
					rcn, err := rcon.NewRcon(m.config, 1)
					m.rcon = rcn
					if err != nil {
						m.history = append(m.history, "Unable to establish connection to the server\n")
						m.rcon = nil
					} else {
						name, err := m.rcon.GetServerName()
						msg := "Successfully connected"
						if err == nil {
							msg += " to " + name
						}
						m.history = append(m.history, msg+"\n")
						m.connected = true

						m.textInput.SetSuggestions(allCommands)
					}
				}
			} else {
				command := m.textInput.Value()

				if command == "clear" {
					m.history = []string{}
					m.viewport.SetContent(strings.Join(m.history, "\n"))
				} else {
					if !isValidCommand(command) {
						m.history = append(m.history, "Invalid command, won't run: "+command)
					} else if !m.connected {
						m.history = append(m.history, "Not connected, can't run: "+command)
					} else {
						response := executeCommand(m.rcon, command)
						m.history = append(m.history, command)
						for _, resp := range response {
							m.history = append(m.history, " "+strings.ReplaceAll(resp, "\n", "\n "))
						}
					}
				}
			}
			m.viewport.SetContent(strings.Join(m.history, "\n"))
			m.viewport.GotoBottom()
			m.textInput.SetValue("")
		case "ctrl+c", "esc":
			if m.rcon != nil {
				m.rcon.Close()
			}
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width - 2
		m.height = msg.Height - 2
		m.textInput.Width = msg.Width - 30

		m.viewport.Width = m.width - 2
		m.viewport.Height = m.height - 5
	}

	var cmd tea.Cmd
	var cmds []tea.Cmd
	m.textInput, cmd = m.textInput.Update(msg)
	cmds = append(cmds, cmd)
	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func isValidCommand(command string) bool {
	for _, cmd := range allCommands {
		if strings.HasPrefix(command, cmd) {
			return true
		}
	}
	return false
}

func (m model) View() string {
	historyBox := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Width(m.width).
		Height(m.height - 3).
		Padding(1).
		Render(m.viewport.View())

	inputBox := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Width(m.width).
		Render(m.textInput.View())

	return fmt.Sprintf("%s\n%s", historyBox, inputBox)
}

func executeCommand(rcn *rcon.Rcon, cmd string) []string {
	responseFormat := config.RF_DIRECT
	for _, listCmd := range indexedListCommands {
		if strings.HasPrefix(strings.ToLower(cmd), listCmd) {
			responseFormat = config.RF_INDEXEDLIST
		}
	}
	for _, listCmd := range unindexedListCommands {
		if strings.HasPrefix(strings.ToLower(cmd), listCmd) {
			responseFormat = config.RF_UNINDEXEDLIST
		}
	}

	response := []string{}
	var err error

	switch responseFormat {
	case config.RF_DIRECT:
		fallthrough
	case config.RF_INDEXEDLIST:
		response, err = rcn.RunCommand(cmd, responseFormat)
	case config.RF_UNINDEXEDLIST:
		if strings.HasPrefix(strings.ToLower(cmd), logsCmd) {
			cmd, _ = strings.CutPrefix(strings.ToLower(cmd), logsCmd)
			minutes := util.ToInt(strings.TrimSpace(cmd))
			response, err = rcn.GetLogs(minutes)
		}
	}

	if err != nil {
		response = []string{fmt.Sprint(err)}
	}

	return response
}

func main() {
	logger.NOPLogger()

	var cfg rcon.ServerConfig
	var version bool

	flag.StringVar(&cfg.Host, "host", "", "hostname of server")
	flag.StringVar(&cfg.Port, "port", "", "port on the server")
	flag.StringVar(&cfg.Password, "password", "", "password of the rcon")
	flag.BoolVar(&version, "version", false, "display version")
	flag.Parse()

	if version {
		if len(CommitSHA) > 7 {
			CommitSHA = CommitSHA[:7]
		}
		if Version == "" {
			Version = "(built from source)"
		}

		fmt.Printf("go-let-loose-cli %s", Version)
		if len(CommitSHA) > 0 {
			fmt.Printf(" (%s)", CommitSHA)
		}

		fmt.Println()
		os.Exit(0)
	}

	p := tea.NewProgram(initialModel(cfg),
		tea.WithAltScreen(),
	)
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
}
