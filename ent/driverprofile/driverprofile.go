// Code generated by ent, DO NOT EDIT.

package driverprofile

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the driverprofile type in the database.
	Label = "driver_profile"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldLicenseNumber holds the string denoting the license_number field in the database.
	FieldLicenseNumber = "license_number"
	// FieldKtpNumber holds the string denoting the ktp_number field in the database.
	FieldKtpNumber = "ktp_number"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldVehicleInfo holds the string denoting the vehicle_info field in the database.
	FieldVehicleInfo = "vehicle_info"
	// FieldCurrentLat holds the string denoting the current_lat field in the database.
	FieldCurrentLat = "current_lat"
	// FieldCurrentLong holds the string denoting the current_long field in the database.
	FieldCurrentLong = "current_long"
	// FieldIsActive holds the string denoting the is_active field in the database.
	FieldIsActive = "is_active"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeTripsDriver holds the string denoting the trips_driver edge name in mutations.
	EdgeTripsDriver = "trips_driver"
	// Table holds the table name of the driverprofile in the database.
	Table = "driver_profiles"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "driver_profiles"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
	// TripsDriverTable is the table that holds the trips_driver relation/edge.
	TripsDriverTable = "trips"
	// TripsDriverInverseTable is the table name for the Trip entity.
	// It exists in this package in order to avoid circular dependency with the "trip" package.
	TripsDriverInverseTable = "trips"
	// TripsDriverColumn is the table column denoting the trips_driver relation/edge.
	TripsDriverColumn = "driver_profile_trips_driver"
)

// Columns holds all SQL columns for driverprofile fields.
var Columns = []string{
	FieldID,
	FieldLicenseNumber,
	FieldKtpNumber,
	FieldStatus,
	FieldVehicleInfo,
	FieldCurrentLat,
	FieldCurrentLong,
	FieldIsActive,
	FieldUserID,
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
	// DefaultIsActive holds the default value on creation for the "is_active" field.
	DefaultIsActive bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Status defines the type for the "status" enum field.
type Status string

// StatusPending is the default value of the Status enum.
const DefaultStatus = StatusPending

// Status values.
const (
	StatusPending  Status = "pending"
	StatusApproved Status = "approved"
	StatusReject   Status = "reject"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusPending, StatusApproved, StatusReject:
		return nil
	default:
		return fmt.Errorf("driverprofile: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the DriverProfile queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByLicenseNumber orders the results by the license_number field.
func ByLicenseNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLicenseNumber, opts...).ToFunc()
}

// ByKtpNumber orders the results by the ktp_number field.
func ByKtpNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldKtpNumber, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByVehicleInfo orders the results by the vehicle_info field.
func ByVehicleInfo(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVehicleInfo, opts...).ToFunc()
}

// ByCurrentLat orders the results by the current_lat field.
func ByCurrentLat(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCurrentLat, opts...).ToFunc()
}

// ByCurrentLong orders the results by the current_long field.
func ByCurrentLong(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCurrentLong, opts...).ToFunc()
}

// ByIsActive orders the results by the is_active field.
func ByIsActive(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsActive, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}

// ByTripsDriverCount orders the results by trips_driver count.
func ByTripsDriverCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTripsDriverStep(), opts...)
	}
}

// ByTripsDriver orders the results by trips_driver terms.
func ByTripsDriver(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTripsDriverStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
	)
}
func newTripsDriverStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TripsDriverInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TripsDriverTable, TripsDriverColumn),
	)
}
