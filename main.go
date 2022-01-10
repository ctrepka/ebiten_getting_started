package main

import (
	"ebiten_getting_started/story"
	"fmt"
)

func main() {
	s, err := story.SequenceFromJSON("./story/storyJson/story1.json")
	if err != nil {
		fmt.Printf("%s", err)
	} else {
		s.StartSequenceTree()
	}
}
