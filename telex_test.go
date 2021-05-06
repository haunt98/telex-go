package telex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConverText(t *testing.T) {
	tests := []struct {
		name string
		text string
		want string
	}{
		{
			name: "toans hocj",
			text: "toans hocj",
			want: "toán học",
		},
		{
			name: "vaatj lys",
			text: "vaatj lys",
			want: "vật lý",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := ConvertText(tc.text)
			assert.Equal(t, tc.want, got)
		})
	}
}
