package commandlinetool

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type Command struct {
	usage string
	short string
	long  string

	options CliOptions

	command []*Command

	args    cobra.PositionalArgs // Args 个数设置
	runFunc RunCommandFunc
}

type CommandOption func(*Command)

// RunCommandFunc 相关的 action
type RunCommandFunc func(args []string) error

// AddCommand 添加子命令行
func (c *Command) AddCommand(cmd *Command) {
	c.command = append(c.command, cmd)
}

func (c *Command) AddCommands(cmds ...*Command) {
	c.command = append(c.command, cmds...)
}

// BuildCobraCommand 构建命令行及其子命令行
func (c *Command) BuildCobraCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   c.usage,
		Short: c.short,
		Long:  c.long,
		Args:  c.args,
		// Run: c.runFunc,
	}

	cmd.SetOut(os.Stdout)
	// 判断该命令行是否存在子命令行
	if len(c.command) > 0 {
		for _, command := range c.command {
			cmd.AddCommand(command.BuildCobraCommand())
		}

	}

	if c.runFunc != nil {
		cmd.Run = c.runCommand
	}

	if c.options != nil {
		for _, f := range c.options.Flags().FlagSets {
			cmd.Flags().AddFlagSet(f)
		}
	}

	addHelpCommandFlag(c.usage, cmd.Flags())
	return cmd
}

func (c *Command) runCommand(cmd *cobra.Command, args []string) {
	if c.runFunc != nil {
		if err := c.runFunc(args); err != nil {
			fmt.Printf("%v %v\n", color.RedString("Error:"), err)
			os.Exit(1)
		}
	}
}
