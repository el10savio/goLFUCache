package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInitDoublyLinkedListTestSuite = []struct {
	name          string
	inputValue    []int
	expectedValue []int
	expectedError error
}{
	{""},
}

// TestInitDoublyLinkedList tests the functionality
// of the InitDoublyLinkedList() function
func TestInitDoublyLinkedList(t *testing.T) {
	for _, testCase := range testInitDoublyLinkedListTestSuite {
		t.Run(testCase.name, func(t *testing.T) {
			actualValue, actualError := InitDoublyLinkedList(testCase.inputValue)
			assert.Equal(t, testCase.expectedValue, actualValue)
			assert.Equal(t, testCase.expectedError, actualError)
		})
	}
}
