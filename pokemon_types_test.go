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
