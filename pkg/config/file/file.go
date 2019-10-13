package file

type file struct {
	val1 string
	val2 int
	val3 bool
}

// New creates a new config
func New() *file {
	return &file{}
}

func (c *file) Get(v string) string {
	return ""
}
func (c *file) Set(v string) bool {
	return true
}
func (c *file) Del(v string) bool {
	return true
}
