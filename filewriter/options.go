package filewriter

// options .
type options struct {
	// output 日志文件路径
	output string
	// maxrolls 日志文件保留天数
	maxrolls int
}

// Path .
func Path(file string) func(o *options) {
	return func(o *options) {
		if len(file) != 0 {
			o.output = file
		}
	}
}

// MaxRolls .
func MaxRolls(days int) func(o *options) {
	return func(o *options) {
		if days != 0 {
			o.maxrolls = days
		}
	}
}
