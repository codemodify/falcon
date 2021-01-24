package persisters

import "github.com/codemodify/falcon/data"

type RecordPersister interface {
	SaveRecord(*data.Record) error
}

type RecordLoader interface {
	LoadRecord(id string) (*data.Record, error)
}
