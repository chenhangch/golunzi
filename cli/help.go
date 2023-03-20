package cli

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const (
	flagHelp = "help"
	flagHelpShorthand = "H"
)

func helpCommand(name string) *cobra.Command {
	return &cobra.Command{
		Use: "help [command]",
		Short: "Help about any command",
		Long: `Help providesn help for any cmmand in the applicaton.
		Simply type ` + name + `help [path to command] for full detatils`,

		Run: func(c *cobra.Command, args []string) {
			cmd, _, e := c.Root().Find(args)
			if cmd == nil || e != nil {
				c.Printf("Unknown help topic %#q\n", args)
				_ = c.Root().Usage()
			} else {
				cmd.InitDefaultHelpFlag() // make possible 'help' flag to be shown
				_ = cmd.Help()
			}
		},
	}
}

// addHelpCommandFlag 添加帮助命令行参数
func addHelpCommandFlag(usage string, fs *pflag.FlagSet)  {
	fs.BoolP(
		flagHelp,
		flagHelpShorthand,
		false,
		fmt.Sprintf("Help for the %s command.", color.GreenString(strings.Split(usage, " ")[0])),
)
}