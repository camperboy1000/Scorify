// Code generated by ent, DO NOT EDIT.

package kothcheck

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the kothcheck type in the database.
	Label = "koth_check"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldFile holds the string denoting the file field in the database.
	FieldFile = "file"
	// FieldHost holds the string denoting the host field in the database.
	FieldHost = "host"
	// FieldTopic holds the string denoting the topic field in the database.
	FieldTopic = "topic"
	// FieldWeight holds the string denoting the weight field in the database.
	FieldWeight = "weight"
	// EdgeStatuses holds the string denoting the statuses edge name in mutations.
	EdgeStatuses = "statuses"
	// Table holds the table name of the kothcheck in the database.
	Table = "koth_checks"
	// StatusesTable is the table that holds the statuses relation/edge.
	StatusesTable = "koth_status"
	// StatusesInverseTable is the table name for the KothStatus entity.
	// It exists in this package in order to avoid circular dependency with the "kothstatus" package.
	StatusesInverseTable = "koth_status"
	// StatusesColumn is the table column denoting the statuses relation/edge.
	StatusesColumn = "check_id"
)

// Columns holds all SQL columns for kothcheck fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldName,
	FieldFile,
	FieldHost,
	FieldTopic,
	FieldWeight,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// FileValidator is a validator for the "file" field. It is called by the builders before save.
	FileValidator func(string) error
	// HostValidator is a validator for the "host" field. It is called by the builders before save.
	HostValidator func(string) error
	// TopicValidator is a validator for the "topic" field. It is called by the builders before save.
	TopicValidator func(string) error
	// WeightValidator is a validator for the "weight" field. It is called by the builders before save.
	WeightValidator func(int) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the KothCheck queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByUpdateTime orders the results by the update_time field.
func ByUpdateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateTime, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByFile orders the results by the file field.
func ByFile(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFile, opts...).ToFunc()
}

// ByHost orders the results by the host field.
func ByHost(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHost, opts...).ToFunc()
}

// ByTopic orders the results by the topic field.
func ByTopic(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTopic, opts...).ToFunc()
}

// ByWeight orders the results by the weight field.
func ByWeight(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWeight, opts...).ToFunc()
}

// ByStatusesCount orders the results by statuses count.
func ByStatusesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newStatusesStep(), opts...)
	}
}

// ByStatuses orders the results by statuses terms.
func ByStatuses(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newStatusesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newStatusesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(StatusesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, StatusesTable, StatusesColumn),
	)
}
