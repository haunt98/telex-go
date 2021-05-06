package telex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestViCharToRune(t *testing.T) {
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
			got := tc.c.toRune()
			assert.Equal(t, string(tc.want), string(got))
		})
	}
}

func TestViCharPlus(t *testing.T) {
	tests := []struct {
		name  string
		c     viChar
		r     rune
		wantC viChar
		want  bool
	}{
		{
			name: "ắ + w",
			c: viChar{
				main: 'a',
				sub:  'w',
				mask: 's',
			},
			r: 'w',
			wantC: viChar{
				main: 'a',
				mask: 's',
			},
			want: false,
		},
		{
			name: "ắ + s",
			c: viChar{
				main: 'a',
				sub:  'w',
				mask: 's',
			},
			r: 's',
			wantC: viChar{
				main: 'a',
				sub:  'w',
			},
			want: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.c.plus(tc.r)
			assert.Equal(t, tc.wantC, tc.c)
			assert.Equal(t, tc.want, got)
		})
	}
}
