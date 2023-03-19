package cli

import "github.com/spf13/pflag"

var config string

func init()  {
	pflag.StringVarP(&config,"config","c","config.yaml","set you config fild || file type is yaml")
}

