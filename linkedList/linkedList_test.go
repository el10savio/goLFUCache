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

var testPrependTestSuite = []struct {
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

// TestPrepend tests the functionality
// of the Prepend() function
func TestPrepend(t *testing.T) {
	for _, testCase := range testPrependTestSuite {
		t.Run(testCase.name, func(t *testing.T) {
			dList, _ := InitDoublyLinkedList(testCase.inputLength)
			defer dList.ClearList()

			for _, inputValue := range testCase.inputValues {
				dList.Prepend(inputValue)
			}

			actualValues, actualValuesReversed := dList.List(), dList.ListReverse()

			assert.Equal(t, testCase.expectedValues, actualValues)
			assert.Equal(t, testCase.expectedValuesInReverse, actualValuesReversed)
		})
	}
}

var testAppendTestSuite = []struct {
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

// TestAppend tests the functionality
// of the Append() function
func TestAppend(t *testing.T) {
	for _, testCase := range testAppendTestSuite {
		t.Run(testCase.name, func(t *testing.T) {
			dList, _ := InitDoublyLinkedList(testCase.inputLength)
			defer dList.ClearList()

			for _, inputValue := range testCase.inputValues {
				dList.Append(inputValue)
			}

			actualValues, actualValuesReversed := dList.List(), dList.ListReverse()

			assert.Equal(t, testCase.expectedValues, actualValues)
			assert.Equal(t, testCase.expectedValuesInReverse, actualValuesReversed)
		})
	}
}

var testInsertBetweenTestSuite = []struct {
	name                    string
	inputLength             uint
	inputValues             []int
	inputValueToAdd         int
	inputValuePrevious      int
	inputValueNext          int
	expectedValues          []int
	expectedValuesInReverse []int
	expectedError           error
}{
	{"BasicFuntionality", 10, []int{1, 2, 3, 4, 5}, 7, 3, 4, []int{1, 2, 3, 7, 4, 5}, []int{5, 4, 7, 3, 2, 1}, nil},
	{"TwoElements", 10, []int{1, 2}, 7, 1, 2, []int{1, 7, 2}, []int{2, 7, 1}, nil},
	{"EmptyList", 10, []int{}, 7, 3, 4, []int{}, []int{}, errors.New("head does not exist")},
	{"SingleElement", 10, []int{1}, 7, 1, 4, []int{1}, []int{1}, errors.New("range does not exist")},
	{"RangeDoesNotExist", 10, []int{1, 2, 3, 4, 5}, 7, 11, 21, []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}, errors.New("range does not exist")},
	{"RangeStartDoesNotExist", 10, []int{1, 2, 3, 4, 5}, 7, 11, 4, []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}, errors.New("range does not exist")},
	{"RangeEndDoesNotExist", 10, []int{1, 2, 3, 4, 5}, 7, 3, 21, []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}, errors.New("range does not exist")},
	{"RangeValuesSameAsInsert", 10, []int{1, 2, 3, 3, 5}, 3, 3, 3, []int{1, 2, 3, 3, 3, 5}, []int{5, 3, 3, 3, 2, 1}, nil},
	{"RangeStartSameAsInsert", 10, []int{1, 2, 3, 4, 5}, 3, 3, 4, []int{1, 2, 3, 3, 4, 5}, []int{5, 4, 3, 3, 2, 1}, nil},
	{"RangeEndSameAsInsert", 10, []int{1, 2, 3, 4, 5}, 4, 3, 4, []int{1, 2, 3, 4, 4, 5}, []int{5, 4, 4, 3, 2, 1}, nil},
	{"RangeNotInBetween", 10, []int{1, 2, 3, 4, 5}, 7, 3, 5, []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}, errors.New("range does not exist")},
	{"RangeInBetweenEnd", 10, []int{1, 2, 3, 4, 5}, 7, 4, 5, []int{1, 2, 3, 4, 7, 5}, []int{5, 7, 4, 3, 2, 1}, nil},
}

