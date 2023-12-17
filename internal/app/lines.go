package app

import (
	"bufio"
	"errors"
	"strings"
)

func ProcessLines(input string, f func(line string) error) error {
	var errs []error
	s := bufio.NewScanner(strings.NewReader(input))
	for s.Scan() {
		l := s.Text()
		if l != "" {
			err := f(l)
			if err != nil {
				errs = append(errs, err)
			}
		}
	}
	if err := s.Err(); err != nil {
		return err
	}
	if len(errs) != 0 {
		return errors.Join(errs...)
	}

	return nil
}
