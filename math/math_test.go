package math

import (
	"math"
	"testing"
)

var abstests = []struct {
	num int
	abs int
}{
	{-13, 13},
	{123456, 123456},
	{0, 0},
	{math.MaxInt32, math.MaxInt32},
	{math.MinInt32, -math.MinInt32},
}

func TestAbsi(t *testing.T) {
	for _, test := range abstests {
		a := Absi(test.num)
		if a != test.abs {
			t.Errorf("expected '%v' but got '%v'", test.abs, a)
		}
	}
}

var roundtests = []struct {
	num   float64
	round int
}{
	{123.4999999, 123},
	{0.500000000, 1},
	{1.0, 1},
	{-99.9999, -100},
	{-0.1, 0},
}

func TestRound(t *testing.T) {
	for _, test := range roundtests {
		r := Round(test.num)
		if r != test.round {
			t.Errorf("expected '%v' but got '%v'", test.round, r)
		}
	}
}

var mintests = []struct {
	num1, num2, min int
}{
	{1, 2, 1},
	{2, 1, 1},
	{-99, -1, -99},
	{-1, 1, -1},
	{0, 0, 0},
}

func TestMini(t *testing.T) {
	for _, test := range mintests {
		m := Mini(test.num1, test.num2)
		if m != test.min {
			t.Errorf("expected '%v' but got '%v'", test.min, m)
		}
	}
}

var maxtests = []struct {
	num1, num2, max int
}{
	{1, 2, 2},
	{2, 1, 2},
	{-99, -1, -1},
	{-1, 1, 1},
	{0, 0, 0},
}

func TestMaxi(t *testing.T) {
	for _, test := range maxtests {
		m := Maxi(test.num1, test.num2)
		if m != test.max {
			t.Errorf("expected '%v' but got '%v'", test.max, m)
		}
	}
}
