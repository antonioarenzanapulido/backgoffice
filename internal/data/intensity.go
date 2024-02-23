package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Intensity int32

var ErrInvalidIntensityFormat = errors.New("invalid intensity format")

func (i Intensity) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d / 10", i)

	quotedJSON := strconv.Quote(jsonValue)

	return []byte(quotedJSON), nil
}

func (i *Intensity) UnmarshalJSON(jsonValue []byte) error {
	unquotedJSONValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidIntensityFormat
	}

	parts := strings.Split(unquotedJSONValue, " ")

	if len(parts) != 3 || parts[1] != "/" || parts[2] != "10" {
		return ErrInvalidIntensityFormat
	}

	v, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return ErrInvalidIntensityFormat
	}

	*i = Intensity(v)
	return nil
}
