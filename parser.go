package dcssa

import (
	"errors"
	"io"
)

// Parser represents a parser.
type Parser struct {
	s   *Scanner
	buf struct {
		tok Token  // last read token
		lit string // last read literal
		n   int    // buffer size (max=1)
	}
}

// NewParser returns a new instance of Parser.
func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewScanner(r)}
}

// scan returns the next token from the underlying scanner.
// If a token has been unscanned then read that instead.
func (p *Parser) scan() (tok Token, lit string) {
	// If we have a token on the buffer, then return it.
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.lit
	}

	// Otherwise read the next token from the scanner.
	tok, lit = p.s.Scan()

	// Save it to the buffer in case we unscan later.
	p.buf.tok, p.buf.lit = tok, lit

	return
}

// unscan pushes the previously read token back onto the buffer.
func (p *Parser) unscan() { p.buf.n = 1 }

// scanIgnoreWhitespace scans the next non-whitespace token.
func (p *Parser) scanIgnoreWhitespace() (tok Token, lit string) {
	tok, lit = p.scan()
	if tok == WS {
		tok, lit = p.scan()
	}
	return
}

// scanUntilNewline scans until a newline, multiple newlines or EOF.
func (p *Parser) scanUntilNewline() (tok Token, lit string) {
	for tok != EOF && tok != NL && tok != MNL {
		tok, lit = p.scan()
	}
	return tok, lit
}

// ParseVersion parses the DCSS version.
func (p *Parser) ParseVersion() (string, error) {
	var tok Token
	var lit string
	for tok != EOF && tok != NUMBER && tok != NL && tok != MNL {
		tok, lit = p.scan()
	}
	if tok != NUMBER {
		return "", errors.New("could not find a number")
	}
	version := lit
	tok, lit = p.scan()
	if tok == DOT {
		version += lit
		tok, lit = p.scan()
		if tok == NUMBER {
			version += lit
			tok, lit = p.scan()
			if tok == DOT {
				version += lit
				tok, lit = p.scan()
				if tok == NUMBER {
					version += lit
					tok, lit = p.scan()
				}
			}
		}
	}
	if tok != NL && tok != MNL {
		p.scanUntilNewline()
	}
	return version, nil
}
