package data

import (
	"fmt"
	"strconv"
)

// Declare a custom runtime type with underlying type of int32.
type Runtime int32

// Implement a MarshalJSON() method on the Runtime type so it satisfies the
// json.Marshaler interface. This should return the JSON-encoded value for the movie
// runtime it returns "<runtime> mins"
func (r Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)
	// Use the strconv.Quote() func on the string to wrap it in double quotes.
	// It should be surrounded by double quotes.
	quotedJSONValue := strconv.Quote(jsonValue)
	return []byte(quotedJSONValue), nil
}
