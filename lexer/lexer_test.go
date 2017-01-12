package lexer

import (
	"testing"

	"FirstLanguage/token"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	test := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
	}
}
