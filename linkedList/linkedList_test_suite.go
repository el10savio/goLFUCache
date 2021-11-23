package linkedlist

var testInitDoublyLinkedListTestSuite = []struct {
	name                     string
	expectedDoublyLinkedList *DoublyLinkedList
}{
	{"BasicFunctionality", &DoublyLinkedList{nil, nil}},
}

var testPrependTestSuite = []struct {
	name                    string
	inputValues             []int
	expectedValues          []int
	expectedValuesInReverse []int
}{
	{"BasicFuntionality", []int{10}, []int{10}, []int{10}},
	{"EmptyValues", []int{}, []int{}, []int{}},
	{"SizeOne", []int{10}, []int{10}, []int{10}},
	{"TwoValues", []int{1, 2}, []int{2, 1}, []int{1, 2}},
	{"MultipleValues", []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
}

var testAppendTestSuite = []struct {
	name                    string
	inputValues             []int
	expectedValues          []int
	expectedValuesInReverse []int
}{
	{"BasicFuntionality", []int{10}, []int{10}, []int{10}},
	{"EmptyValues", []int{}, []int{}, []int{}},
	{"SizeOne", []int{10}, []int{10}, []int{10}},
	{"TwoValues", []int{1, 2}, []int{1, 2}, []int{2, 1}},
	{"MultipleValues", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
}

var testInsertBetweenTestSuite = []struct {
	name                    string
	inputValues             []int
	inputValueToAdd         int
	inputValuePrevious      int
	inputValueNext          int
	expectedValues          []int
	expectedValuesInReverse []int
	expectedError           error
}{
	{"BasicFuntionality", []int{1, 2, 3, 4, 5}, 7, 3, 4, []int{1, 2, 3, 7, 4, 5}, []int{5, 4, 7, 3, 2, 1}, nil},
	{"TwoElements", []int{1, 2}, 7, 1, 2, []int{1, 7, 2}, []int{2, 7, 1}, nil},
	{"EmptyList", []int{}, 7, 3, 4, []int{}, []int{}, errHeadDoesNotExist},
	{"SingleElement", []int{1}, 7, 1, 4, []int{1}, []int{1}, errRangeDoesNotExist},
	{"RangeDoesNotExist", []int{1, 2, 3, 4, 5}, 7, 11, 21, []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}, errRangeDoesNotExist},
	{"RangeStartDoesNotExist", []int{1, 2, 3, 4, 5}, 7, 11, 4, []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}, errRangeDoesNotExist},
	{"RangeEndDoesNotExist", []int{1, 2, 3, 4, 5}, 7, 3, 21, []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}, errRangeDoesNotExist},
	{"RangeValuesSameAsInsert", []int{1, 2, 3, 3, 5}, 3, 3, 3, []int{1, 2, 3, 3, 3, 5}, []int{5, 3, 3, 3, 2, 1}, nil},
	{"RangeStartSameAsInsert", []int{1, 2, 3, 4, 5}, 3, 3, 4, []int{1, 2, 3, 3, 4, 5}, []int{5, 4, 3, 3, 2, 1}, nil},
	{"RangeEndSameAsInsert", []int{1, 2, 3, 4, 5}, 4, 3, 4, []int{1, 2, 3, 4, 4, 5}, []int{5, 4, 4, 3, 2, 1}, nil},
	{"RangeNotInBetween", []int{1, 2, 3, 4, 5}, 7, 3, 5, []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}, errRangeDoesNotExist},
	{"RangeInBetweenEnd", []int{1, 2, 3, 4, 5}, 7, 4, 5, []int{1, 2, 3, 4, 7, 5}, []int{5, 7, 4, 3, 2, 1}, nil},
}

var testGetHeadTestSuite = []struct {
	name          string
	inputValues   []int
	expectedHead  int
	expectedError error
}{
	{"BasicFuntionality", []int{10}, 10, nil},
	{"EmptyList", []int{}, 0, errHeadDoesNotExist},
	{"TwoValues", []int{1, 2}, 1, nil},
	{"MultipleValues", []int{1, 2, 3, 4, 5}, 1, nil},
}

var testGetTailTestSuite = []struct {
	name          string
	inputValues   []int
	expectedTail  int
	expectedError error
}{
	{"BasicFuntionality", []int{10}, 10, nil},
	{"EmptyList", []int{}, 0, errTailDoesNotExist},
	{"TwoValues", []int{1, 2}, 2, nil},
	{"MultipleValues", []int{1, 2, 3, 4, 5}, 5, nil},
}

