package service

import (
	"errors"
	"strings"

	"github.com/demidenkoalex/sprint6-final/pkg/morse"
)

var (
	ErrEmptyInput = errors.New("empty input")
)

func IsMorse(s string) (bool, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return false, ErrEmptyInput
	}
	hasDotOrDash := false
	for _, r := range s {
		switch r {
		case '.', '-':
			hasDotOrDash = true
		case ' ', '\t', '\n', '/':
		default:
			return false, nil
		}
	}
	return hasDotOrDash, nil
}

func Convert(s string) (string, error) {

	result, err := IsMorse(s)

	if err != nil {
		return "", err
	}

	if result {
		return morse.ToText(s), err
	}
	return morse.ToMorse(s), err
}
