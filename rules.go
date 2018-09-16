package main

import "strings"

type evaluator interface {
	Eval(content string) bool
}
type sexp struct {
	op   string
	args []evaluator
}

func (s sexp) Eval(content string) bool {
	switch s.op {
	case "not":
		if len(s.args) == 0 {
			return true
		}
		return !s.args[0].Eval(content)

	case "and":
		for _, a := range s.args {
			if !a.Eval(content) {
				return false
			}
		}
		return true

	case "or":
		for _, a := range s.args {
			if a.Eval(content) {
				return true
			}
		}
		return false
	}
	return false
}

type matcher string

func (m matcher) Eval(content string) bool {
	return strings.Contains(content, string(m))
}
