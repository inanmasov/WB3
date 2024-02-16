package main

import (
	"log"
	"testing"
)

func TestEmpty(t *testing.T) {
	args := &Args{
		flags: Flags{
			after:      0,
			before:     0,
			context:    0,
			count:      false,
			ignoreCase: false,
			invert:     false,
			fixed:      false,
			lineNum:    false,
		},
		pattern: "",
		files:   []string{"input1.txt"},
	}

	var expected = []string{
		"a",
		"b",
		"c",
		"d",
		"e",
		"f",
	}

	results, err := processFile(args)
	if err != nil {
		log.Fatalln(err)
	}
	res := results[0]

	if len(res) != len(expected) {
		t.Errorf("Expected %d lines, but got %d", len(expected), len(res))
	}

	for i, line := range res {
		if line != expected[i] {
			t.Errorf("Expected %v, but got %v", expected[i], line)
		}
	}
}

func TestCount(t *testing.T) {
	args := &Args{
		flags: Flags{
			after:      0,
			before:     0,
			context:    0,
			count:      true,
			ignoreCase: false,
			invert:     false,
			fixed:      false,
			lineNum:    false,
		},
		pattern: "",
		files:   []string{"input1.txt"},
	}

	var expected = []string{
		"6",
	}

	results, err := processFile(args)
	if err != nil {
		log.Fatalln(err)
	}
	res := results[0]

	if len(res) != len(expected) {
		t.Errorf("Expected %d lines, but got %d", len(expected), len(res))
	}

	for i, line := range res {
		if line != expected[i] {
			t.Errorf("Expected %v, but got %v", expected[i], line)
		}
	}
}

func TestAfter(t *testing.T) {
	args := &Args{
		flags: Flags{
			after:      3,
			before:     0,
			context:    0,
			count:      false,
			ignoreCase: false,
			invert:     false,
			fixed:      false,
			lineNum:    false,
		},
		pattern: "c",
		files:   []string{"input1.txt"},
	}

	var expected = []string{
		"c",
		"d",
		"e",
		"f",
	}

	results, err := processFile(args)
	if err != nil {
		log.Fatalln(err)
	}
	res := results[0]

	if len(res) != len(expected) {
		t.Errorf("Expected %d lines, but got %d", len(expected), len(res))
	}

	for i, line := range res {
		if line != expected[i] {
			t.Errorf("Expected %v, but got %v", expected[i], line)
		}
	}
}

func TestBefore(t *testing.T) {
	args := &Args{
		flags: Flags{
			after:      0,
			before:     3,
			context:    0,
			count:      false,
			ignoreCase: false,
			invert:     false,
			fixed:      false,
			lineNum:    false,
		},
		pattern: "c",
		files:   []string{"input1.txt"},
	}

	var expected = []string{
		"a",
		"b",
		"c",
	}

	results, err := processFile(args)
	if err != nil {
		log.Fatalln(err)
	}
	res := results[0]

	if len(res) != len(expected) {
		t.Errorf("Expected %d lines, but got %d", len(expected), len(res))
	}

	for i, line := range res {
		if line != expected[i] {
			t.Errorf("Expected %v, but got %v", expected[i], line)
		}
	}
}

func TestContext(t *testing.T) {
	args := &Args{
		flags: Flags{
			after:      0,
			before:     0,
			context:    3,
			count:      false,
			ignoreCase: false,
			invert:     false,
			fixed:      false,
			lineNum:    false,
		},
		pattern: "c",
		files:   []string{"input1.txt"},
	}

	var expected = []string{
		"a",
		"b",
		"c",
		"d",
		"e",
		"f",
	}

	results, err := processFile(args)
	if err != nil {
		log.Fatalln(err)
	}
	res := results[0]

	if len(res) != len(expected) {
		t.Errorf("Expected %d lines, but got %d", len(expected), len(res))
	}

	for i, line := range res {
		if line != expected[i] {
			t.Errorf("Expected %v, but got %v", expected[i], line)
		}
	}
}

func TestIgnoreCase(t *testing.T) {
	args := &Args{
		flags: Flags{
			after:      0,
			before:     0,
			context:    0,
			count:      false,
			ignoreCase: true,
			invert:     false,
			fixed:      false,
			lineNum:    false,
		},
		pattern: "a",
		files:   []string{"input2.txt"},
	}

	var expected = []string{
		"Ab",
		"abr",
	}

	results, err := processFile(args)
	if err != nil {
		log.Fatalln(err)
	}
	res := results[0]

	if len(res) != len(expected) {
		t.Errorf("Expected %d lines, but got %d", len(expected), len(res))
	}

	for i, line := range res {
		if line != expected[i] {
			t.Errorf("Expected %v, but got %v", expected[i], line)
		}
	}
}

func TestInvert(t *testing.T) {
	args := &Args{
		flags: Flags{
			after:      0,
			before:     0,
			context:    0,
			count:      false,
			ignoreCase: false,
			invert:     true,
			fixed:      false,
			lineNum:    false,
		},
		pattern: "a",
		files:   []string{"input2.txt"},
	}

	var expected = []string{
		"Ab",
		"b",
		"C",
	}

	results, err := processFile(args)
	if err != nil {
		log.Fatalln(err)
	}
	res := results[0]

	if len(res) != len(expected) {
		t.Errorf("Expected %d lines, but got %d", len(expected), len(res))
	}

	for i, line := range res {
		if line != expected[i] {
			t.Errorf("Expected %v, but got %v", expected[i], line)
		}
	}
}

func TestLineNum(t *testing.T) {
	args := &Args{
		flags: Flags{
			after:      0,
			before:     0,
			context:    0,
			count:      false,
			ignoreCase: false,
			invert:     false,
			fixed:      false,
			lineNum:    true,
		},
		pattern: "a",
		files:   []string{"input2.txt"},
	}

	var expected = []string{
		"4:abr",
	}

	results, err := processFile(args)
	if err != nil {
		log.Fatalln(err)
	}
	res := results[0]

	if len(res) != len(expected) {
		t.Errorf("Expected %d lines, but got %d", len(expected), len(res))
	}

	for i, line := range res {
		if line != expected[i] {
			t.Errorf("Expected %v, but got %v", expected[i], line)
		}
	}
}
