package resolver

type Resolver struct {
	args []string

	Command Command
	Count   int
	Entity  string
}

func New(args []string) *Resolver {
	return &Resolver{args: args}
}

func (r *Resolver) Resolve() (err error) {
	return nil
}

func (r *Resolver) Do() (err error) {
	return nil
}
