package story

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Event struct {
	Title         string
	Description   string
	EventTriggers []*Event
}
type Sequence struct {
	Label         string
	Prompt        string
	Choices       []*Sequence
	Requirements  []*Requirement
	EventTriggers []*Event
}
type Requirement struct {
	Event       []*Event
	Title       string
	Description string
}

func (node *Sequence) StartSequenceTree() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(node.Prompt)
	if len(node.Choices) > 0 {
		for i, r := range node.Choices {
			fmt.Printf("%v )\t %s \n", i+1, r.Label)
		}
		scanner.Scan()
		resp := scanner.Text()
		respAsInt, err := strconv.Atoi(resp)

		if err == nil {
			if len(node.EventTriggers) > 0 {
				for _, t := range node.EventTriggers {
					fmt.Printf("Triggered event %s", t.Title)
				}
			}
			next, err := node.SelectSequenceNode(respAsInt - 1)

			if err != nil {
				fmt.Printf("\n%s. \n", err)
				next.StartSequenceTree()
			} else {
				next.StartSequenceTree()
			}
		}
	}
}

func (node *Sequence) SelectSequenceNode(nodeIndex int) (s *Sequence, e error) {
	choiceInBounds := nodeIndex < len(node.Choices) && nodeIndex >= 0
	if choiceInBounds {
		return node.Choices[nodeIndex], nil
	} else {
		err := errors.New("error: choice index invalid, please choose again")
		return node, err
	}
}

var Seq_1_root = Sequence{
	"",
	"Choose your path:",
	[]*Sequence{
		&seq_1_r1a,
		&seq_1_r1b,
	},
	nil,
	nil,
}

var seq_1_r1a = Sequence{
	"Choose sequence r1a",
	"Okay, you chose r1a. What next?",
	[]*Sequence{
		&seq_1_r2a,
	},
	nil,
	nil,
}

var seq_1_r1b = Sequence{
	"Choose sequence r1b",
	"Okay, you chose r1b. Goodbye.",
	nil,
	nil,
	nil,
}

var seq_1_r2a = Sequence{
	"Choose sequence r2a",
	"Okay, you won!",
	nil,
	nil,
	nil,
}
