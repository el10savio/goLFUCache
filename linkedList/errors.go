package linkedlist

import "errors"

var (
	ErrHeadDoesNotExist            = errors.New("head does not exist")
	ErrTailDoesNotExist            = errors.New("tail does not exist")
	ErrRangeDoesNotExist           = errors.New("range does not exist")
	ErrElementDoesNotExist         = errors.New("element does not exist")
	ErrNextElementDoesNotExist     = errors.New("next element does not exist")
	ErrPreviousElementDoesNotExist = errors.New("previous element does not exist")
)
