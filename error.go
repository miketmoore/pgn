package pgn

import (
	"reflect"
	"strconv"
)

// An UnmarshalError describes a PGN value that led to an unexported
// (and therefore unwritable) struct field.
type UnmarshalError struct {
	Value string
	Type  reflect.Type
	Field reflect.StructField
}

// Error satisfies the Error interface and returns a human readable error message
func (e UnmarshalError) Error() string {
	return "pgn: cannot unmarshal value " + strconv.Quote(e.Value) + " into unexported field " + e.Field.Name + " of type " + e.Type.String()
}
