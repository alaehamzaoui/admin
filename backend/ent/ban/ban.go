// Code generated by ent, DO NOT EDIT.

package ban

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the ban type in the database.
	Label = "ban"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCheck holds the string denoting the check field in the database.
	FieldCheck = "check"
	// FieldMessage holds the string denoting the message field in the database.
	FieldMessage = "message"
	// FieldValidUntil holds the string denoting the validuntil field in the database.
	FieldValidUntil = "valid_until"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// Table holds the table name of the ban in the database.
	Table = "bans"
)

// Columns holds all SQL columns for ban fields.
var Columns = []string{
	FieldID,
	FieldCheck,
	FieldMessage,
	FieldValidUntil,
	FieldCreatedAt,
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
	// DefaultCreatedAt holds the default value on creation for the "createdAt" field.
	DefaultCreatedAt func() time.Time
)

// OrderOption defines the ordering options for the Ban queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCheck orders the results by the check field.
func ByCheck(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCheck, opts...).ToFunc()
}

// ByMessage orders the results by the message field.
func ByMessage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMessage, opts...).ToFunc()
}

// ByValidUntil orders the results by the validUntil field.
func ByValidUntil(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldValidUntil, opts...).ToFunc()
}

// ByCreatedAt orders the results by the createdAt field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}
