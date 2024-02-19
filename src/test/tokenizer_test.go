package internal_test

import (
	"RAWSH/src/internal"
	"testing"
)

// func TestTokenizerClassifier(t *testing.T) {
// 	result, _ := internal.Tokenizer("echo 'batata com arroz' | grep arroz")
//     expected := []string{"word", "word", "operator", "word", "word"}
//
//     for i, token := range result {
//         if token.TokenType != expected[i] {
//             t.Errorf("Expected %s, got %s", expected[i], token.TokenType)
//         }
//     }
// }

func TestTokenizerDoubleQuoted(t *testing.T) {
	result, _ := internal.Tokenizer("echo \"batata\"")
	expected := []string{"echo", "\"batata\""}

	length := len(expected)

	if len(result) != length {
		t.Errorf("Lengths are different")
	}

	for i := 0; i < length; i++ {
		if result[i].Value != expected[i] {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	}
}

func TestTokenizerSingleQuoted(t *testing.T) {
	result, _ := internal.Tokenizer("echo 'batata com arroz' | grep arroz")
	expected := []string{"echo", "'batata com arroz'", "|", "grep", "arroz"}

	length := len(expected)

	if len(result) != length {
		t.Errorf("Lengths are different")
	}

	for i := 0; i < length; i++ {
		if result[i].Value != expected[i] {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	}
}

func TestTokenizerMixedQuotes(t *testing.T) {
	result, _ := internal.Tokenizer("\"echo\" \"arroz com 'batata e' feijao\"")
	expected := []string{"\"echo\"", "\"arroz com 'batata e' feijao\""}

	length := len(expected)

	if len(result) != length {
		t.Errorf("Lengths are different")
	}

	for i := 0; i < length; i++ {
		if result[i].Value != expected[i] {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	}
}

func TestTokenizerUnquoted(t *testing.T) {
	result, _ := internal.Tokenizer("   ls -a ~/*")
	expected := []string{"ls", "-a", "~/*"}

	length := len(expected)

	if len(result) != length {
		t.Errorf("Lengths are different")
	}

	for i := 0; i < length; i++ {
		if result[i].Value != expected[i] {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	}
}
