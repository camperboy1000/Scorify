// Code generated by ent, DO NOT EDIT.

package audit

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the audit type in the database.
	Label = "audit"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAction holds the string denoting the action field in the database.
	FieldAction = "action"
	// FieldIP holds the string denoting the ip field in the database.
	FieldIP = "ip"
	// FieldTimestamp holds the string denoting the timestamp field in the database.
	FieldTimestamp = "timestamp"
	// FieldMessage holds the string denoting the message field in the database.
	FieldMessage = "message"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the audit in the database.
	Table = "audits"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "audits"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "audit_user"
)

// Columns holds all SQL columns for audit fields.
var Columns = []string{
	FieldID,
	FieldAction,
	FieldIP,
	FieldTimestamp,
	FieldMessage,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "audits"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"audit_user",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// IPValidator is a validator for the "ip" field. It is called by the builders before save.
	IPValidator func(string) error
	// DefaultTimestamp holds the default value on creation for the "timestamp" field.
	DefaultTimestamp func() time.Time
	// MessageValidator is a validator for the "message" field. It is called by the builders before save.
	MessageValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Action defines the type for the "action" enum field.
type Action string

// Action values.
const (
	ActionAuthLogin             Action = "auth_login"
	ActionAuthLogout            Action = "auth_logout"
	ActionAuthFailedLogin       Action = "auth_failed_login"
	ActionAdminLogin            Action = "admin_login"
	ActionAdminBecome           Action = "admin_become"
	ActionUserChangePassword    Action = "user_change_password"
	ActionUserCreate            Action = "user_create"
	ActionUserUpdate            Action = "user_update"
	ActionUserDelete            Action = "user_delete"
	ActionCheckCreate           Action = "check_create"
	ActionCheckUpdate           Action = "check_update"
	ActionCheckDelete           Action = "check_delete"
	ActionCheckValidate         Action = "check_validate"
	ActionCheckConfig           Action = "check_config"
	ActionNotificationCreate    Action = "notification_create"
	ActionEngineStart           Action = "engine_start"
	ActionEngineStop            Action = "engine_stop"
	ActionInjectCreate          Action = "inject_create"
	ActionInjectUpdate          Action = "inject_update"
	ActionInjectDelete          Action = "inject_delete"
	ActionInjectSubmit          Action = "inject_submit"
	ActionInjectGrade           Action = "inject_grade"
	ActionMinionRegister        Action = "minion_register"
	ActionMinionDeactivate      Action = "minion_deactivate"
	ActionMinionActivate        Action = "minion_activate"
	ActionWipeAll               Action = "wipe_all"
	ActionWipeCheckConfigs      Action = "wipe_check_configs"
	ActionWipeInjectSubmissions Action = "wipe_inject_submissions"
	ActionWipeStatuses          Action = "wipe_statuses"
	ActionWipeScores            Action = "wipe_scores"
	ActionWipeRound             Action = "wipe_round"
	ActionWipeCache             Action = "wipe_cache"
)

func (a Action) String() string {
	return string(a)
}

// ActionValidator is a validator for the "action" field enum values. It is called by the builders before save.
func ActionValidator(a Action) error {
	switch a {
	case ActionAuthLogin, ActionAuthLogout, ActionAuthFailedLogin, ActionAdminLogin, ActionAdminBecome, ActionUserChangePassword, ActionUserCreate, ActionUserUpdate, ActionUserDelete, ActionCheckCreate, ActionCheckUpdate, ActionCheckDelete, ActionCheckValidate, ActionCheckConfig, ActionNotificationCreate, ActionEngineStart, ActionEngineStop, ActionInjectCreate, ActionInjectUpdate, ActionInjectDelete, ActionInjectSubmit, ActionInjectGrade, ActionMinionRegister, ActionMinionDeactivate, ActionMinionActivate, ActionWipeAll, ActionWipeCheckConfigs, ActionWipeInjectSubmissions, ActionWipeStatuses, ActionWipeScores, ActionWipeRound, ActionWipeCache:
		return nil
	default:
		return fmt.Errorf("audit: invalid enum value for action field: %q", a)
	}
}

// OrderOption defines the ordering options for the Audit queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByAction orders the results by the action field.
func ByAction(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAction, opts...).ToFunc()
}

// ByIP orders the results by the ip field.
func ByIP(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIP, opts...).ToFunc()
}

// ByTimestamp orders the results by the timestamp field.
func ByTimestamp(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTimestamp, opts...).ToFunc()
}

// ByMessage orders the results by the message field.
func ByMessage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMessage, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
	)
}
