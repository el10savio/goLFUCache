package linkedlist

import "errors"

var (
	errHeadDoesNotExist            = errors.New("head does not exist")
	errTailDoesNotExist            = errors.New("tail does not exist")
	errRangeDoesNotExist           = errors.New("range does not exist")
	errElementDoesNotExist         = errors.New("element does not exist")
	errNextElementDoesNotExist     = errors.New("next element does not exist")
	errPreviousElementDoesNotExist = errors.New("previous element does not exist")
)
