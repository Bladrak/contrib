// Code generated by entc, DO NOT EDIT.

package validmessage

const (
	// Label holds the string label denoting the validmessage type in the database.
	Label = "valid_message"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldTs holds the string denoting the ts field in the database.
	FieldTs = "ts"
	// FieldUUID holds the string denoting the uuid field in the database.
	FieldUUID = "uuid"
	// Table holds the table name of the validmessage in the database.
	Table = "valid_messages"
)

// Columns holds all SQL columns for validmessage fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldTs,
	FieldUUID,
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
