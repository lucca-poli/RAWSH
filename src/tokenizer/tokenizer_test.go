package tokenizer

import (
	"testing"
)

func TestTokenizerDoubleQuoted(t *testing.T) {
	result, _ := Tokenize("echo \"batata\"")
	expected := []string{"echo", "batata"}

	length := len(expected)

	if len(result) != length {
		t.Errorf("Lengths are different")
	}

	for i := 0; i < length; i++ {
		if result[i] != expected[i] {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	}
}

func TestTokenizerSingleQuoted(t *testing.T) {
	result, _ := Tokenize("echo 'batata'")
	expected := []string{"echo", "batata"}

	length := len(expected)

	if len(result) != length {
		t.Errorf("Lengths are different")
	}

	for i := 0; i < length; i++ {
		if result[i] != expected[i] {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	}
}

func TestTokenizerMixedQuotes(t *testing.T) {
	result, _ := Tokenize("\"echo\" \"arroz com 'batata e' feijao\"")
	expected := []string{"echo", "arroz com 'batata e' feijao"}

	length := len(expected)

	if len(result) != length {
		t.Errorf("Lengths are different")
	}

	for i := 0; i < length; i++ {
		if result[i] != expected[i] {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	}
}

func TestTokenizerOperator(t *testing.T) {
	result, _ := Tokenize("echo \"batata com arroz\"|grep arroz")
	expected := []string{"echo", "batata com arroz", "|", "grep", "arroz"}

	length := len(expected)

	if len(result) != length {
		t.Errorf("Lengths are different")
	}

	for i := 0; i < length; i++ {
		if result[i] != expected[i] {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	}
}

func TestTokenizerUnquoted(t *testing.T) {
	result, _ := Tokenize("   ls -a ~/*")
	expected := []string{"ls", "-a", "~/*"}

	length := len(expected)

	if len(result) != length {
		t.Errorf("Lengths are different")
	}

	for i := 0; i < length; i++ {
		if result[i] != expected[i] {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	}
}
