package control

type Option func(c *Context)

func System() Option {
	return func(c *Context) {
		c.namespace = "system"
	}
}

func User() Option {
	return func(c *Context) {
		c.namespace = "user"
	}
}

func SetFlags(flags uint64) Option {
	return func(c *Context) {
		c.flags = flags
	}
}
