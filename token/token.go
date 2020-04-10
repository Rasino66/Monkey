package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// 識別子＋リテラル
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   //123456

	// 識別子
	ASSIGN = "="
	PLUS   = "+"

	// デリメタ
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPANEN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// キーワード
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
