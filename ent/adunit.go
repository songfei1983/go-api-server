// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/songfei1983/go-api-server/ent/adunit"
)

// Adunit is the model entity for the Adunit schema.
type Adunit struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Adunit) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // name
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Adunit fields.
func (a *Adunit) assignValues(values ...interface{}) error {
	if m, n := len(values), len(adunit.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	a.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[0])
	} else if value.Valid {
		a.Name = value.String
	}
	return nil
}

// Update returns a builder for updating this Adunit.
// Note that, you need to call Adunit.Unwrap() before calling this method, if this Adunit
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Adunit) Update() *AdunitUpdateOne {
	return (&AdunitClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (a *Adunit) Unwrap() *Adunit {
	tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Adunit is not a transactional entity")
	}
	a.config.driver = tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Adunit) String() string {
	var builder strings.Builder
	builder.WriteString("Adunit(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	builder.WriteString(", name=")
	builder.WriteString(a.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Adunits is a parsable slice of Adunit.
type Adunits []*Adunit

func (a Adunits) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
