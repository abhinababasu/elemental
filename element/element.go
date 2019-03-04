package element

import (
	"strings"
)

type Element struct {
	Symbol      string
	AtNumber    uint
	Name        string
	Real        bool // real or fake element
	Properties  string
	Description string
}

type Elements struct {
	elementList map[string]Element
}

func (this *Elements) getElementsForInnerWord(word string, charCount int, useFake bool) []Element {
	result := make([]Element, 0)
	s := string(word[0:charCount])

	val, ok := this.elementList[s]

	// If a element was found with the Symbol and it is fake then we cannot use it
	// if useFake was not passed
	if ok && !val.Real && !useFake {
		ok = false
	}

	if ok {
		newWord := string(word[charCount:])
		if len(newWord) > 0 {
			resultTemp := this.getElementsForWordWorker(newWord, useFake)
			if len(resultTemp) > 0 {
				result = append(result, val)
				result = append(result, resultTemp...)
			}
		} else {
			result = append(result, val)
		}
	}

	return result
}

var elements *Elements = newElements()

func newElements() *Elements {
	var elements *Elements = new(Elements)
	elements.elementList = elementList
	return elements
}

func GetElements() *Elements {
	return elements
}

func (this *Elements) getElementsForWordWorker(word string, useFake bool) []Element {
	result := make([]Element, 0)
	if len(word) > 0 {
		// First try with 2 char elements
		if len(word) > 1 {
			result = this.getElementsForInnerWord(word, 2, useFake)
		}

		// didn't find anything, so try with single char elements
		if len(result) == 0 {
			result = this.getElementsForInnerWord(word, 1, useFake)
		}
	}

	return result
}

func (this *Elements) GetElementsForWord(word string) []Element {
	word = strings.ToLower(word)

	// First try without any fake elements. E.g. Prokriti get Pr O Kr I Ti
	// If first pass we attempt fake elements then we get P R O K R I Ti
	// That is because when we try to generate at P it fails to generate a word for ROKRITI
	// and starts using fakes. Which succeeds.
	// But that failure makes the algo split at Pr - OKRITI which generates Pr O Kr I Ti
	result := this.getElementsForWordWorker(word, false)
	if len(result) == 0 {
		result = this.getElementsForWordWorker(word, true)
	}

	return result
}

func (this *Elements) GetAllElements() []Element {
	elements := make([]Element, len(elementList))
	i := 0
	for _, v := range elementList {
		elements[i] = v
		i++
	}

	return elements
}
