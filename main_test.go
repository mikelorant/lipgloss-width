package main

import (
	"testing"
)

func TestRender(t *testing.T) {
	debug = false

	tests := []struct {
		name string
		text Text
		want int
	}{
		{
			name: "empty_plain",
			text: Text{str: "", bold: false},
			want: 1,
		},
		{
			name: "empty_bold",
			text: Text{str: "", bold: true},
			want: 1,
		},
		{
			name: "foo_plain",
			text: Text{str: "foo", bold: false},
			want: 4,
		},
		{
			name: "foo_bold",
			text: Text{str: "foo", bold: true},
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Width(tt.text)
			if got != tt.want {
				t.Errorf("want: %d got: %d", tt.want, got)
			}
		})
	}
}
