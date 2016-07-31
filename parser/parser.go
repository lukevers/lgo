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

	// The variable `tcount` is used for keeping track of the number of back
	// ticks (`) in a row that we've hit. The back ticks are for spotting the
	// code we actually want, like the following:
	//
	// ```
	// // some code here...
	// ```
	tcount := 0

	// The variable `teol` is used for keeping track of the EOL on the same line
	// as the triple back ticks. We want to be able to specify a language type,
	// if the user desires, on the back ticks like the following:
	//
	// ```go
	// // some code here...
	// ```
	teol := false

	for {
		// Read a field.
		tok, lit := p.scan()
		if tcount > 0 && tcount != 3 {
			if tok != TICK {
				tcount = 0
				teol = false
			}
		}

		if tok == TICK {
			tcount++
			continue
		} else {
			if tcount == 3 {
				// Now if the character hasn't hit the EOL yet, wait for the
				// EOL before we allow the characters to be added to the file.
				if !teol && tok == EOL {
					teol = true
				}

				// If we've hit the point where we're not on the backticks line
				// anymore, we can add the character.
				if teol {
					file += lit
				}
			} else if tcount > 3 {
				tcount = 0
				teol = false
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
