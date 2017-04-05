package token

// Type - The type for definiting tokens
type Type string

// Token - Token object
type Token struct {
	Type    Type
	Literal string
}
