package resolver

import (
	"errors"
	"github.com/google/uuid"
	"strings"
)

const (
	entityWord  = "word"
	entityWords = "words"
	entityUUID  = "uuid"
	entityUUIDs = "uuids"
)

var entites = []string{
	entityWord, entityWords, entityUUID, entityUUIDs,
}

type Generate func() []string

func (r *Resolver) setEntity() error {
	for _, arg := range r.args[countIndex:] {
		for _, entity := range entites {
			if strings.EqualFold(arg, entity) {
				r.Func = r.resolveGenFunction(arg)
				return nil
			}
		}
	}

	return errors.New("unknown entity")
}

func (r *Resolver) resolveGenFunction(entity string) Generate {
	switch entity {
	case entityWord, entityWords:
		return nil
	case entityUUID, entityUUIDs:
		return func() []string {
			uuids := make([]string, r.Count)
			for i := 0; i < r.Count; i++ {
				uuids[i] = uuid.New().String()
			}

			return uuids
		}
	}

	return nil
}
