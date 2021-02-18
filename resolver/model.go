package resolver

import (
	"golang.org/x/sync/errgroup"
)

type Resolver struct {
	args []string

	Command Command
	Count   int
	Func    Generate
}

func New(args []string) *Resolver {
	return &Resolver{args: args}
}

func (r *Resolver) Resolve() (err error) {
	var eg errgroup.Group

	eg.Go(func() error {
		return r.setCommand()
	})

	eg.Go(func() error {
		return r.setCount()
	})

	err = eg.Wait()
	if err != nil {
		return
	}

	err = r.setEntity()
	if err != nil {
		return
	}

	//Мы бежим по списку агрументов
	// Первым всегда будет команда
	// Далее всегда идет кол-во
	// Далее идут модификаторы (random, unique, shor, only-digit и т.д.
	// Перебираем модификаторы, до тех пор, пока не встретим слово ентити (words/word, uuids/uuid и т.д.)
	// потом смотрим слова дальше (там будет либо order by, либо into {fileName}
	// по ентити определяем функцию, добавляем ей список модификаторов
	// (нужно продумать, как они будут применяться)
	// прокидываем получившуюся функцию в команду и сохраняем в структуре

	return nil
}

func (r *Resolver) Do() (err error) {
	return r.Command(r.Func)
}
