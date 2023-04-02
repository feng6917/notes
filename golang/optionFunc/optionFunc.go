package optionFunc

type repo struct {
	optionFirst string
}

// options options
type options func(*repo)

func NewRepo(opts ...options) *repo {
	srv := &repo{}
	for _, o := range opts {
		o(srv)
	}
	return srv
}

func withOptionFirst(t string) options {
	return func(c *repo) {
		c.optionFirst = t
	}
}
