package telex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertViChar(t *testing.T) {
	tests := []struct {
		name string
		c    viChar
		want rune
	}{
		{
			name: "ă",
			c: viChar{
				main: 'a',
				sub:  'w',
			},
			want: 'ă',
		},
		{
			name: "ầ",
			c: viChar{
				main: 'a',
				sub:  'a',
				mask: 'f',
			},
			want: 'ầ',
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := converViChar(tc.c)
			assert.Equal(t, string(tc.want), string(got))
		})
	}
}
