package dictionary

import (
	"fmt"
	"time"

	"github.com/dgraph-io/badger/v3"
)

// Dictionary ..
type Dictionary struct {
	db *badger.DB
}

// Entry is database data structure
type Entry struct {
	Word       string
	Definition string
	CreatedAt  time.Time
}

func (e Entry) String() string {
	created := e.CreatedAt.Format(time.Stamp)
	return fmt.Sprintf("%-10v\t%-50v%-6v", e.Word, e.Definition, created)
}

// New function inits db connection
func New(dir string) (*Dictionary, error) {
	options := badger.DefaultOptions(dir)

	db, err := badger.Open(options.WithLogger(nil))
	if err != nil {
		return nil, err
	}

	dict := &Dictionary{
		db: db,
	}

	return dict, nil
}

// Close stops db connection
func (d *Dictionary) Close() {
	d.db.Close()
}
