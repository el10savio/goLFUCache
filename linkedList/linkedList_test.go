package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
