package parser

import (
	"0x822a5b87/test-pratt-parsing/ast"
	"0x822a5b87/test-pratt-parsing/token"
	"fmt"
	"github.com/bzick/tokenizer"
)

type Precedence int

const (
	Number  Precedence = 10 // Number integer,number,identifier...
	Sum     Precedence = 20 // +
	Product Precedence = 30 // *
	Prefix  Precedence = 40 // -X or !X
)

var precedences = map[tokenizer.TokenKey]Precedence{
	tokenizer.TokenInteger: Number,
	token.TypePlus:         Sum,
	token.TypeAsterisk:     Product,
	token.TypeExclamation:  Prefix,
}

// prefixParseFn parse prefix operator
type prefixParseFn func() ast.ExprNode

// infixParseFn parse infix operator
type infixParseFn func(lhs ast.ExprNode) ast.ExprNode

type Parser struct {
	stream         *tokenizer.Stream
	prefixParseFns map[tokenizer.TokenKey]prefixParseFn
	infixParseFns  map[tokenizer.TokenKey]infixParseFn
}

func NewParser(input string) *Parser {
	tk := token.NewTokenizer()
	p := Parser{
		stream:         tk.ParseString(input),
		prefixParseFns: make(map[tokenizer.TokenKey]prefixParseFn),
		infixParseFns:  make(map[tokenizer.TokenKey]infixParseFn),
	}

	p.registerPrefixParseFn(tokenizer.TokenInteger, p.parseInteger)
	p.registerPrefixParseFn(token.TypeMinus, p.parsePrefixOperator)

	p.registerInfixParseFn(token.TypePlus, p.parseInfixOperator)
	p.registerInfixParseFn(token.TypeAsterisk, p.parseInfixOperator)
	p.registerInfixParseFn(token.TypeExclamation, p.parseSuffixOperator)

	return &p
}

func (p *Parser) ParseExpression(precedence Precedence) ast.ExprNode {
	prefixFn := p.getPrefixParseFn(p.currentTokenType())
	// the left-hand side
	lhs := prefixFn()
	// if we don't reach the end of stream and the precedence of the peek token is greater than the current token
	// then we must parse peek token first, and the parsing result of the current token
	// should be lhs of the peek token
	for p.currentTokenIsValid() && precedence < p.peekPrecedence() {
		infixFn := p.getInfixParseFn(p.currentTokenType())
		lhs = infixFn(lhs)
	}
	return lhs
}

func (p *Parser) Close() {
	p.stream.Close()
}

func (p *Parser) goNext() {
	p.stream.GoNext()
}

func (p *Parser) peekPrecedence() Precedence {
	precedence, ok := precedences[p.currentTokenType()]
	if !ok {
		panic(fmt.Errorf("prefix parse precedence not found for [%d]", p.currentTokenType()))
	}
	return precedence
}

func (p *Parser) currentTokenStringValue() string {
	return p.stream.CurrentToken().ValueString()
}

func (p *Parser) currentTokenType() tokenizer.TokenKey {
	return p.stream.CurrentToken().Key()
}

func (p *Parser) currentTokenIsValid() bool {
	return p.stream.IsValid()
}

func (p *Parser) peekTokenIs(tokenType tokenizer.TokenKey) bool {
	return p.stream.NextToken().Is(tokenType)
}

func (p *Parser) getPrefixParseFn(tokenType tokenizer.TokenKey) prefixParseFn {
	fn, ok := p.prefixParseFns[tokenType]
	if !ok {
		panic(fmt.Errorf("prefix parse fn not found for [%d]", tokenType))
	}
	return fn
}

func (p *Parser) getInfixParseFn(tokenType tokenizer.TokenKey) infixParseFn {
	fn, ok := p.infixParseFns[tokenType]
	if !ok {
		panic(fmt.Errorf("infix parse fn not found for [%d]", tokenType))
	}
	return fn
}

func (p *Parser) registerPrefixParseFn(tokenType tokenizer.TokenKey, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfixParseFn(tokenType tokenizer.TokenKey, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

func (p *Parser) parseInteger() ast.ExprNode {
	node := ast.ValueNode{Value: p.stream.CurrentToken().ValueInt64()}
	p.goNext()
	return &node
}

func (p *Parser) parsePrefixOperator() ast.ExprNode {
	node := ast.PrefixOperatorNode{
		Operator: p.currentTokenStringValue(),
	}
	p.goNext()
	node.Rhs = p.ParseExpression(Prefix)
	return &node
}

func (p *Parser) parseInfixOperator(lhs ast.ExprNode) ast.ExprNode {
	node := ast.InfixOperatorNode{
		Lhs:      lhs,
		Operator: p.stream.CurrentToken().ValueString(),
	}

	precedence := precedences[p.currentTokenType()]
	p.goNext()
	if p.currentTokenIsValid() {
		node.Rhs = p.ParseExpression(precedence)
	}
	return &node
}

// parseSuffixOperator handle suffix operator
// the main difference between infix operators and suffix operators is that
// infix operators is followed by number/bool/identifier, whereas suffix operators is followed by an infix operator
func (p *Parser) parseSuffixOperator(lhs ast.ExprNode) ast.ExprNode {
	node := ast.InfixOperatorNode{
		Lhs:      lhs,
		Operator: p.stream.CurrentToken().ValueString(),
	}

	p.goNext()

	if p.currentTokenIsValid() {
		infixFn := p.getInfixParseFn(p.currentTokenType())
		return infixFn(&node)
	} else {
		return &node
	}
}
