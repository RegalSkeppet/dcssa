package dcssa

import (
	"bufio"
	"bytes"
	"io"
)

type Token int

// All possible tokens.
const (
	// Special tokens
	ILLEGAL Token = iota
	EOF           // End of file
	WS            // Whitespace
	NL            // Newline
	MNL           // Multiple newlines

	// Literals
	WORD   // A word
	NUMBER // A number

	// Misc characters
	CR          // \r
	COMMA       // ,
	DOT         // .
	COLON       // :
	BPAREN      // (
	EPAREN      // )
	SLASH       // /
	ASTERISK    // *
	AT          // @
	BCURLY      // {
	ECURLY      // }
	PLUS        // +
	MINUS       // -
	HASHTAG     // #
	PERCENT     // %
	EXCLAMATION // !
	QUESTION    // ?
	EQUALS      // =
	AMPERSAND   // &
	PIPE        // |
	INFINITY    // ∞
	BBRACKET    // []
	EBRACKET    // ]
	UNDERSCORE  // _
	APOSTROPHE  // '
)

var eof = rune(0)

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t'
}

func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

func isNewline(ch rune) bool {
	return ch == '\r' || ch == '\n'
}

// Scanner represents a lexical scanner.
type Scanner struct {
	r *bufio.Reader
}

// NewScanner returns a new instance of Scanner.
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

// read reads the next rune from the bufferred reader.
// Returns the rune(0) if an error occurs (or io.EOF is returned).
func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

// unread places the previously read rune back on the reader.
func (s *Scanner) unread() { _ = s.r.UnreadRune() }

// Scan returns the next token and literal value.
func (s *Scanner) Scan() (tok Token, lit string) {
	// Read the next rune.
	ch := s.read()

	// If we see whitespace then consume all contiguous whitespace.
	// If we see a letter then consume as an ident or reserved word.
	if isWhitespace(ch) {
		s.unread()
		return s.scanWhitespace()
	} else if isLetter(ch) {
		s.unread()
		return s.scanWord()
	} else if isDigit(ch) {
		s.unread()
		return s.scanNumber()
	} else if isNewline(ch) {
		s.unread()
		return s.scanNewline()
	}

	// Otherwise read the individual character.
	switch ch {
	case eof:
		return EOF, ""
	case ',':
		return COMMA, string(ch)
	case '.':
		return DOT, string(ch)
	case ':':
		return COLON, string(ch)
	case '(':
		return BPAREN, string(ch)
	case ')':
		return EPAREN, string(ch)
	case '/':
		return SLASH, string(ch)
	case '*':
		return ASTERISK, string(ch)
	case '@':
		return AT, string(ch)
	case '{':
		return BCURLY, string(ch)
	case '}':
		return ECURLY, string(ch)
	case '+':
		return PLUS, string(ch)
	case '-':
		return MINUS, string(ch)
	case '#':
		return HASHTAG, string(ch)
	case '%':
		return PERCENT, string(ch)
	case '!':
		return EXCLAMATION, string(ch)
	case '?':
		return QUESTION, string(ch)
	case '=':
		return EQUALS, string(ch)
	case '&':
		return AMPERSAND, string(ch)
	case '|':
		return PIPE, string(ch)
	case '∞':
		return INFINITY, string(ch)
	case '[':
		return BBRACKET, string(ch)
	case ']':
		return EBRACKET, string(ch)
	case '_':
		return EBRACKET, string(ch)
	case '\'':
		return APOSTROPHE, string(ch)
	}

	return ILLEGAL, string(ch)
}

// scanWhitespace consumes the current rune and all contiguous whitespace.
func (s *Scanner) scanWhitespace() (tok Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent whitespace character into the buffer.
	// Non-whitespace characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isWhitespace(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return WS, buf.String()
}

// scanIdent consumes the current rune and all contiguous ident runes.
func (s *Scanner) scanWord() (tok Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent ident character into the buffer.
	// Non-ident characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isLetter(ch) && !isDigit(ch) && ch != '_' {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	// Otherwise return as a regular identifier.
	return WORD, buf.String()
}

// scanIdent consumes the current rune and all contiguous ident runes.
func (s *Scanner) scanNumber() (tok Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent ident character into the buffer.
	// Non-ident characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isDigit(ch) {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	// Otherwise return as a regular identifier.
	return NUMBER, buf.String()
}

// scanIdent consumes the current rune and all contiguous ident runes.
func (s *Scanner) scanNewline() (tok Token, lit string) {
	count := 0
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isNewline(ch) {
			s.unread()
			break
		} else if ch == '\n' {
			count++
		}
	}

	if count > 1 {
		return MNL, ""
	}

	return NL, ""
}
