package cli

import "github.com/spf13/pflag"

// AppFlagSets 存放app各子命令行的FlagSet
type AppFlagSets struct {
	FlagSets map[string]*pflag.FlagSet
}

// FlagSet 返回指定名称的FlagSet
func (fs *AppFlagSets) FlagSet(name string) *pflag.FlagSet {
	if fs.FlagSets == nil {
		fs.FlagSets = make(map[string]*pflag.FlagSet)
	}
	if _, ok := fs.FlagSets[name]; !ok {
		fs.FlagSets[name] = pflag.NewFlagSet(name,pflag.ExitOnError)
	}
	return fs.FlagSets[name]
}