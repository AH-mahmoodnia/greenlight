package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")

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

// This will change the Runtime and so we make it pointer receiver
func (r *Runtime) UnmarshalJSON(jsonValue []byte) error {
	// The jsonValue is in the format of "<runtime> mins"
	unquotedJSONValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	parts := strings.Split(unquotedJSONValue, " ")

	if len(parts) != 2 || parts[1] != "mins" {
		return ErrInvalidRuntimeFormat
	}

	// convert to int32
	i, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return ErrInvalidRuntimeFormat
	}
	*r = Runtime(i)
	return nil
}
