package main

import "testing"

func TestSexp(t *testing.T) {

	sexp := sexp{
		op: "and",
		args: []evaluator{
			sexp{op: "or", args: []evaluator{matcher("too"), matcher("boo")}},
			sexp{op: "not", args: []evaluator{matcher("crumble")}},
		},
	}

	if sexp.Eval("footle boolte") != true {
		t.Errorf("expected true, have false")
	}

	if sexp.Eval("boo crumble") != false {
		t.Errorf("expected false, have true")
	}
}

func TestSexpNot(t *testing.T) {

	sexp := sexp{
		op:   "not",
		args: []evaluator{matcher("footle")},
	}

	if sexp.Eval("boolte tootle") != true {
		t.Errorf("expected true, have false")
	}

	if sexp.Eval("footle mcdootle") != false {
		t.Errorf("expected false, have true")
	}
}
