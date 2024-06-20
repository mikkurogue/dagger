package installer

import (
	"os/exec"

	"github.com/charmbracelet/huh/spinner"
	"github.com/fatih/color"
)

func Aliases(aliases []string, current_os string, curr_step int) {
	for _, alias := range aliases {
		switch alias {
		case "git-purge":
			_ = spinner.New().Title("Setting git-purge alias").Action(func() {
				curr_step += 1
				if current_os == "windows" {
					color.Red("can not set git-purge alias on this operating system \n")
					return
				}

				_, err := exec.Command("/bin/zsh",
					"-c",
					"alias git-purge=\"git fetch -p && git branch --merged | grep -v '*' | grep -v 'master' | xargs git branch -d\"").
					Output()
				if err != nil {
					color.Red("Can not set git purge alias\n")
				}
			}).Run()
		case "skip":
			continue
		}
	}
}
