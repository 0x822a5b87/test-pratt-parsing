package token

import (
	"github.com/bzick/tokenizer"
)

const (
	_ = iota
	TypePlus
	TypeMinus
	TypeAsterisk
	TypeSlash
	TypeExclamation
	TypeCaret
	TypeLPa
	TypeRPa
	TypeSemicolon
)

func NewTokenizer() *tokenizer.Tokenizer {
	// configure tokenizer
	tokenizers := tokenizer.New()
	tokenizers.DefineTokens(TypePlus, []string{"+"})
	tokenizers.DefineTokens(TypeMinus, []string{"-"})
	tokenizers.DefineTokens(TypeAsterisk, []string{"*"})
	tokenizers.DefineTokens(TypeSlash, []string{"/"})
	tokenizers.DefineTokens(TypeExclamation, []string{"!"})
	tokenizers.DefineTokens(TypeCaret, []string{"^"})
	tokenizers.DefineTokens(TypeLPa, []string{"("})
	tokenizers.DefineTokens(TypeRPa, []string{")"})
	tokenizers.DefineTokens(TypeSemicolon, []string{";"})
	return tokenizers
}
