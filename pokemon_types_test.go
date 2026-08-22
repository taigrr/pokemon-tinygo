package main

import "testing"

func TestPokemonWinsAgainst(t *testing.T) {
	tests := []struct {
		name     string
		pokemon  Pokemon
		opponent Pokemon
		want     bool
	}{
		{
			name:     "strength matches opponent weakness",
			pokemon:  Pokemon{strength: "fire"},
			opponent: Pokemon{weakness: "fire"},
			want:     true,
		},
		{
			name:     "strength does not match opponent weakness",
			pokemon:  Pokemon{strength: "water"},
			opponent: Pokemon{weakness: "fire"},
			want:     false,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			got := testCase.pokemon.WinsAgainst(&testCase.opponent)
			if got != testCase.want {
				t.Fatalf("WinsAgainst() = %v, want %v", got, testCase.want)
			}
		})
	}
}

func TestValidSelection(t *testing.T) {
	tests := []struct {
		name     string
		selected int
		choices  int
		want     bool
	}{
		{name: "first option", selected: 0, choices: 3, want: true},
		{name: "last option", selected: 2, choices: 3, want: true},
		{name: "negative selection", selected: -1, choices: 3, want: false},
		{name: "past end", selected: 3, choices: 3, want: false},
		{name: "no choices", selected: 0, choices: 0, want: false},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			got := ValidSelection(testCase.selected, testCase.choices)
			if got != testCase.want {
				t.Fatalf("ValidSelection(%d, %d) = %v, want %v", testCase.selected, testCase.choices, got, testCase.want)
			}
		})
	}
}
