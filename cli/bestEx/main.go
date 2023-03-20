package main

import (
	"github.com/hangcodebug/golunzi/cli"
	"github.com/hangcodebug/golunzi/cli/bestEx/appoptions"
)

func NewApp(basename string) *cli.AppCli {
	opts := appoptions.NewOptions()
	applilcation := cli.NewAppCli(
		basename,
		"APPCLI",
		cli.WithOptions(opts),
		cli.WithDescription("desc"),
		cli.WithRunFunc(run(opts)),
		cli.WithConfig(false),
	)
	return applilcation
}

func run(opts *appoptions.Options) cli.RunFunc {
	return func(basename string) error {
		// log.Init(opts.Log)
		// defer log.Flush()

		// cfg, err := config.CreateConfigFromOptions(opts)
		// if err != nil {
		// 	return err
		// }

		// return Run(cfg)
		return nil
	}
}

func main() {
	NewApp("test").Run()
}

