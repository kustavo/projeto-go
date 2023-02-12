package domain

import "strings"

type ErrorsList struct {
	Errs []error
}

func NewErrorsList(errs []error) *ErrorsList {
	if len(errs) == 0 {
		return nil
	}

	return &ErrorsList{
		Errs: errs,
	}
}

func (ce *ErrorsList) Error() string {
	if ce.Errs == nil {
		return ""
	}

	var errs []string
	for _, err := range ce.Errs {
		errs = append(errs, err.Error())
	}

	return strings.Join(errs[:], ", ")
}
