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
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := converViChar(tc.c)
			assert.Equal(t, tc.want, got)
		})
	}
}
