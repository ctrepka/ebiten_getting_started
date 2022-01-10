package main

import (
	"ebiten_getting_started/story"
	"fmt"
)

func main() {
	s, err := story.SequenceFromJSON("./story/storyJson/story1.json")
	//story.Seq_1_root.StartSequenceTree()
	if err != nil {
		fmt.Printf("%s", err)
	} else {
		s.StartSequenceTree()
	}
}
