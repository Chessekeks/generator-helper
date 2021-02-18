package resolver

import (
	"errors"
	"log"
)

const (
	commandIndex = 0
	commandPrint = "print"
	commandSave  = "save"
)

type Command func(gen Generate) error

func (r *Resolver) setCommand() error {
	var err error
	switch r.args[commandIndex] {
	case commandPrint:
		r.Command = func(gen Generate) error {
			log.Println(gen())
			return nil
		}
	default:
		err = errors.New("unknown command")
	}

	return err
}
