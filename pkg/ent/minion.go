// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/scorify/scorify/pkg/ent/minion"
)

// Minion is the model entity for the Minion schema.
type Minion struct {
	config `json:"-"`
	// ID of the ent.
	// The uuid of the minion
	ID uuid.UUID `json:"id"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// The name of the minion
	Name string `json:"name"`
	// The ip of the minion
	IP string `json:"ip"`
	// The deactivation status of the minion
	Deactivated bool `json:"deactivated"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MinionQuery when eager-loading is set.
	Edges        MinionEdges `json:"edges"`
	selectValues sql.SelectValues
}

// MinionEdges holds the relations/edges for other nodes in the graph.
type MinionEdges struct {
	// Statuses holds the value of the statuses edge.
	Statuses []*Status `json:"status"`
	// KothStatuses holds the value of the kothStatuses edge.
	KothStatuses []*KothStatus `json:"kothStatuses"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// StatusesOrErr returns the Statuses value or an error if the edge
// was not loaded in eager-loading.
func (e MinionEdges) StatusesOrErr() ([]*Status, error) {
	if e.loadedTypes[0] {
		return e.Statuses, nil
	}
	return nil, &NotLoadedError{edge: "statuses"}
}

// KothStatusesOrErr returns the KothStatuses value or an error if the edge
// was not loaded in eager-loading.
func (e MinionEdges) KothStatusesOrErr() ([]*KothStatus, error) {
	if e.loadedTypes[1] {
		return e.KothStatuses, nil
	}
	return nil, &NotLoadedError{edge: "kothStatuses"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Minion) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case minion.FieldDeactivated:
			values[i] = new(sql.NullBool)
		case minion.FieldName, minion.FieldIP:
			values[i] = new(sql.NullString)
		case minion.FieldCreateTime, minion.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case minion.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Minion fields.
func (m *Minion) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case minion.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				m.ID = *value
			}
		case minion.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				m.CreateTime = value.Time
			}
		case minion.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				m.UpdateTime = value.Time
			}
		case minion.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				m.Name = value.String
			}
		case minion.FieldIP:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ip", values[i])
			} else if value.Valid {
				m.IP = value.String
			}
		case minion.FieldDeactivated:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field deactivated", values[i])
			} else if value.Valid {
				m.Deactivated = value.Bool
			}
		default:
			m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Minion.
// This includes values selected through modifiers, order, etc.
func (m *Minion) Value(name string) (ent.Value, error) {
	return m.selectValues.Get(name)
}

// QueryStatuses queries the "statuses" edge of the Minion entity.
func (m *Minion) QueryStatuses() *StatusQuery {
	return NewMinionClient(m.config).QueryStatuses(m)
}

// QueryKothStatuses queries the "kothStatuses" edge of the Minion entity.
func (m *Minion) QueryKothStatuses() *KothStatusQuery {
	return NewMinionClient(m.config).QueryKothStatuses(m)
}

// Update returns a builder for updating this Minion.
// Note that you need to call Minion.Unwrap() before calling this method if this Minion
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Minion) Update() *MinionUpdateOne {
	return NewMinionClient(m.config).UpdateOne(m)
}

// Unwrap unwraps the Minion entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Minion) Unwrap() *Minion {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Minion is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Minion) String() string {
	var builder strings.Builder
	builder.WriteString("Minion(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("create_time=")
	builder.WriteString(m.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(m.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(m.Name)
	builder.WriteString(", ")
	builder.WriteString("ip=")
	builder.WriteString(m.IP)
	builder.WriteString(", ")
	builder.WriteString("deactivated=")
	builder.WriteString(fmt.Sprintf("%v", m.Deactivated))
	builder.WriteByte(')')
	return builder.String()
}

// Minions is a parsable slice of Minion.
type Minions []*Minion
