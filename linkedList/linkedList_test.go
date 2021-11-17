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
	{"SizeZero", 0, nil, errors.New("list length cannot be zero")},
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

var testInsertAtFrontTestSuite = []struct {
	name                    string
	inputLength             uint
	inputValues             []int
	expectedValues          []int
	expectedValuesInReverse []int
}{
	{"BasicFuntionality", 2, []int{10}, []int{10}, []int{10}},
	{"EmptyValues", 5, []int{}, []int{}, []int{}},
	{"SizeOne", 1, []int{10}, []int{10}, []int{10}},
	{"TwoValues", 5, []int{1, 2}, []int{2, 1}, []int{1, 2}},
	{"MultipleValues", 5, []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
}

// TestInsertAtFront tests the functionality
// of the InsertAtFront() function
func TestInsertAtFront(t *testing.T) {
	for _, testCase := range testInsertAtFrontTestSuite {
		t.Run(testCase.name, func(t *testing.T) {
			dList, _ := InitDoublyLinkedList(testCase.inputLength)
			defer dList.ClearList()

			for _, inputValue := range testCase.inputValues {
				dList.InsertAtFront(inputValue)
			}

			actualValues, actualValuesReversed := dList.List(), dList.ListReverse()

			assert.Equal(t, testCase.expectedValues, actualValues)
			assert.Equal(t, testCase.expectedValuesInReverse, actualValuesReversed)
		})
	}
}

var testInsertAtEndTestSuite = []struct {
	name                    string
	inputLength             uint
	inputValues             []int
	expectedValues          []int
	expectedValuesInReverse []int
}{
	{"BasicFuntionality", 2, []int{10}, []int{10}, []int{10}},
	{"EmptyValues", 5, []int{}, []int{}, []int{}},
	{"SizeOne", 1, []int{10}, []int{10}, []int{10}},
	{"TwoValues", 5, []int{1, 2}, []int{1, 2}, []int{2, 1}},
	{"MultipleValues", 5, []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
}

// TestInsertAtEnd tests the functionality
// of the InsertAtEnd() function
func TestInsertAtEnd(t *testing.T) {
	for _, testCase := range testInsertAtEndTestSuite {
		t.Run(testCase.name, func(t *testing.T) {
			dList, _ := InitDoublyLinkedList(testCase.inputLength)
			defer dList.ClearList()

			for _, inputValue := range testCase.inputValues {
				dList.InsertAtEnd(inputValue)
			}

			actualValues, actualValuesReversed := dList.List(), dList.ListReverse()

			assert.Equal(t, testCase.expectedValues, actualValues)
			assert.Equal(t, testCase.expectedValuesInReverse, actualValuesReversed)
		})
	}
}
