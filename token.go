package main

type Token int

const (
	ILLEGAL Token = iota
	EOF
	WS

	IDENT
	TICK
)
