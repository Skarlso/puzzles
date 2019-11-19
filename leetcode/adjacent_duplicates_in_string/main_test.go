package adjacent_duplicates_in_string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasic(t *testing.T) {
	s := removeDuplicates("abccba", 2)
	assert.Equal(t, "", s, "should have removed everything")
}

func Test2Basic(t *testing.T) {
	s := removeDuplicates("deeedbbcccbdaa", 3)
	assert.Equal(t, "", s, "should have removed everything")
}
