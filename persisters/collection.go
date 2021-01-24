package persisters

import "github.com/codemodify/falcon/data"

type CollectionPersister interface {
	SaveCollection(*data.Collection) error
}

type CollectionLoader interface {
	LoadCollection(id string) (*data.Collection, error)
}
