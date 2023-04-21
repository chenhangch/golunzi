package cli

// CliOptions 抽象用于从命令行读取参数的配置选项。
type CliOptions interface {
	// Flags() add pflag
	Flags() (fss AppFlagSets)
	// 验证
	Validate() []error
}

// ConfigurableOptions 抽象用于从配置文件读取参数的配置选项。
type ConfigurableOptions interface {
	ApplyFlags() []error
}

// CompleteableOptions  抽象可以完成/编译的options
type CompleteableOptions interface {
	Complete() error
}

// PrintableOptions 抽象可以打印的options
type PrintableOptions interface {
	String() string
}
