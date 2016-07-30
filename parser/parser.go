package parser

import (
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

// Parse parses a file
func (p *Parser) Parse() (string, error) {
	file := ""

	tcount := 0
	for {
		// Read a field.
		tok, lit := p.scan()
		if tcount > 0 && tcount != 3 {
			if tok != TICK {
				tcount = 0
			}
		}

		if tok == TICK {
			tcount++
			continue
		} else {
			if tcount == 3 {
				file += lit
			} else if tcount > 3 {
				tcount = 0
			}
		}

		if tok == 1 {
			break
		}
	}

	return file, nil
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
