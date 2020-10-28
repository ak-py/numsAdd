package main

import (
	"testing"
)

func TestAddExample(t *testing.T) {
	result := Add([]int{1, 2, 3, 4, 5, 6}, []int{1, 2, 3, 4, 5, 6})

	expected := "777777"
	if result != expected {
		t.Errorf("result was incorrect, got: %#v, want: %#v.", result, expected)
	}
}

func TestAddNilAndEmpty(t *testing.T) {
	result := Add(nil, nil)
	expected := ""
	if result != expected {
		t.Errorf("result was incorrect, got: %#v, want: %#v.", result, expected)
	}

	result = Add([]int{}, []int{})
	if result != expected {
		t.Errorf("result was incorrect, got: %#v, want: %#v.", result, expected)
	}
}

func TestAddLarge(t *testing.T) {

	var (
		num1 []int
		num2 []int
	)

	num1length := 15
	for i := 1; i <= num1length; i++ {
		num1 = append(num1, i%10)
	}

	num2length := 6
	for i := 1; i <= num2length; i++ {
		num2 = append(num2, i%10)
	}

	result := Add(num1, num2)
	expected := "123456789666666"
	if result != expected {
		t.Errorf("result was incorrect, got: %#v, want: %#v.", result, expected)
	}

	result = Add(num2, num1)
	expected = "543210987777777"
	if result != expected {
		t.Errorf("result was incorrect, got: %#v, want: %#v.", result, expected)
	}
}

func TestAddSmall(t *testing.T) {

	var (
		num1 []int
		num2 []int
	)

	num1length := 5
	for i := 1; i <= num1length; i++ {
		num1 = append(num1, i%10)
	}

	num2length := 8
	for i := 1; i <= num2length; i++ {
		num2 = append(num2, i%10)
	}

	result := Add(num1, num2)
	expected := "87666666"
	if result != expected {
		t.Errorf("result was incorrect, got: %#v, want: %#v.", result, expected)
	}

	result = Add(num2, num1)
	expected = "12399999"
	if result != expected {
		t.Errorf("result was incorrect, got: %#v, want: %#v.", result, expected)
	}
}
