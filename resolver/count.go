package resolver

import "strconv"

const (
	countIndex = 1
)

func (r *Resolver) setCount() error {
	var err error
	r.Count, err = strconv.Atoi(r.args[countIndex])

	return err
}
