package cli

import "github.com/spf13/cobra"

type AppCli struct {
	basename string
	name string
	description string
	option CliOption
	runFunc RunFunc

	commands []*Command

	args cobra.PositionalArgs
	cmd *cobra.Command
}

type Option func(*AppCli) 

func WithOptions(opt CliOption) Option {
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



func (ac *AppCli) NewAppCli(basename, name string, opts ...Option) *AppCli {
	a := &AppCli{
		basename: basename,
		name: name,
	}

	for _, o := range opts {
		o(a)
	}

	a.buildCommand()

	return a
}

func (ac *AppCli) buildCommand()  {
	cmd := &cobra.Command{
		Use: ac.basename,
		Short: ac.name,
		Long: ac.description,
		Args: ac.args,
	}


	ac.cmd = cmd
}


func (ac *AppCli) Run() {
	if err := ac.cmd.Execute();  err != nil {
		panic(err)
	}
}

