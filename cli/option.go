package cli

// CliOption 命令行接口
type CliOption interface {
	// Flags() add pflag
	Flags() (fss AppFlagSets)
	// 验证
	Validate() []error
}

type ConfigurableOption interface {
	ApplyFlags() []error
}

