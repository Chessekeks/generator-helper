package resolver

import "log"

const (
	commandIndex = 0
	commandPrint = "print"
	commandSave  = "save"
)

type Command func(gen Generate) error

func (r *Resolver) setCommand() {
	switch r.args[commandIndex] {
	case commandPrint:
		r.Command = func(gen Generate) error {
			log.Println(gen())
			return nil
		}
	}
}
