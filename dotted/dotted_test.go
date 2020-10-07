package dotted_test

import (
	"testing"

	"github.com/FrancescoIlario/gocg/dotted"
	"github.com/matryer/is"
)

func TestToMap(t *testing.T) {
	// Arrange
	is := is.New(t)

	s := "a.b=c"
	e := map[string]interface{}{
		"a": map[string]interface{}{
			"b": "c",
		},
	}

	// Act
	m, err := dotted.ToMap(s)

	// Assert
	is.NoErr(err)
	is.Equal(e, m)
}