var testRemoveHeadTestSuite = []struct {
	name                    string
	inputValues             []int
	expectedValues          []int
	expectedValuesInReverse []int
	expectedError           error
}{
	{"BasicFuntionality", []int{10}, []int{}, []int{}, nil},
	{"EmptyValues", []int{}, []int{}, []int{}, errHeadDoesNotExist},
	{"SizeOne", []int{10}, []int{}, []int{}, nil},
	{"TwoValues", []int{1, 2}, []int{2}, []int{2}, nil},
	{"MultipleValues", []int{1, 2, 3, 4, 5}, []int{2, 3, 4, 5}, []int{5, 4, 3, 2}, nil},
}

var testRemoveTailTestSuite = []struct {
	name                    string
	inputValues             []int
	expectedValues          []int
	expectedValuesInReverse []int
	expectedError           error
}{
	{"BasicFuntionality", []int{10}, []int{}, []int{}, nil},
	{"EmptyValues", []int{}, []int{}, []int{}, errTailDoesNotExist},
	{"SizeOne", []int{10}, []int{}, []int{}, nil},
	{"TwoValues", []int{1, 2}, []int{1}, []int{1}, nil},
	{"MultipleValues", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4}, []int{4, 3, 2, 1}, nil},
}

var testRemoveElementTestSuite = []struct {
	name                    string
	inputValues             []int
	inputValueToRemove      int
	expectedValues          []int
	expectedValuesInReverse []int
	expectedError           error
}{
	{"BasicFunctionality", []int{1, 2, 3, 4, 5}, 3, []int{1, 2, 4, 5}, []int{5, 4, 2, 1}, nil},
	{"SingleElement", []int{1}, 1, []int{}, []int{}, nil},
	{"TwoElementsValueAtStart", []int{1, 2}, 1, []int{2}, []int{2}, nil},
	{"TwoElementsValueAtEnd", []int{1, 2}, 2, []int{1}, []int{1}, nil},
	{"EmptyList", []int{}, 3, []int{}, []int{}, errHeadDoesNotExist},
	{"ValueDoesNotExist", []int{1, 2, 3, 4, 5}, 7, []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}, errElementDoesNotExist},
	{"MultipleValuesExist", []int{1, 3, 3, 4, 5}, 3, []int{1, 3, 4, 5}, []int{5, 4, 3, 1}, nil},
	{"ValueExistsAtStart", []int{1, 2, 3, 4, 5}, 1, []int{2, 3, 4, 5}, []int{5, 4, 3, 2}, nil},
	{"ValueExistsAtEnd", []int{1, 2, 3, 4, 5}, 5, []int{1, 2, 3, 4}, []int{4, 3, 2, 1}, nil},
}

var testFindElementTestSuite = []struct {
	name              string
	inputValues       []int
	inputValueToFind  int
	expectedNodeValue int
	expectedError     error
}{
	{"BasicFunctionality", []int{1, 2, 3, 4, 5}, 3, 3, nil},
	{"SingleElement", []int{1}, 1, 1, nil},
	{"TwoElements", []int{1, 2}, 2, 2, nil},
	{"EmptyList", []int{}, 11, 0, errHeadDoesNotExist},
	{"NodeNotPresent", []int{1, 2, 3, 4, 5}, 11, 0, errElementDoesNotExist},
}

var testGetNextElementTestSuite = []struct {
	name                string
	inputValues         []int
	inputValueToGetNext int
	expectedNodeValue   int
	expectedError       error
}{
	{"BasicFunctionality", []int{1, 2, 3, 4, 5}, 3, 4, nil},
	{"SingleElement", []int{1}, 1, 0, errNextElementDoesNotExist},
	{"TwoElements", []int{1, 2}, 1, 2, nil},
	{"EmptyList", []int{}, 11, 0, errElementDoesNotExist},
	{"NodeNotPresent", []int{1, 2, 3, 4, 5}, 11, 0, errElementDoesNotExist},
	{"NextNodeNotPresent", []int{1, 2, 3, 4, 5}, 5, 0, errNextElementDoesNotExist},
}

var testGetPreviousElementTestSuite = []struct {
	name                    string
	inputValues             []int
	inputValueToGetPrevious int
	expectedNodeValue       int
	expectedError           error
}{
	{"BasicFunctionality", []int{1, 2, 3, 4, 5}, 3, 2, nil},
	{"SingleElement", []int{1}, 1, 0, errPreviousElementDoesNotExist},
	{"TwoElements", []int{1, 2}, 2, 1, nil},
	{"EmptyList", []int{}, 11, 0, errElementDoesNotExist},
	{"NodeNotPresent", []int{1, 2, 3, 4, 5}, 11, 0, errElementDoesNotExist},
	{"PreviousNodeNotPresent", []int{1, 2, 3, 4, 5}, 1, 0, errPreviousElementDoesNotExist},
}
