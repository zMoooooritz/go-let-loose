package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/zMoooooritz/go-let-loose/internal/util"
	"github.com/zMoooooritz/go-let-loose/pkg/config"
	"github.com/zMoooooritz/go-let-loose/pkg/rcon"
)

var (
	logsCmd = "showlog"

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
		"banprofanity",
		"unbanprofanity",
	}

	unindexedListCommands = []string{
		logsCmd,
	}
)

type inputType int

const (
	inputHostname inputType = iota
	inputPort
	inputPassword
	mainCLI
)

type model struct {
	textInput textinput.Model
	viewport  viewport.Model
	history   []string
	step      inputType
	config    rcon.ServerConfig
	rcon      *rcon.Rcon
	connected bool
	width     int
	height    int
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "Enter host..."
	ti.Focus()
	ti.Prompt = " > "

	ti.PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	ti.TextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("227"))

	vp := viewport.New(0, 0)
	vp.KeyMap = viewport.KeyMap{
		Up:   key.NewBinding(key.WithKeys("up")),
		Down: key.NewBinding(key.WithKeys("down")),
	}

	return model{
		textInput: ti,
		viewport:  vp,
		step:      inputHostname,
		history:   []string{},
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
			switch m.step {
			case inputHostname:
				m.config.Host = m.textInput.Value()
				m.textInput.SetValue("")
				m.textInput.Placeholder = "Enter port..."
				m.history = append(m.history, fmt.Sprintf(" Host: %s", m.config.Host))
				m.viewport.SetContent(strings.Join(m.history, "\n"))
				m.step = inputPort
			case inputPort:
				_, err := strconv.Atoi(m.textInput.Value())
				if err != nil {
					m.textInput.SetValue("")
					m.textInput.Placeholder = "Invalid port, try again..."
				} else {
					m.config.Port = m.textInput.Value()
					m.textInput.SetValue("")
					m.textInput.Placeholder = "Enter password..."
					m.textInput.EchoMode = textinput.EchoPassword
					m.history = append(m.history, fmt.Sprintf(" Port: %s", m.config.Port))
					m.viewport.SetContent(strings.Join(m.history, "\n"))
					m.step = inputPassword
				}
			case inputPassword:
				m.config.Password = m.textInput.Value()
				m.textInput.SetValue("")
				m.textInput.Placeholder = "Enter command..."
				m.textInput.EchoMode = textinput.EchoNormal
				m.history = append(m.history, fmt.Sprintf(" Pass: %s\n", strings.Repeat("*", len(m.config.Password))))
				rcn, err := rcon.NewRcon(m.config, 1)
				m.rcon = rcn
				if err != nil {
					m.history = append(m.history, "Unable to establish connection to the server\n")
					m.rcon = nil
				} else {
					m.history = append(m.history, "Successfully connected\n")
					m.connected = true
				}
				m.viewport.SetContent(strings.Join(m.history, "\n"))
				m.step = mainCLI
			case mainCLI:
				command := m.textInput.Value()
				if m.connected {
					response := executeCommand(m.rcon, command)
					m.history = append(m.history, command)
					for _, resp := range response {
						m.history = append(m.history, " "+resp)
					}
				} else {
					m.history = append(m.history, "Not connected, can't run: "+command)
				}
				m.viewport.SetContent(strings.Join(m.history, "\n"))
				m.viewport.GotoBottom()
				m.textInput.SetValue("")
			}
		case "ctrl+c", "esc":
			if m.rcon != nil {
				m.rcon.Close()
			}
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width - 2
		m.height = msg.Height - 2
		m.textInput.Width = msg.Width - 6

		m.viewport.Width = m.width
		m.viewport.Height = m.height - 3
	}

	var cmd tea.Cmd
	var cmds []tea.Cmd
	m.textInput, cmd = m.textInput.Update(msg)
	cmds = append(cmds, cmd)
	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	// historyView := strings.Join(m.history, "\n")
	historyBox := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Width(m.width).
		Height(m.height - 3). // Height for the history area
		Render(m.viewport.View())

	inputBox := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Width(m.width).
		Render(m.textInput.View())

	return fmt.Sprintf("%s\n%s", historyBox, inputBox)
}

func main() {
	p := tea.NewProgram(initialModel(),
		tea.WithAltScreen(),
	)
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
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
	err := func() error { return nil }()

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
		response = []string{fmt.Sprintln(err)}
	}

	return response
}
