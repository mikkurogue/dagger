package config

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/charmbracelet/lipgloss"
)

const (
	configFileDir  = "/.dagger"
	configFileName = ".cfg"
	configFilePath = "/.dagger/.cfg"
)

// contents of the cfg
type Config struct {
	cliTools   []string
	langs       []string
	aliases     []string
	codeEditor string
}

func (c *Config) ReadConfig() {

	usrHomeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(lipgloss.NewStyle().
			Background(lipgloss.Color("#ff0000")).
			Foreground(lipgloss.Color("#d3d3d3")).
			Italic(true).
			Padding(0, 1).
			Render(err.Error()))
	}

	if _, err := os.Stat(usrHomeDir + configFilePath); err != nil {
		fmt.Println(lipgloss.NewStyle().
			Background(lipgloss.Color("#ff0000")).
			Foreground(lipgloss.Color("#d3d3d3")).
			Italic(true).
			Padding(0, 1).
			Render("No dagger config found, creating new one"))

	}
}

func OpenConfig(c *Config) {
	usrHomeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(lipgloss.NewStyle().
			Background(lipgloss.Color("#ff0000")).
			Foreground(lipgloss.Color("#d3d3d3")).
			Italic(true).
			Padding(0, 1).
			Render(err.Error()))
	}

	config, err := os.OpenFile(usrHomeDir+configFilePath, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println(lipgloss.NewStyle().
			Background(lipgloss.Color("#ff0000")).
			Foreground(lipgloss.Color("#d3d3d3")).
			Italic(true).
			Padding(0, 1).
			Render("Could not open file... double check directory otherwise nuke"))

		c.aliases = []string{}
		c.cliTools = []string{}
		c.codeEditor = ""
		c.langs = []string{}

		Create(c)
		return
	}
	defer config.Close()

	// do something with the content
	_, err = io.ReadAll(config)
	if err != nil {
		fmt.Println(lipgloss.NewStyle().
			Background(lipgloss.Color("#ff0000")).
			Foreground(lipgloss.Color("#d3d3d3")).
			Italic(true).
			Padding(0, 1).
			Render("giga err", err.Error()))
		return
	}
}

func Create(config *Config) {
	dir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(lipgloss.NewStyle().
			Background(lipgloss.Color("#ff0000")).
			Foreground(lipgloss.Color("#d3d3d3")).
			Italic(true).
			Padding(0, 1).
			Render("Can not access user home dir"))
	}

	daggerDir := filepath.Join(dir, ".dagger")

	if err := os.MkdirAll(daggerDir, 0755); err != nil {
		fmt.Println(lipgloss.NewStyle().
			Background(lipgloss.Color("#ff0000")).
			Foreground(lipgloss.Color("#d3d3d3")).
			Italic(true).
			Padding(0, 1).
			Render("Error creating ~/.dagger directory:" + err.Error()))
		return
	}

	// Create ~/.dagger/.cfg file
	cfg_file := filepath.Join(daggerDir, ".cfg")
	file, err := os.OpenFile(cfg_file, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println(lipgloss.NewStyle().
			Background(lipgloss.Color("#ff0000")).
			Foreground(lipgloss.Color("#d3d3d3")).
			Italic(true).
			Padding(0, 1).
			Render("Error creating ~/.dagger/.cfg file:" + err.Error()))
		return
	}
	defer file.Close()

	content := fmt.Sprintf(
		"# dagger config\nAliases: %v\nCLI Tools: %v\nCode Editor: %s\nLangs: %v\n",
		config.aliases,
		config.cliTools,
		config.codeEditor,
		config.langs,
	)

	if _, err := file.WriteString(content); err != nil {
		fmt.Println(lipgloss.NewStyle().
			Background(lipgloss.Color("#ff0000")).
			Foreground(lipgloss.Color("#d3d3d3")).
			Italic(true).
			Padding(0, 1).
			Render("Error writing to ~/.dagger/.cfg file:" + err.Error()))
		return
	}
}

func (c *Config) DeleteConfig() {
	// allow user to delete the config just incase it somehow corrupts
	usrHomeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(lipgloss.NewStyle().
			Background(lipgloss.Color("#ff0000")).
			Foreground(lipgloss.Color("#d3d3d3")).
			Italic(true).
			Padding(0, 1).
			Render(err.Error()))
	}
	exec.Command("rm", "-rf", usrHomeDir+configFilePath)
}

func (c *Config) UpdateConfig(
	aliases []string,
	cliTools []string,
	codeEditor string,
	langs []string) {

	c.cliTools = cliTools
	c.aliases = aliases
	c.codeEditor = codeEditor
	c.langs = langs

	// re-create the config
	Create(c)
}