// TestInsertBetween tests the functionality
// of the InsertBetween() function
func TestInsertBetween(t *testing.T) {
	for _, testCase := range testInsertBetweenTestSuite {
		t.Run(testCase.name, func(t *testing.T) {
			dList, _ := InitDoublyLinkedList(testCase.inputLength)
			defer dList.ClearList()

			for _, inputValue := range testCase.inputValues {
				dList.Append(inputValue)
			}

			actualError := dList.InsertBetween(testCase.inputValueToAdd, testCase.inputValuePrevious, testCase.inputValueNext)
			assert.Equal(t, testCase.expectedError, actualError)

			actualValues, actualValuesReversed := dList.List(), dList.ListReverse()

			assert.Equal(t, testCase.expectedValues, actualValues)
			assert.Equal(t, testCase.expectedValuesInReverse, actualValuesReversed)
		})
	}
}

var testGetHeadTestSuite = []struct {
	name          string
	inputLength   uint
	inputValues   []int
	expectedHead  int
	expectedError error
}{
	{"BasicFuntionality", 2, []int{10}, 10, nil},
	{"EmptyList", 2, []int{}, 0, errors.New("head does not exist")},
	{"TwoValues", 2, []int{1, 2}, 1, nil},
	{"MultipleValues", 5, []int{1, 2, 3, 4, 5}, 1, nil},
}

// TestGetHead tests the functionality
// of the GetHead() function
func TestGetHead(t *testing.T) {
	for _, testCase := range testGetHeadTestSuite {
		t.Run(testCase.name, func(t *testing.T) {
			dList, _ := InitDoublyLinkedList(testCase.inputLength)
			defer dList.ClearList()

			for _, inputValue := range testCase.inputValues {
				dList.Append(inputValue)
			}

			actualHead, actualError := dList.GetHead()

			assert.Equal(t, testCase.expectedHead, actualHead)
			assert.Equal(t, testCase.expectedError, actualError)
		})
	}
}

var testGetTailTestSuite = []struct {
	name          string
	inputLength   uint
	inputValues   []int
	expectedTail  int
	expectedError error
}{
	{"BasicFuntionality", 2, []int{10}, 10, nil},
	{"EmptyList", 2, []int{}, 0, errors.New("tail does not exist")},
	{"TwoValues", 2, []int{1, 2}, 2, nil},
	{"MultipleValues", 5, []int{1, 2, 3, 4, 5}, 5, nil},
}

// TestGetTail tests the functionality
// of the GetTail() function
func TestGetTail(t *testing.T) {
	for _, testCase := range testGetTailTestSuite {
		t.Run(testCase.name, func(t *testing.T) {
			dList, _ := InitDoublyLinkedList(testCase.inputLength)
			defer dList.ClearList()

			for _, inputValue := range testCase.inputValues {
				dList.Append(inputValue)
			}

			actualTail, actualError := dList.GetTail()

			assert.Equal(t, testCase.expectedTail, actualTail)
			assert.Equal(t, testCase.expectedError, actualError)
		})
	}
}

var testRemoveHeadTestSuite = []struct {
	name                    string
	inputLength             uint
	inputValues             []int
	expectedValues          []int
	expectedValuesInReverse []int
	expectedError           error
}{
	{"BasicFuntionality", 2, []int{10}, []int{}, []int{}, nil},
	{"EmptyValues", 5, []int{}, []int{}, []int{}, errors.New("head does not exist")},
	{"SizeOne", 1, []int{10}, []int{}, []int{}, nil},
	{"TwoValues", 5, []int{1, 2}, []int{2}, []int{2}, nil},
	{"MultipleValues", 5, []int{1, 2, 3, 4, 5}, []int{2, 3, 4, 5}, []int{5, 4, 3, 2}, nil},
}

// TestRemoveHead tests the functionality
// of the RemoveHead() function
func TestRemoveHead(t *testing.T) {
	for _, testCase := range testRemoveHeadTestSuite {
		t.Run(testCase.name, func(t *testing.T) {
			dList, _ := InitDoublyLinkedList(testCase.inputLength)
			defer dList.ClearList()

			for _, inputValue := range testCase.inputValues {
				dList.Append(inputValue)
			}

			actualError := dList.RemoveHead()
			assert.Equal(t, testCase.expectedError, actualError)

			actualValues, actualValuesReversed := dList.List(), dList.ListReverse()

			assert.Equal(t, testCase.expectedValues, actualValues)
			assert.Equal(t, testCase.expectedValuesInReverse, actualValuesReversed)
		})
	}
}

