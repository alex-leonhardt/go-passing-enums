package env

type env struct {
	val1 string
	val2 int
	val3 bool
}

// New creates a new config
func New() *env {
	return &env{}
}

func (c *env) Get(v string) string {
	return ""
}
func (c *env) Set(v string) bool {
	return true
}
func (c *env) Del(v string) bool {
	return true
}
