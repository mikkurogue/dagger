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
	config_file_dir  = "/.dagger"
	config_file_name = ".cfg"
	config_file_path = "/.dagger/.cfg"
)

// contents of the cfg
type Config struct {
	cli_tools   []string
	langs       []string
	aliases     []string
	code_editor string
}

func (c *Config) ReadConfig() {

	user_home_dir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(lipgloss.NewStyle().
			Background(lipgloss.Color("#ff0000")).
			Foreground(lipgloss.Color("#d3d3d3")).
			Italic(true).
			Padding(0, 1).
			Render(err.Error()))
	}

	if _, err := os.Stat(user_home_dir + config_file_path); err != nil {
		fmt.Println(lipgloss.NewStyle().
			Background(lipgloss.Color("#ff0000")).
			Foreground(lipgloss.Color("#d3d3d3")).
			Italic(true).
			Padding(0, 1).
			Render("No dagger config found, creating new one"))

	}
}

func OpenConfig(c *Config) {
	user_home_dir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(lipgloss.NewStyle().
			Background(lipgloss.Color("#ff0000")).
			Foreground(lipgloss.Color("#d3d3d3")).
			Italic(true).
			Padding(0, 1).
			Render(err.Error()))
	}

	fmt.Println(user_home_dir + config_file_path)

	config, err := os.OpenFile(user_home_dir+config_file_path, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println(lipgloss.NewStyle().
			Background(lipgloss.Color("#ff0000")).
			Foreground(lipgloss.Color("#d3d3d3")).
			Italic(true).
			Padding(0, 1).
			Render("Could not open file... double check directory otherwise nuke"))

		c.aliases = []string{"a1", "a2"}
		c.cli_tools = []string{"tool1", "tool2"}
		c.code_editor = "default_editor"
		c.langs = []string{"lang1", "lang2"}

		Create(c)
		return
	}
	defer config.Close()

	content, err := io.ReadAll(config)
	if err != nil {
		fmt.Println(lipgloss.NewStyle().
			Background(lipgloss.Color("#ff0000")).
			Foreground(lipgloss.Color("#d3d3d3")).
			Italic(true).
			Padding(0, 1).
			Render("giga err", err.Error()))
		return
	}

	fmt.Println("config content:\n", string(content))
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

	// TODO Write the defaults here
	//
	content := fmt.Sprintf(
		"# dagger config\nAliases: %v\nCLI Tools: %v\nCode Editor: %s\nLangs: %v\n",
		config.aliases,
		config.cli_tools,
		config.code_editor,
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

	fmt.Println("Configuration file created with default values")
}

func (c *Config) DeleteConfig() {
	// allow user to delete the config just incase it somehow corrupts
	//
	user_home_dir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(lipgloss.NewStyle().
			Background(lipgloss.Color("#ff0000")).
			Foreground(lipgloss.Color("#d3d3d3")).
			Italic(true).
			Padding(0, 1).
			Render(err.Error()))
	}
	exec.Command("rm", "-rf", user_home_dir+config_file_path)
}
