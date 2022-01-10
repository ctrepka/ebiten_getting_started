package character

import (
	"ebiten_getting_started/story"
	"ebiten_getting_started/world"
)

type NPC struct {
	Character Character
	Position  world.Position
	Dialogue  *story.Sequence
}

type Character struct {
	Attributes *Attributes
	Inventory  []*world.Item
}

type Attributes struct {
	Charisma  int
	Strength  int
	Wisdom    int
	Agility   int
	Speed     int
	Integrity int
	Guile     int
	Ingenuity int
	Luck      int
}
