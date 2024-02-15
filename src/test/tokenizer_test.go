package tokenizer

import (
	tokenizer "RAWSH/src/pkg"
	"testing"
)

func TestTokenizerDoubleQuoted(t *testing.T) {
	result, _ := tokenizer.Tokenize("echo \"batata\"")
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
	result, _ := tokenizer.Tokenize("echo 'batata'")
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
	result, _ := tokenizer.Tokenize("\"echo\" \"arroz com 'batata e' feijao\"")
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

func TestTokenizerUnquoted(t *testing.T) {
	result, _ := tokenizer.Tokenize("   ls -a ~/*")
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
