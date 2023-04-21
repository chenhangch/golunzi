package cli

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/prometheus/common/log"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type AppCli struct {
	basename    string
	name        string
	description string
	option      CliOptions
	runFunc     RunFunc

	noConfig  bool
	noVersion bool

	commands []*Command

	args cobra.PositionalArgs
	cmd  *cobra.Command
}

type Option func(*AppCli)

func WithOptions(opt CliOptions) Option {
	return func(ac *AppCli) {
		ac.option = opt
	}
}

type RunFunc func(basename string) error

func WithRunFunc(run RunFunc) Option {
	return func(ac *AppCli) {
		ac.runFunc = run
	}
}

func WithDescription(desc string) Option {
	return func(ac *AppCli) {
		ac.description = desc
	}
}

func WithConfig(noConfig bool) Option {
	return func(ac *AppCli) {
		ac.noConfig = noConfig
	}
}

func NewAppCli(basename, name string, opts ...Option) *AppCli {
	a := &AppCli{
		basename: basename,
		name:     name,
	}

	for _, o := range opts {
		o(a)
	}

	a.buildCommand()

	return a
}

func (ac *AppCli) buildCommand() {
	cmd := &cobra.Command{
		Use:   ac.basename,
		Short: ac.name,
		Long:  ac.description,
		Args:  ac.args,
	}
	cmd.SetOut(os.Stdout)
	cmd.SetErr(os.Stderr)
	// 设置命令行翻译器
	cmd.Flags().SetNormalizeFunc(wordSepNormalizeFunc)

	if len(ac.commands) > 0 {
		for _, command := range ac.commands {
			cmd.AddCommand(command.BuildCobraCommand())
		}
		// 设置help命令行
		cmd.SetHelpCommand(helpCommand(ac.basename))
	}

	if ac.runFunc != nil {
		cmd.RunE = ac.runCommand
	}

	// 构建命令行参数解析
	var appFlagSets AppFlagSets
	if ac.option != nil {
		appFlagSets = ac.option.Flags()
		fs := cmd.Flags()
		for _, f := range appFlagSets.FlagSets {
			fs.AddFlagSet(f)
		}
	}

	//TODO: cmd 开启 version 版本信息
	if !ac.noVersion {

	}

	// 命令行 config 命令是否开启
	if !ac.noConfig {
		addConfigFlag(ac.basename, appFlagSets.FlagSet("global"))
	}

	appFlagSets.FlagSet("global").BoolP("help", "h", false, fmt.Sprintf("help for %s", color.GreenString(ac.name)))
	// 将全新的全局标志集添加到cmd FlagSet
	cmd.Flags().AddFlagSet(appFlagSets.FlagSet("global"))

	ac.cmd = cmd
}

// Run 开启运行命令行程序
func (ac *AppCli) Run() {
	fmt.Println(color.YellowString(time.Now().String()) + color.BlueString("===========> app cli run"))
	if err := ac.cmd.Execute(); err != nil {
		fmt.Printf("%v %v\n", color.RedString("Error:"), err)
		os.Exit(1)
	}
	fmt.Println(color.BlueString("===========> app cli cmd.Exec finish"))
}

func (ac *AppCli) runCommand(cmd *cobra.Command, args []string) error {
	// 日志记录cmd下所有flag对应的value
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		log.Debugf("FLAG: --%s=%q", color.BlueString(f.Name), color.GreenString(f.Value.String()))
	})
	// TODO:输出app --version 的信息

	if !ac.noConfig {
		if err := viper.BindPFlags(cmd.Flags()); err != nil {
			return err
		}
		if err := viper.Unmarshal(ac.option); err != nil {
			return err
		}
	}
	// if ac.option != nil {
	// }

	if ac.runFunc != nil {
		return ac.runFunc(ac.basename)
	}
	return nil
}
