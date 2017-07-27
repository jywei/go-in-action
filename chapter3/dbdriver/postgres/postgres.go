package postgres

import (
	"database/sql"
	"database/sql/driver"
	"errors"
)

// Driver provides our implementation for the sql package
type Driver struct{}

// Open provides a connection to the database
func (dr Driver) Open(string) (driver.Conn, error) {
	return nil, errors.New("Unimplemented")
}

var d *Driver

func init() {
	d = new(Driver)
	sql.Register("postgres", d)
}
