package parser

type Token int

const (
	ILLEGAL Token = iota

	EOF   // end of file
	WS    // whitespace
	EOL   // new line
	IDENT // anything else
	TICK  // backticks
)
