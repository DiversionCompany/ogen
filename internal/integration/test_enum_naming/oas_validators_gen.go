// Code generated by ogen, DO NOT EDIT.

package api

import (
	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/validate"
)

func (s CustomNamingInt) Validate() error {
	switch s {
	case 1:
		return nil
	case 2:
		return nil
	case 3:
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s CustomNamingStr) Validate() error {
	switch s {
	case "a":
		return nil
	case "b":
		return nil
	case "e":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s PascalExceptionStrat) Validate() error {
	switch s {
	case "1":
		return nil
	case "-2":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s PascalSpecialStrat) Validate() error {
	switch s {
	case "2+2":
		return nil
	case "2-2":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s PascalStrat) Validate() error {
	switch s {
	case "in-sync":
		return nil
	case "out-of-sync":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s *ProbeLivenessOK) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := s.VeryBadEnum.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "VeryBadEnum",
			Error: err,
		})
	}
	if err := func() error {
		if err := s.PascalStrat.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "PascalStrat",
			Error: err,
		})
	}
	if err := func() error {
		if err := s.PascalSpecialStrat.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "PascalSpecialStrat",
			Error: err,
		})
	}
	if err := func() error {
		if err := s.PascalExceptionStrat.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "PascalExceptionStrat",
			Error: err,
		})
	}
	if err := func() error {
		if err := s.CustomNamingStr.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "CustomNamingStr",
			Error: err,
		})
	}
	if err := func() error {
		if err := s.CustomNamingInt.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "CustomNamingInt",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s VeryBadEnum) Validate() error {
	switch s {
	case " ":
		return nil
	case "!":
		return nil
	case "\"":
		return nil
	case "#":
		return nil
	case "$":
		return nil
	case "%":
		return nil
	case "&":
		return nil
	case "'":
		return nil
	case "(":
		return nil
	case ")":
		return nil
	case "*":
		return nil
	case "+":
		return nil
	case ",":
		return nil
	case "-":
		return nil
	case ".":
		return nil
	case "/":
		return nil
	case "0":
		return nil
	case "1":
		return nil
	case "2":
		return nil
	case "3":
		return nil
	case "4":
		return nil
	case "5":
		return nil
	case "6":
		return nil
	case "7":
		return nil
	case "8":
		return nil
	case "9":
		return nil
	case ":":
		return nil
	case ";":
		return nil
	case "<":
		return nil
	case "=":
		return nil
	case ">":
		return nil
	case "?":
		return nil
	case "@":
		return nil
	case "A":
		return nil
	case "B":
		return nil
	case "C":
		return nil
	case "D":
		return nil
	case "E":
		return nil
	case "F":
		return nil
	case "G":
		return nil
	case "H":
		return nil
	case "I":
		return nil
	case "J":
		return nil
	case "K":
		return nil
	case "L":
		return nil
	case "M":
		return nil
	case "N":
		return nil
	case "O":
		return nil
	case "P":
		return nil
	case "Q":
		return nil
	case "R":
		return nil
	case "S":
		return nil
	case "T":
		return nil
	case "U":
		return nil
	case "V":
		return nil
	case "W":
		return nil
	case "X":
		return nil
	case "Y":
		return nil
	case "Z":
		return nil
	case "[":
		return nil
	case "\\":
		return nil
	case "]":
		return nil
	case "^":
		return nil
	case "_":
		return nil
	case "`":
		return nil
	case "a":
		return nil
	case "b":
		return nil
	case "c":
		return nil
	case "d":
		return nil
	case "e":
		return nil
	case "f":
		return nil
	case "g":
		return nil
	case "h":
		return nil
	case "i":
		return nil
	case "j":
		return nil
	case "k":
		return nil
	case "l":
		return nil
	case "m":
		return nil
	case "n":
		return nil
	case "o":
		return nil
	case "p":
		return nil
	case "q":
		return nil
	case "r":
		return nil
	case "s":
		return nil
	case "t":
		return nil
	case "u":
		return nil
	case "v":
		return nil
	case "w":
		return nil
	case "x":
		return nil
	case "y":
		return nil
	case "z":
		return nil
	case "{":
		return nil
	case "|":
		return nil
	case "}":
		return nil
	case "~":
		return nil
	case "id":
		return nil
	case "id desc":
		return nil
	case "classifyAs":
		return nil
	case "classifyAs desc":
		return nil
	case "senderEmailAddress":
		return nil
	case "senderEmailAddress desc":
		return nil
	case "childFolders":
		return nil
	case "messageRules":
		return nil
	case "messages":
		return nil
	case "multiValueExtendedProperties":
		return nil
	case "singleValueExtendedProperties":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
