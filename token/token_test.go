package token

import (
	"github.com/bzick/tokenizer"
	"testing"
)

func TestTokenizer(t *testing.T) {
	input := `(10 + 1) - 2! * -3 / 4^10^10`

	expected := []struct {
		tokenType    tokenizer.TokenKey
		literalValue string
	}{
		{TypeLPa, "("},
		{tokenizer.TokenInteger, "10"},
		{TypePlus, "+"},
		{tokenizer.TokenInteger, "1"},
		{TypeRPa, ")"},
		{TypeMinus, "-"},
		{tokenizer.TokenInteger, "2"},
		{TypeExclamation, "!"},
		{TypeAsterisk, "*"},
		{TypeMinus, "-"},
		{tokenizer.TokenInteger, "3"},
		{TypeSlash, "/"},
		{tokenizer.TokenInteger, "4"},
		{TypeCaret, "^"},
		{tokenizer.TokenInteger, "10"},
		{TypeCaret, "^"},
		{tokenizer.TokenInteger, "10"},
	}

	tk := NewTokenizer()
	stream := tk.ParseString(input)
	defer stream.Close()

	for tokenOffset, expectedToken := range expected {
		currentToken := stream.CurrentToken()

		// check type
		if !currentToken.Is(expectedToken.tokenType) {
			t.Fatalf("token offset = [%d] type not match, expected [%d], got [%d]",
				tokenOffset, expectedToken.tokenType, currentToken.Key())
		}

		// check literal value
		if currentToken.ValueString() != expectedToken.literalValue {
			t.Fatalf("token offset = [%d], expected [%s], got [%s]",
				tokenOffset, currentToken.ValueString(), expectedToken.literalValue)
		}

		stream.GoNext()
		tokenOffset++
	}
}
