package dcssa

import (
	"bytes"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
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

// scanUntilNewline scans the next newline or multiple newlines ignoring everything else.
func (p *Parser) scanNextNewline() (tok Token, lit string) {
	var buf bytes.Buffer
	tok, lit = p.scan()
	for tok != EOF && tok != NL && tok != MNL {
		buf.WriteString(lit)
		tok, lit = p.scan()
	}
	return tok, buf.String()
}

// scanNextMultiNewline scans the next newline ignoring everything else.
func (p *Parser) scanNextMultiNewline() (tok Token, lit string) {
	for tok != EOF && tok != MNL {
		tok, lit = p.scan()
	}
	return tok, lit
}

// scanIgnoreIndented scans the first non-indented, non-whitespace token.
func (p *Parser) scanIgnoreIndented() (tok Token, lit string) {
	for {
		tok, lit = p.scan()
		if tok == EOF {
			return tok, lit
		}
		if tok == NL || tok == MNL {
			continue
		}
		if tok == WS {
			tok, lit = p.scanNextNewline()
			if tok == EOF {
				return tok, lit
			}
			continue
		}
		break
	}
	return tok, lit
}

// scanNextTokenIgnoringOthers scans the next chosen token ignoring everything else.
func (p *Parser) scanNextTokenIgnoringOthers(ignoredLit *string, tokens ...Token) (tok Token, lit string) {
	var buf bytes.Buffer
	for {
		tok, lit = p.scan()
		if tok == EOF {
			break
		}
		done := false
		for _, t := range tokens {
			if tok == t {
				done = true
				break
			}
		}
		if !done {
			buf.WriteString(lit)
			continue
		}
		break
	}
	if ignoredLit != nil {
		*ignoredLit = buf.String()
	}
	return tok, lit
}

// ParseFile parses a morgue file into data.
func ParseFile(path string, data *Data) {
	file, err := os.Open(path)
	if err != nil {
		data.FailedReads[path] = err.Error()
		return
	}

	parser := NewParser(file)
	run, err := parser.ParseRun()
	if err != nil {
		data.FailedReads[path] = err.Error()
		return
	}

	data.Runs = append(data.Runs, run)
}

// ParseRun parses a morgue file.
func (p *Parser) ParseRun() (*Run, error) {
	var tok Token
	var lit string
	var err error

	// Parse version.
	for tok != EOF && tok != NUMBER && tok != NL && tok != MNL {
		tok, lit = p.scan()
	}
	if tok != NUMBER {
		return nil, errors.New("could not find a number when parsing version")
	}
	run := NewRun()
	run.Version = lit
	tok, lit = p.scan()
	if tok == DOT {
		run.Version += lit
		tok, lit = p.scan()
		if tok == NUMBER {
			run.Version += lit
			tok, lit = p.scan()
			if tok == DOT {
				run.Version += lit
				tok, lit = p.scan()
				if tok == NUMBER {
					run.Version += lit
					tok, lit = p.scan()
				}
			}
		}
	}
	p.unscan()

	tok, lit = p.scanNextMultiNewline()
	if tok != MNL {
		return nil, errors.New("expected multiple newlines after version: " + lit)
	}

	// Parse score.
	if tok, lit = p.scan(); tok != NUMBER {
		return nil, errors.New("expected number when parsing score: " + lit)
	}
	run.Score, err = strconv.Atoi(lit)
	if err != nil {
		return nil, errors.New("could not parse score int: " + err.Error())
	}
	if tok, lit = p.scanNextNewline(); tok == EOF {
		return nil, errors.New("EOF after score")
	}

	// Parse name and title.
	tok, lit = p.scanIgnoreIndented()
	if tok != WORD && tok != NUMBER && tok != MINUS && tok != DOT && tok != UNDERSCORE && lit != " " {
		return nil, errors.New("expected valid name: " + lit)
	}
	nameAndTitle := []string{}
	for tok != EOF && tok != BPAREN {
		nameAndTitle = append(nameAndTitle, lit)
		tok, lit = p.scan()
	}
	if tok == EOF {
		return nil, errors.New("EOF when parsing name and title")
	}
	nameAndTitleSeparatorIndex := -1
	for i := len(nameAndTitle) - 1; i >= 0; i-- {
		if nameAndTitle[i] == "the" {
			nameAndTitleSeparatorIndex = i
			break
		}
	}
	if nameAndTitleSeparatorIndex < 0 {
		return nil, errors.New("expected 'the' somewhere in name and title")
	}
	run.Name = strings.Trim(strings.Join(nameAndTitle[:nameAndTitleSeparatorIndex], ""), " ")
	run.Title = strings.Trim(strings.Join(nameAndTitle[nameAndTitleSeparatorIndex+2:], ""), " ")

	// Parse race.
	tok, lit = p.scan()
	if tok != WORD {
		return nil, errors.New("expected word when parsing race")
	}
	run.Race = lit
	switch lit {
	case "Deep", "High", "Hill", "Vine":
		tok, lit = p.scanIgnoreWhitespace()
		if tok != WORD {
			return nil, errors.New("expected second word when parsing race")
		}
		run.Race += " "
		run.Race += lit
	}

	// Parse background.
	tok, lit = p.scanIgnoreWhitespace()
	if tok != WORD {
		return nil, errors.New("expected word when parsing background")
	}
	run.Background = lit
	for tok, lit = p.scanIgnoreWhitespace(); tok == WORD; tok, lit = p.scanIgnoreWhitespace() {
		run.Background += " "
		run.Background += lit
	}
	if tok != EPAREN {
		return nil, errors.New("expected EPAREN when parsing background")
	}

	// Parse turns.
	tok, lit = p.scanNextTokenIgnoringOthers(nil, NUMBER)
	if tok == EOF {
		return nil, errors.New("EOF when parsing turns")
	}
	run.Turns, err = strconv.Atoi(lit)
	if err != nil {
		return nil, errors.New("could not parse turns int: " + err.Error())
	}

	// Parse time.
	tok, lit = p.scanNextTokenIgnoringOthers(nil, NUMBER)
	if tok == EOF {
		return nil, errors.New("EOF when parsing turns")
	}
	run.Time = lit
	if tok, lit = p.scan(); tok != COLON {
		return nil, errors.New("expected COLON when parsing time")
	}
	run.Time += ":"
	if tok, lit = p.scan(); tok != NUMBER {
		return nil, errors.New("expected second NUMBER when parsing time")
	}
	run.Time += lit
	if tok, lit = p.scan(); tok != COLON {
		return nil, errors.New("expected second COLON when parsing time")
	}
	run.Time += ":"
	if tok, lit = p.scan(); tok != NUMBER {
		return nil, errors.New("expected third NUMBER when parsing time")
	}
	run.Time += lit
	if tok, lit = p.scanNextNewline(); tok != NL && tok != MNL {
		return nil, errors.New("expected newline after time")
	}

	// Parse stats.
	currentKey := ""
	for {
		if currentKey == "" {
			tok, lit = p.scanIgnoreWhitespace()
			if tok != WORD {
				return nil, errors.New("expected WORD when parsing stats: " + lit)
			}
			currentKey = lit
			if tok, lit = p.scan(); tok != COLON {
				return nil, errors.New("expected COLON when parsing stats")
			}
		}
		currentAttributeLits := []string{}
		lastTok := EOF
		for {
			tok, lit = p.scan()
			if tok == EOF {
				return nil, errors.New("EOF when parsing stats")
			}
			if tok == MNL || tok == NL {
				run.Stats[currentKey] = strings.Trim(strings.Join(currentAttributeLits, ""), " ")
				currentKey = ""
				break
			}
			if tok == COLON && lastTok == WORD {
				newKey := currentAttributeLits[len(currentAttributeLits)-1]
				currentAttributeLits = currentAttributeLits[:len(currentAttributeLits)-1]
				run.Stats[currentKey] = strings.Trim(strings.Join(currentAttributeLits, ""), " ")
				currentKey = newKey
				break
			}
			currentAttributeLits = append(currentAttributeLits, lit)
			lastTok = tok
		}
		if tok == MNL {
			break
		}
	}

	// Parse resistances and equipped.
	currentKey = ""
	for {
		if currentKey == "" {
			tok, lit = p.scanIgnoreWhitespace()
			if tok == BPAREN {
				tok, lit = p.scanNextNewline()
				if tok == MNL || tok == EOF {
					break
				}
				continue
			}
			if tok != WORD {
				return nil, errors.New("expected WORD when parsing resistances: " + lit)
			}
			currentKey = lit
		}
		tok, lit = p.scanIgnoreWhitespace()
		if tok == EOF {
			return nil, errors.New("EOF when parsing resistances")
		}
		if tok == MINUS {
			tok, lit = p.scanNextNewline()
			if tok == EOF {
				return nil, errors.New("EOF when parsing equipped")
			}
			run.Equipped = append(run.Equipped, strings.Trim(lit, " "))
			if tok == MNL {
				break
			}
			currentKey = ""
			continue
		}
		currentValue := lit
		for {
			tok, lit = p.scan()
			if tok == EOF {
				return nil, errors.New("EOF when parsing resistances")
			}
			if tok == MNL || tok == NL {
				run.Resistances[currentKey] = strings.Trim(currentValue, " ")
				currentKey = ""
				break
			}
			if tok == WORD || tok == BPAREN {
				run.Resistances[currentKey] = strings.Trim(currentValue, " ")
				currentKey = ""
				p.unscan()
				break
			}
			currentValue += lit
		}
		if tok == MNL {
			break
		}
	}

	// Parse attributes.
	currentKey = ""
	for {
		if currentKey == "" {
			tok, lit = p.scan()
			if tok == EOF {
				return nil, errors.New("EOF when reading attributes")
			}
			currentKey = lit
			if tok, lit = p.scan(); tok != COLON {
				return nil, errors.New("expected COLON when parsing attributes")
			}
		}
		currentValue := []string{}
		currentAttributeLits := []string{}
		lastToks := []Token{NL, NL}
		for {
			tok, lit = p.scan()
			if tok == EOF {
				return nil, errors.New("EOF when parsing attributes")
			}
			if tok == MNL {
				currentValue = append(currentValue, strings.Trim(strings.Join(currentAttributeLits, ""), " "))
				run.Attributes[currentKey] = currentValue
				break
			}
			if tok == COLON && lastToks[1] == NL {
				newKey := currentAttributeLits[len(currentAttributeLits)-1]
				currentAttributeLits = currentAttributeLits[:len(currentAttributeLits)-1]
				currentValue = append(currentValue, strings.Trim(strings.Join(currentAttributeLits, ""), " "))
				run.Attributes[currentKey] = currentValue
				currentKey = newKey
				currentValue = []string{}
				break
			}
			if tok == NL {
				currentAttributeLits = append(currentAttributeLits, " ")
			} else if tok == COLON {
				currentValue = []string{}
				currentAttributeLits = []string{}
			} else if tok == COMMA {
				currentValue = append(currentValue, strings.Trim(strings.Join(currentAttributeLits, ""), " "))
				currentAttributeLits = []string{}
			} else {
				currentAttributeLits = append(currentAttributeLits, lit)
			}
			lastToks[1] = lastToks[0]
			lastToks[0] = tok
		}
		if tok == MNL {
			break
		}
	}

	// Parse escaped.
	for {
		tok, lit = p.scanNextNewline()
		if tok == EOF {
			return nil, errors.New("EOF when parsing escaped")
		}
		if lit == "   Skills:" {
			break
		}
		if lit == "You escaped." {
			run.Escaped = true
		}
	}

	// Parse skills.
	for {
		tok, lit = p.scanIgnoreWhitespace()
		if tok == EOF {
			return nil, errors.New("EOF when parsing skills")
		}
		if tok == NL || tok == MNL {
			break
		}

		skill := Skill{}
		switch tok {
		case WORD:
			switch lit {
			case "O":
				skill.State = MASTERED
			case "Level":
				skill.State = UNUSED
				p.unscan()
			default:
				return nil, errors.New("unexpected word when parsing skills: " + lit)
			}
		case MINUS:
			skill.State = USING
		case PLUS:
			skill.State = TRAINING
		case ASTERISK:
			skill.State = FOCUSING
		default:
			return nil, errors.New("unexpected token when parsing skills: " + lit)
		}
		tok, lit = p.scanIgnoreWhitespace() // Discard "Level"
		if lit != "Level" {
			return nil, errors.New("unexpted literal when parsing skills: " + lit)
		}
		tok, lit = p.scanIgnoreWhitespace()
		if tok != NUMBER {
			return nil, errors.New("expected NUMBER when parsing skills: " + lit)
		}

		skill.Level, err = strconv.Atoi(lit)
		if err != nil {
			return nil, errors.New("failed to parse skill level: " + err.Error())
		}
		tok, lit = p.scan()
		if tok == DOT {
			tok, lit = p.scan()
			if tok != NUMBER {
				return nil, errors.New("expected NUMBER when parsing skills decimal: " + lit)
			}
			skill.LevelDecimal, err = strconv.Atoi(lit)
			if err != nil {
				return nil, errors.New("failed to parse skill level decimal: " + err.Error())
			}
			tok, lit = p.scan()
		}
		if tok == BPAREN {
			tok, lit = p.scan()
			if tok != NUMBER {
				return nil, errors.New("expected NUMBER when parsing crosstrained skill level: " + lit)
			}
			skill.Level, err = strconv.Atoi(lit)
			if err != nil {
				return nil, errors.New("failed to parse crosstrained skill level: " + err.Error())
			}
			tok, lit = p.scan()
			if tok == DOT {
				tok, lit = p.scan()
				if tok != NUMBER {
					return nil, errors.New("expected NUMBER when parsing crosstrained skills decimal: " + lit)
				}
				skill.LevelDecimal, err = strconv.Atoi(lit)
				if err != nil {
					return nil, errors.New("failed to parse crosstrained skill level decimal: " + err.Error())
				}
				tok, lit = p.scan()
			} else {
				skill.LevelDecimal = 0
			}
			if tok != EPAREN {
				return nil, errors.New("expected EPAREN when parsing crosstrained skills decimal: " + lit)
			}
			tok, lit = p.scan()
		}
		if tok != WS {
			return nil, errors.New("expected WS when parsing skills: " + lit)
		}
		tok, lit = p.scanNextNewline()
		skill.Name = lit
		run.Skills = append(run.Skills, skill)
		if tok == MNL {
			break
		}
	}

	// Parse spells.
	hasSpells := false
	for {
		tok, lit = p.scanNextNewline()
		if tok == EOF {
			return nil, errors.New("EOF when searching for spells")
		}
		if lit == "You didn't know any spells." {
			break
		}
		if lit == "You knew the following spells:" {
			hasSpells = true
			p.scanNextNewline() // Discard headers
			break
		}
	}
	for hasSpells {
		tok, lit = p.scan()
		if tok != WORD {
			return nil, errors.New("expected word when parsing spells: " + lit)
		}
		tok, lit = p.scanIgnoreWhitespace()
		if tok != MINUS {
			return nil, errors.New("expected MINUS when parsing spells: " + lit)
		}
		lits := []string{}
		for {
			tok, lit = p.scan()
			if tok == EOF {
				return nil, errors.New("EOF when parsing spells")
			}
			if tok == WORD && len(lits) > 0 && strings.HasSuffix(lits[len(lits)-1], "  ") {
				p.unscan()
				break
			}
			lits = append(lits, lit)
		}
		spell := Spell{
			Name: strings.Trim(strings.Join(lits, ""), " "),
		}
		for {
			tok, lit = p.scan()
			if tok == EOF {
				return nil, errors.New("EOF when parsing spells")
			}
			if tok == WS {
				break
			}
			spell.Type += lit
		}
		for {
			tok, lit = p.scan()
			if tok == EOF {
				return nil, errors.New("EOF when parsing spells")
			}
			if tok == WS {
				break
			}
			spell.Power += lit
		}
		tok, lit = p.scan()
		if tok != NUMBER {
			return nil, errors.New("expected NUMBER when parsing spell failure: " + lit)
		}
		spell.Failure, err = strconv.Atoi(lit)
		if err != nil {
			return nil, errors.New("failed to parse spell failure: " + err.Error())
		}
		tok, lit = p.scan()
		if tok != PERCENT {
			return nil, errors.New("expected PERCENT when parsing spell failure: " + lit)
		}
		tok, lit = p.scanIgnoreWhitespace()
		if tok != NUMBER {
			return nil, errors.New("expected NUMBER when parsing spell level: " + lit)
		}
		spell.Level, err = strconv.Atoi(lit)
		if err != nil {
			return nil, errors.New("failed to parse spell level: " + err.Error())
		}
		tok, lit = p.scan()
		if tok != WS {
			return nil, errors.New("expected WS when parsing spell level: " + lit)
		}
		tok, lit = p.scanNextNewline()
		spell.Hunger = lit
		run.Spells = append(run.Spells, spell)
		if tok == MNL {
			break
		}
	}
	for {
		tok, lit = p.scanNextNewline()
		if tok == EOF {
			break
		}
		if lit == "Innate Abilities, Weirdness & Mutations" {
			for {
				tok, lit = p.scanNextNewline()
				if tok == EOF {
					break
				}
				run.Mutations = append(run.Mutations, lit)
				if tok == MNL {
					break
				}
			}
		}
	}
	return run, nil
}
