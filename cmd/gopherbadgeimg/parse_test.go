package main

import (
	"testing"
)

func TestParseRatio(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantX   int
		wantY   int
		wantErr bool
	}{
		{name: "valid square", input: "128x128", wantX: 128, wantY: 128},
		{name: "valid rectangle", input: "246x128", wantX: 246, wantY: 128},
		{name: "asymmetric", input: "100x200", wantX: 100, wantY: 200},
		{name: "uppercase", input: "128X64", wantX: 128, wantY: 64},
		{name: "missing separator", input: "128128", wantErr: true},
		{name: "too many parts", input: "128x64x32", wantErr: true},
		{name: "empty string", input: "", wantErr: true},
		{name: "non-numeric x", input: "abcx128", wantErr: true},
		{name: "non-numeric y", input: "128xabc", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotX, gotY, err := ParseRatio(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Errorf("ParseRatio(%q) expected error, got x=%d y=%d", tt.input, gotX, gotY)
				}
				return
			}
			if err != nil {
				t.Errorf("ParseRatio(%q) unexpected error: %v", tt.input, err)
				return
			}
			if gotX != tt.wantX || gotY != tt.wantY {
				t.Errorf("ParseRatio(%q) = (%d, %d), want (%d, %d)", tt.input, gotX, gotY, tt.wantX, tt.wantY)
			}
		})
	}
}
