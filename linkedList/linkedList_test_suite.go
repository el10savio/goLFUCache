package linkedlist

import "errors"

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