var testRemoveTailTestSuite = []struct {
	name                    string
	inputLength             uint
	inputValues             []int
	expectedValues          []int
	expectedValuesInReverse []int
	expectedError           error
}{
	{"BasicFuntionality", 2, []int{10}, []int{}, []int{}, nil},
	{"EmptyValues", 5, []int{}, []int{}, []int{}, errors.New("tail does not exist")},
	{"SizeOne", 1, []int{10}, []int{}, []int{}, nil},
	{"TwoValues", 5, []int{1, 2}, []int{1}, []int{1}, nil},
	{"MultipleValues", 5, []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4}, []int{4, 3, 2, 1}, nil},
}

// TestRemoveTail tests the functionality
// of the RemoveTail() function
func TestRemoveTail(t *testing.T) {
	for _, testCase := range testRemoveTailTestSuite {
		t.Run(testCase.name, func(t *testing.T) {
			dList, _ := InitDoublyLinkedList(testCase.inputLength)
			defer dList.ClearList()

			for _, inputValue := range testCase.inputValues {
				dList.Append(inputValue)
			}

			actualError := dList.RemoveTail()
			assert.Equal(t, testCase.expectedError, actualError)

			actualValues, actualValuesReversed := dList.List(), dList.ListReverse()

			assert.Equal(t, testCase.expectedValues, actualValues)
			assert.Equal(t, testCase.expectedValuesInReverse, actualValuesReversed)
		})
	}
}

var testRemoveElementTestSuite = []struct {
	name                    string
	inputLength             uint
	inputValues             []int
	inputValueToRemove      int
	expectedValues          []int
	expectedValuesInReverse []int
	expectedError           error
}{
	{"BasicFunctionality", 5, []int{1, 2, 3, 4, 5}, 3, []int{1, 2, 4, 5}, []int{5, 4, 2, 1}, nil},
	{"SingleElement", 1, []int{1}, 1, []int{}, []int{}, nil},
	{"TwoElementsValueAtStart", 2, []int{1, 2}, 1, []int{2}, []int{2}, nil},
	{"TwoElementsValueAtEnd", 2, []int{1, 2}, 2, []int{1}, []int{1}, nil},
	{"EmptyList", 10, []int{}, 3, []int{}, []int{}, errors.New("head does not exist")},
	{"ValueDoesNotExist", 5, []int{1, 2, 3, 4, 5}, 7, []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}, errors.New("element does not exist")},
	{"MultipleValuesExist", 5, []int{1, 3, 3, 4, 5}, 3, []int{1, 3, 4, 5}, []int{5, 4, 3, 1}, nil},
	{"ValueExistsAtStart", 5, []int{1, 2, 3, 4, 5}, 1, []int{2, 3, 4, 5}, []int{5, 4, 3, 2}, nil},
	{"ValueExistsAtEnd", 5, []int{1, 2, 3, 4, 5}, 5, []int{1, 2, 3, 4}, []int{4, 3, 2, 1}, nil},
}

// TestRemoveElement tests the functionality
// of the RemoveElement() function
func TestRemoveElement(t *testing.T) {
	for _, testCase := range testRemoveElementTestSuite {
		t.Run(testCase.name, func(t *testing.T) {
			dList, _ := InitDoublyLinkedList(testCase.inputLength)
			defer dList.ClearList()

			for _, inputValue := range testCase.inputValues {
				dList.Append(inputValue)
			}

			actualError := dList.RemoveElement(testCase.inputValueToRemove)
			assert.Equal(t, testCase.expectedError, actualError)

			actualValues, actualValuesReversed := dList.List(), dList.ListReverse()

			assert.Equal(t, testCase.expectedValues, actualValues)
			assert.Equal(t, testCase.expectedValuesInReverse, actualValuesReversed)
		})
	}
}

var testFindElementTestSuite = []struct {
	name              string
	inputLength       uint
	inputValues       []int
	inputValueToFind  int
	expectedNodeValue int
	expectedError     error
}{
	{"BasicFunctionality", 5, []int{1, 2, 3, 4, 5}, 3, 3, nil},
	{"SingleElement", 1, []int{1}, 1, 1, nil},
	{"TwoElements", 2, []int{1, 2}, 2, 2, nil},
	{"EmptyList", 5, []int{}, 11, 0, errors.New("head does not exist")},
	{"NodeNotPresent", 5, []int{1, 2, 3, 4, 5}, 11, 0, errors.New("element does not exist")},
}

