package story

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Event struct {
	Title         string   `json:"title"`
	Description   string   `json:"description"`
	EventTriggers []*Event `json:"eventTriggers"`
}
type Sequence struct {
	Label         string         `json:"label"`
	Prompt        string         `json:"prompt"`
	Choices       []*Sequence    `json:"choices"`
	Requirements  []*Requirement `json:"requirements"`
	EventTriggers []*Event       `json:"eventTriggers"`
}
type Requirement struct {
	Events      []*Event `json:"events"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
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

func SequenceFromJSON(path string) (s *Sequence, e error) {
	jsonFile, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer jsonFile.Close()

	byteVal, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var sequence Sequence

	json.Unmarshal(byteVal, &sequence)

	return &sequence, err
}

//TODO: Add event triggers API routing
//TODO: Trigger events on choice selection
//TODO: Add method to add a Sequence to choices
//TODO: Add method to remove Sequence from choices
