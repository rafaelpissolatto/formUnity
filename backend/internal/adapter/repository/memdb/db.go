package memdb

import "github.com/hashicorp/go-memdb"

// DB is wrapper for in-memory database connection
type DB struct {
	*memdb.MemDB
}

// NewDB creates a new in-memory database connection
func NewDB() (*DB, error) {
	// Create the DB schema (dynamic based on the Volunteer type)
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"volunteer": {
				Name: "volunteer",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
					"email": {
						Name:    "email",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Email"},
					},
					"phone": {
						Name:    "phone",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Phone"},
					},
				},
			},
		},
	}

	// Create a new data base
	db, err := memdb.NewMemDB(schema)
	if err != nil {
		return nil, err
	}

	return &DB{
		MemDB: db,
	}, nil
}
