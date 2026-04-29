package main

// Pokemon describes a battler and its elemental matchup data.
type Pokemon struct {
	//lint:ignore U1000 used in tinygo-only rendering code
	name string
	//lint:ignore U1000 used in tinygo-only rendering code
	sprite   []uint8
	strength string
	weakness string
}

func (pokemon *Pokemon) WinsAgainst(opponent *Pokemon) bool {
	return pokemon.strength == opponent.weakness
}