// TestFindElement tests the functionality
// of the FindElement() function
func TestFindElement(t *testing.T) {
	for _, testCase := range testFindElementTestSuite {
		t.Run(testCase.name, func(t *testing.T) {
			dList, _ := InitDoublyLinkedList(testCase.inputLength)
			defer dList.ClearList()

			for _, inputValue := range testCase.inputValues {
				dList.Append(inputValue)
			}

			actualNode, actualError := dList.FindElement(testCase.inputValueToFind)

			if actualNode == nil {
				assert.Equal(t, testCase.expectedNodeValue, 0)
			} else {
				assert.Equal(t, testCase.expectedNodeValue, actualNode.Value)
			}

			assert.Equal(t, testCase.expectedError, actualError)
		})
	}
}

var testGetNextElementTestSuite = []struct {
	name                string
	inputLength         uint
	inputValues         []int
	inputValueToGetNext int
	expectedNodeValue   int
	expectedError       error
}{
	{"BasicFunctionality", 5, []int{1, 2, 3, 4, 5}, 3, 4, nil},
	{"SingleElement", 1, []int{1}, 1, 0, errors.New("next element does not exist")},
	{"TwoElements", 2, []int{1, 2}, 1, 2, nil},
	{"EmptyList", 5, []int{}, 11, 0, errors.New("element does not exist")},
	{"NodeNotPresent", 5, []int{1, 2, 3, 4, 5}, 11, 0, errors.New("element does not exist")},
	{"NextNodeNotPresent", 5, []int{1, 2, 3, 4, 5}, 5, 0, errors.New("next element does not exist")},
}

// TestGetNextElement tests the functionality
// of the GetNextElement() function
func TestGetNextElement(t *testing.T) {
	for _, testCase := range testGetNextElementTestSuite {
		t.Run(testCase.name, func(t *testing.T) {
			dList, _ := InitDoublyLinkedList(testCase.inputLength)
			defer dList.ClearList()

			for _, inputValue := range testCase.inputValues {
				dList.Append(inputValue)
			}

			searchNode, _ := dList.FindElement(testCase.inputValueToGetNext)
			actualNode, actualError := GetNextElement(searchNode)

			if actualNode == nil {
				assert.Equal(t, testCase.expectedNodeValue, 0)
			} else {
				assert.Equal(t, testCase.expectedNodeValue, actualNode.Value)
			}

			assert.Equal(t, testCase.expectedError, actualError)
		})
	}
}

var testGetPreviousElementTestSuite = []struct {
	name                    string
	inputLength             uint
	inputValues             []int
	inputValueToGetPrevious int
	expectedNodeValue       int
	expectedError           error
}{
	{"BasicFunctionality", 5, []int{1, 2, 3, 4, 5}, 3, 2, nil},
	{"SingleElement", 1, []int{1}, 1, 0, errors.New("previous element does not exist")},
	{"TwoElements", 2, []int{1, 2}, 2, 1, nil},
	{"EmptyList", 5, []int{}, 11, 0, errors.New("element does not exist")},
	{"NodeNotPresent", 5, []int{1, 2, 3, 4, 5}, 11, 0, errors.New("element does not exist")},
	{"PreviousNodeNotPresent", 5, []int{1, 2, 3, 4, 5}, 1, 0, errors.New("previous element does not exist")},
}

// TestGetPreviousElement tests the functionality
// of the GetPreviousElement() function
func TestGetPreviousElement(t *testing.T) {
	for _, testCase := range testGetPreviousElementTestSuite {
		t.Run(testCase.name, func(t *testing.T) {
			dList, _ := InitDoublyLinkedList(testCase.inputLength)
			defer dList.ClearList()

			for _, inputValue := range testCase.inputValues {
				dList.Append(inputValue)
			}

			searchNode, _ := dList.FindElement(testCase.inputValueToGetPrevious)
			actualNode, actualError := GetPreviousElement(searchNode)

			if actualNode == nil {
				assert.Equal(t, testCase.expectedNodeValue, 0)
			} else {
				assert.Equal(t, testCase.expectedNodeValue, actualNode.Value)
			}

			assert.Equal(t, testCase.expectedError, actualError)
		})
	}
}
