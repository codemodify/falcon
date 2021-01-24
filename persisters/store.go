package persisters

import "github.com/codemodify/falcon/data"

type StorePersister interface {
	SaveStore(*data.Store) error
}

type StoreLoader interface {
	LoadStore() (*data.Store, error)
}
