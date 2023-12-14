package appoptions

import (
	"github.com/chenhangch/golunzi/cli"
	"github.com/chenhangch/golunzi/cli/bestEx/exoptions"
)

type Options struct {
	MySQLOptions *exoptions.MySQLOptions `json:"mysql"    mapstructure:"mysql"`
}

func NewOptions() *Options {
	o := Options{
		MySQLOptions: exoptions.NewMySQLOptions(),
	}

	return &o
}

func (o *Options) Flags() (fss cli.AppFlagSets) {
	o.MySQLOptions.AddFlags(fss.FlagSet("mysql"))

	return fss
}
