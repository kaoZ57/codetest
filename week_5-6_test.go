package week5_6

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFizzWord(t *testing.T) {
	testCases := []struct {
		input [2]int
		want  string
	}{
		{[2]int{1, 1}, "Fizz"},
		{[2]int{2, 1}, "Fizz"},
		{[2]int{2, 3}, "2"},
		{[2]int{9, 7}, "9"},
	}
	for _, test := range testCases {
		got := FizzWord(test.input[0], test.input[1])
		if got != test.want {
			t.Errorf("\nunexpected\n\tgot: %v\n\twant: %v", got, test.want)
		}
	}
}

func TestMultiplicationTable(t *testing.T) {
	testCases := []struct {
		description string
		input       int
		want        []int
	}{
		{"3 time table", 3, []int{3, 6, 9, 12, 15, 18, 21, 24, 27, 30, 33, 36}},
		{"21 time table", 21, []int{21, 42, 63, 84, 105, 126, 147, 168, 189, 210, 231, 252}},
		{"79 time table", 79, []int{79, 158, 237, 316, 395, 474, 553, 632, 711, 790, 869, 948}},
	}
	for _, test := range testCases {
		t.Run(test.description, func(t *testing.T) {
			got := MultiplicationTable(test.input)
			if reflect.DeepEqual(test.want, got) == false {
				t.Errorf("\nunexpected\n\tgot: %v\n\twant: %v", got, test.want)
			}
		})
	}
}

func TestStringMultiplicationTable(t *testing.T) {
	testCases := []struct {
		description string
		input       int
		want        []string
	}{
		{"3 time table", 3, []string{"3 x 1 = 3", "3 x 2 = 6", "3 x 3 = 9", "3 x 4 = 12",
			"3 x 5 = 15", "3 x 6 = 18", "3 x 7 = 21", "3 x 8 = 24", "3 x 9 = 27", "3 x 10 = 30",
			"3 x 11 = 33", "3 x 12 = 36"}},
		{"21 time table", 21, []string{"21 x 1 = 21", "21 x 2 = 42", "21 x 3 = 63", "21 x 4 = 84",
			"21 x 5 = 105", "21 x 6 = 126", "21 x 7 = 147", "21 x 8 = 168", "21 x 9 = 189",
			"21 x 10 = 210", "21 x 11 = 231", "21 x 12 = 252"}},
		{"79 time table", 79, []string{"79 x 1 = 79", "79 x 2 = 158", "79 x 3 = 237", "79 x 4 = 316",
			"79 x 5 = 395", "79 x 6 = 474", "79 x 7 = 553", "79 x 8 = 632", "79 x 9 = 711",
			"79 x 10 = 790", "79 x 11 = 869", "79 x 12 = 948"}},
	}
	for _, test := range testCases {
		t.Run(test.description, func(t *testing.T) {
			got := StringMultiplicationTable(test.input)
			if reflect.DeepEqual(test.want, got) == false {
				t.Errorf("\nunexpected\n\tgot: %v\n\twant: %v", got, test.want)
			}
		})
	}
}

func TestDeleteIntItem(t *testing.T) {
	testCases := []struct {
		description string
		parameter1  []int
		parameter2  int
		want        []int
	}{
		{"empty array", []int{}, 3, []int{}},
		{"not contain delete item", []int{1, 2, 3}, 4, []int{1, 2, 3}},
		{"contain one delete item", []int{1, 2, 3}, 3, []int{1, 2}},
		{"contain more than one delete item", []int{1, 2, 3, 4, 3, 36}, 3, []int{1, 2, 4, 36}},
	}
	for _, test := range testCases {
		t.Run(test.description, func(t *testing.T) {
			got := DeleteIntItem(test.parameter1, test.parameter2)
			if reflect.DeepEqual(test.want, got) == false {
				t.Errorf("\nunexpected\n\tgot: %v\n\twant: %v", got, test.want)
			}
		})
	}
}

func TestGrade(t *testing.T) {
	testCases := []struct {
		input float32
		want  string
	}{
		{0, "F"},
		{49.99, "F"},
		{50, "D"},
		{69, "C"},
		{79, "B"},
		{80, "A"},
	}
	for _, test := range testCases {
		//https://golang.org/pkg/fmt/#Sprintf
		t.Run(fmt.Sprintf("point: %f", test.input), func(t *testing.T) {
			got := Grade(test.input)
			if got != test.want {
				t.Errorf("\nunexpected\n\tgot: %v\n\twant: %v", got, test.want)
			}
		})
	}
}
