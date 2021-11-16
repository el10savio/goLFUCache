package linkedlist

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInitDoublyLinkedListTestSuite = []struct {
	name                     string
	inputLength              uint
	expectedDoublyLinkedList *DoublyLinkedList
	expectedError            error
}{
	{"BasicFunctionality", 2, &DoublyLinkedList{2, nil, nil}, nil},
	{"LargeList", 1000, &DoublyLinkedList{1000, nil, nil}, nil},
	{"ZeroSize", 0, nil, errors.New("list length cannot be zero")},
}

// TestInitDoublyLinkedList tests the functionality
// of the InitDoublyLinkedList() function
func TestInitDoublyLinkedList(t *testing.T) {
	for _, testCase := range testInitDoublyLinkedListTestSuite {
		t.Run(testCase.name, func(t *testing.T) {
			actualDoublyLinkedList, actualError := InitDoublyLinkedList(testCase.inputLength)
			assert.Equal(t, testCase.expectedDoublyLinkedList, actualDoublyLinkedList)
			assert.Equal(t, testCase.expectedError, actualError)
		})
	}
}
