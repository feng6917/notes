package parameterpassing

type ParameterRepo struct {
	ParameterOne string
	
}

// ParameterOptions options
type ParameterOptions func(*ParameterRepo)


func NewParameter(opts ...ParameterOptions) *ParameterRepo{
	srv := &ParameterRepo{}
	for _, o := range opts {
		o(srv)
	}
	return srv
}


func ParameterOne(t string) ParameterOptions {
	return func(c *ParameterRepo) {
		c.ParameterOne = t
	}
}