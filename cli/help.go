package cli

import "github.com/spf13/cobra"

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

		Run: func(cmd *cobra.Command, args []string) {
			
		},
	}
}