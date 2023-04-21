package clog

// Formatter 自定义日志输出格式
type Formatter interface {
	// Maybe in async goroutine
	// Please write the result to buffer
	Format(entry *Entry) error
}
