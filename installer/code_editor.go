package installer

import (
	"os"
	"os/exec"

	"github.com/charmbracelet/huh/spinner"
	"github.com/fatih/color"
)

func CodeEditor(code_editor string, current_os string, curr_step int) {
	if code_editor != "skip" {

		if current_os == "windows" {
			color.Red("can not set install via brew on this operating system\n")
			return
		}

		if current_os == "linux" || code_editor == "zed" {
			color.Red("can not install zed on linux yet...\n")
			code_editor = "skip"
			return
		} else {
			curr_step += 1

			_ = spinner.New().Title("Installing text editor...").Action(func() {
				_, err := exec.Command("brew", "install", "--cask", code_editor).Output()
				if err != nil {
					color.Red("Error installing " + code_editor)
					os.Exit(1)
				}
			})
		}
	}
}
