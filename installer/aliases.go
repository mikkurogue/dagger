package installer

import (
	"dagger/util"

	"github.com/charmbracelet/huh/spinner"
	"github.com/fatih/color"
)

func Aliases(aliases []string, current_os string) {
	for _, alias := range aliases {
		switch alias {
		case "git-purge":
			_ = spinner.New().Title("Setting git-purge alias").Action(func() {
				if current_os == "windows" {
					color.Red("can not set git-purge alias on this operating system \n")
					return
				}
				util.AgnosticConfigUpdater("\n# Added by dagger\nalias git-purge=\"git fetch -p && git branch --merged | grep -v '*' | grep -v 'master' | xargs git branch -d\"")
			}).TitleStyle(util.TITLE_STYLE).Run()
		case "skip":
			continue
		}
	}
}
