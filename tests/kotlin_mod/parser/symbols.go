// Code generated by daniil-ushkov LL(1) parser generator; DO NOT EDIT.
package parser

//go:generate stringer -type=Terminal
type Terminal int

const (
	LP Terminal = iota
	RP
	EOF
	FUN
	ID
	COL
	COM
)

//go:generate stringer -type=NonTerminal
type NonTerminal int

const (
	Ret NonTerminal = iota
	Sig
	Params
	Param
	Params1
	Params2
)