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
	assert.Equal(t, "aa", s, "should have removed everything")
}

func Test3Basic(t *testing.T) {
	st := "mtaknxdfffffffttttttttttttttttttttttttttttttttttttttttttttttttttfffffffffffffffffffffffffffffffffffffffffffpkrzbadhyvrojrqkptqffpxlkzjvbvbppppppppppbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbpppppppprrrrrrrrrrddddddddddddddddddddddddddddddddddddddddddddddddddrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxttttttttttttttttttttttttttttttttttttttttttttttttttxxxpppppppppxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxpppppppppppppppppppppppoktyhzpvnokn"
	s := removeDuplicates(st, 50)
	assert.Equal(t, "mtaknxdpkrzbadhyvrojrqkptqffpxlkzjvbvboktyhzpvnokn", s, "should have removed everything")
}

func BenchmarkBasic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := removeDuplicates("deeedbbcccbdaa", 3)
		if s != "aa" {
			b.Fatal("s: ", s)
		}
	}
}
