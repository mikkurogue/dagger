package config

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/charmbracelet/lipgloss"
)

const (
	config_file_dir  = "~/.dagger"
	config_file_name = ".cfg"
	config_file_path = "~/.dagger/.cfg"
)

// contents of the cfg
type Config struct {
	path string

	cli_tools   *[]string
	langs       *[]string
	aliases     *[]string
	code_editor *string
}

func (c *Config) ReadConfig() {

	if _, err := os.Stat(config_file_path); err != nil {
		fmt.Println(lipgloss.NewStyle().
			Background(lipgloss.Color("orange")).
			Foreground(lipgloss.Color("#d3d3d3")).
			Italic(true).
			Padding(0, 1).
			Render("No dagger config found, creating new one"))

	}
}

func OpenConfig(c *Config) {
	config, err := os.OpenFile(c.path, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println(lipgloss.NewStyle().
			Background(lipgloss.Color("orange")).
			Foreground(lipgloss.Color("#d3d3d3")).
			Italic(true).
			Padding(0, 1).
			Render("Could not open file... double check directory otherwise nuke"))

		c.path = config_file_path
		c.aliases = nil
		c.cli_tools = nil
		c.code_editor = nil
		c.langs = nil

		Create(c)
	}

	fmt.Println(config)
}

func Create(config *Config) {
	fmt.Println("We made it to create")

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
	cfgFile := filepath.Join(daggerDir, ".cfg")
	if _, err := os.Create(cfgFile); err != nil {
		fmt.Println(lipgloss.NewStyle().
			Background(lipgloss.Color("#ff0000")).
			Foreground(lipgloss.Color("#d3d3d3")).
			Italic(true).
			Padding(0, 1).
			Render("Error creating ~/.dagger/.cfg file:" + err.Error()))
		return
	}

	// TODO Write the defaults here
}

func (c *Config) DeleteConfig() {
	// allow user to delete the config just incase it somehow corrupts
	exec.Command("rm", "-rf", c.path)
}
