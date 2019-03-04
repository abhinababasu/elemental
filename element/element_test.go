package element

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestLoadElement(t *testing.T) {
	e := GetElements()

	if e == nil {
		t.Error("Expected not nil elements")
	}
}

func TestAllWords(t *testing.T) {
	e := GetElements()
	f, err := os.Open("./english.txt")
	if err != nil {
		t.Error("Failed to open dictionary file")
	}

	defer f.Close()
	t.Log("Opened dict")
	foundCount := 0
	notFoundCount := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		res := e.GetElementsForWord(line)

		if len(res) > 0 {
			foundCount++
			/*fmt.Print(line, " - ")
			for _, k := range res {
				fmt.Print(k.Symbol, " ")
			}
			fmt.Println()
			*/
		} else {
			//fmt.Println("!===================", line)
			notFoundCount++
		}
	}

	fmt.Println()
	fmt.Println("Found words: ", foundCount)
	fmt.Println("Not Found words: ", notFoundCount)
}

func TestSpecificWords(t *testing.T) {

	e := GetElements()
	res := e.GetElementsForWord("Prokriti")
	if len(res) != 5 {
		t.Errorf("Expected %v got %v", 0, len(res))
	}
}

func TestAllElements(t *testing.T) {
	e := GetElements()
	res := e.GetAllElements()
	if len(res) <= 118 {
		t.Error("Count of elements ")
	}

}

func TestAllElementTags(t *testing.T) {
	elements := GetElements().GetAllElements()
	validProperties := map[string]bool{
		"adaptable":   true,
		"attractive":  true,
		"available":   true,
		"bold":        true,
		"brilliant":   true,
		"creative":    true,
		"cooperative": true,
		"colorful":    true,
		"complex":     true,
		"dependable":  true,
		"earthy":      true,
		"explosive":   true,
		"fickle":      true,
		"fair":        true,
		"generous":    true,
		"honest":      true,
		"lovable":     true,
		"modest":      true,
		"noble":       true,
		"rich":        true,
		"solid":       true,
		"simple":      true,
		"toxic":       true,
	}
	validPropWeight := map[string]bool{
		"lo": true,
		"hi": true,
	}

	for _, e := range elements {
		// TODO: Some elements properties are still not set
		if len(e.Properties) > 0 {
			props := strings.Split(e.Properties, ",")
			for _, prop := range props {
				ind := strings.Split(prop, "-")
				if len(ind) < 2 {
					t.Errorf("Failed on %v(%v) for %v", e.Name, e.AtNumber, ind)
				} else {
					if !validProperties[ind[0]] {
						t.Errorf("%v(%v) => %v is not a valid property", e.Name, prop, ind[0])
					}

					if !(validPropWeight[ind[1]]) {
						t.Errorf("%v(%v) => %v is not a valid property weight", e.Name, prop, ind[1])
					}
				}
			}
		}
		//fmt.Println(e.Properties)
	}
}
