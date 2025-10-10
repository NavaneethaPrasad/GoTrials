package main

import (
	"reflect"
	"testing"
)

func TestFizzbizz(t *testing.T) {
	result := fizzbizz(15)
	expected := []string{"1", "2", "fizz", "4", "bizz", "fizz", "7", "8", "fizz", "bizz", "11", "fizz", "13", "14", "fizzbizz"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func TestPalindrome(t *testing.T) {
	result := palindrome("Mom")
	if result != true {
		t.Errorf("Expected true but got %t\n", result)
	}
	result = palindrome("hai")
	if result != false {
		t.Errorf("Expected false but got %t\n", result)
	}
	result = palindrome("malayalam")
	if result != true {
		t.Errorf("Expected true but got %t\n", result)
	}
}

func TestAbbreviate(t *testing.T) {
	result := abbreviate("Indian Institute of Management")
	if result != "IIM" {
		t.Errorf("Expected IIM but got %s\n", result)
	}
	result = abbreviate("Automated   Teller Machine")
	if result == "atm" {
		t.Errorf("Expected ATM but got %s\n", result)
	}
}

func TestPangram(t *testing.T) {
	result := pangram("The quick brown fox jumps over the lazy dog")
	if result != true {
		t.Errorf("Expected true but got %t\n", result)
	}
	result = pangram("I am Navaneetha")
	if result != false {
		t.Errorf("Expected false but got %t\n", result)
	}
}

func TestFrequency(t *testing.T) {
	result := frequency("Bobby")
	expected := map[string]int{
		"B": 1,
		"o": 1,
		"b": 2,
		"y": 1,
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but %v", expected, result)
	}
	result = frequency("hai")
	expected = map[string]int{
		"h": 1,
		"a": 1,
		"i": 1,
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but %v", expected, result)
	}
}
