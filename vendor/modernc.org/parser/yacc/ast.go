// CAUTION: Generated by yy - DO NOT EDIT.

// Copyright 2015 The parser Authors. All rights reserved.  Use
// of this source code is governed by a BSD-style license that can be found in
// the LICENSE file.
//
// This is a derived work base on the original at
//
// http://pubs.opengroup.org/onlinepubs/009695399/utilities/yacc.html
//
// The original work is
//
// Copyright © 2001-2004 The IEEE and The Open Group, All Rights reserved.
//
// Grammar for the input to yacc.

package parser // import "modernc.org/parser/yacc"

import (
	"go/token"
)

// Action represents data reduced by production:
//
//	Action:
//	        '{' '}'
type Action struct {
	Values []*ActionValue // For backward compatibility.
	Token  *Token
	Token2 *Token
}

func (n *Action) fragment() interface{} { return n }

// String implements fmt.Stringer.
func (n *Action) String() string {
	return prettyString(n)
}

// Pos reports the position of the first component of n or zero if it's empty.
func (n *Action) Pos() token.Pos {
	return n.Token.Pos()
}

// Definition represents data reduced by productions:
//
//	Definition:
//	        START IDENTIFIER
//	|       UNION                      // Case 1
//	|       LCURL RCURL                // Case 2
//	|       ReservedWord Tag NameList  // Case 3
//	|       ReservedWord Tag           // Case 4
//	|       ERROR_VERBOSE              // Case 5
type Definition struct {
	Nlist        []*Name // For backward compatibility.
	Value        string
	Case         int
	NameList     *NameList
	ReservedWord *ReservedWord
	Tag          *Tag
	Token        *Token
	Token2       *Token
}

func (n *Definition) fragment() interface{} { return n }

// String implements fmt.Stringer.
func (n *Definition) String() string {
	return prettyString(n)
}

// Pos reports the position of the first component of n or zero if it's empty.
func (n *Definition) Pos() token.Pos {
	switch n.Case {
	case 3, 4:
		return n.ReservedWord.Pos()
	case 0, 1, 2, 5:
		return n.Token.Pos()
	default:
		panic("internal error")
	}
}

// DefinitionList represents data reduced by productions:
//
//	DefinitionList:
//	        /* empty */
//	|       DefinitionList Definition  // Case 1
type DefinitionList struct {
	Definition     *Definition
	DefinitionList *DefinitionList
}

func (n *DefinitionList) reverse() *DefinitionList {
	if n == nil {
		return nil
	}

	na := n
	nb := na.DefinitionList
	for nb != nil {
		nc := nb.DefinitionList
		nb.DefinitionList = na
		na = nb
		nb = nc
	}
	n.DefinitionList = nil
	return na
}

func (n *DefinitionList) fragment() interface{} { return n.reverse() }

// String implements fmt.Stringer.
func (n *DefinitionList) String() string {
	return prettyString(n)
}

// Pos reports the position of the first component of n or zero if it's empty.
func (n *DefinitionList) Pos() token.Pos {
	if n == nil {
		return 0
	}

	if p := n.DefinitionList.Pos(); p != 0 {
		return p
	}

	return n.Definition.Pos()
}

// LiteralStringOpt represents data reduced by productions:
//
//	LiteralStringOpt:
//	        /* empty */
//	|       STRING_LITERAL  // Case 1
type LiteralStringOpt struct {
	Token *Token
}

func (n *LiteralStringOpt) fragment() interface{} { return n }

// String implements fmt.Stringer.
func (n *LiteralStringOpt) String() string {
	return prettyString(n)
}

// Pos reports the position of the first component of n or zero if it's empty.
func (n *LiteralStringOpt) Pos() token.Pos {
	if n == nil {
		return 0
	}

	return n.Token.Pos()
}

// Name represents data reduced by productions:
//
//	Name:
//	        IDENTIFIER LiteralStringOpt
//	|       IDENTIFIER NUMBER LiteralStringOpt  // Case 1
type Name struct {
	Identifier       interface{} // For backward compatibility.
	Number           int         // For backward compatibility.
	Case             int
	LiteralStringOpt *LiteralStringOpt
	Token            *Token
	Token2           *Token
}

func (n *Name) fragment() interface{} { return n }

// String implements fmt.Stringer.
func (n *Name) String() string {
	return prettyString(n)
}

// Pos reports the position of the first component of n or zero if it's empty.
func (n *Name) Pos() token.Pos {
	return n.Token.Pos()
}

// NameList represents data reduced by productions:
//
//	NameList:
//	        Name
//	|       NameList Name      // Case 1
//	|       NameList ',' Name  // Case 2
type NameList struct {
	Case     int
	Name     *Name
	NameList *NameList
	Token    *Token
}

func (n *NameList) reverse() *NameList {
	if n == nil {
		return nil
	}

	na := n
	nb := na.NameList
	for nb != nil {
		nc := nb.NameList
		nb.NameList = na
		na = nb
		nb = nc
	}
	n.NameList = nil
	return na
}

func (n *NameList) fragment() interface{} { return n.reverse() }

// String implements fmt.Stringer.
func (n *NameList) String() string {
	return prettyString(n)
}

// Pos reports the position of the first component of n or zero if it's empty.
func (n *NameList) Pos() token.Pos {
	switch n.Case {
	case 0:
		return n.Name.Pos()
	case 1, 2:
		return n.NameList.Pos()
	default:
		panic("internal error")
	}
}

// Precedence represents data reduced by productions:
//
//	Precedence:
//	        /* empty */
//	|       PREC IDENTIFIER         // Case 1
//	|       PREC IDENTIFIER Action  // Case 2
//	|       Precedence ';'          // Case 3
type Precedence struct {
	Identifier interface{} // Name string or literal int.
	Action     *Action
	Case       int
	Precedence *Precedence
	Token      *Token
	Token2     *Token
}

func (n *Precedence) fragment() interface{} { return n }

// String implements fmt.Stringer.
func (n *Precedence) String() string {
	return prettyString(n)
}

// Pos reports the position of the first component of n or zero if it's empty.
func (n *Precedence) Pos() token.Pos {
	switch n.Case {
	case 0:
		return 0
	case 3:
		if p := n.Precedence.Pos(); p != 0 {
			return p
		}

		return n.Token.Pos()
	case 1, 2:
		return n.Token.Pos()
	default:
		panic("internal error")
	}
}

// ReservedWord represents data reduced by productions:
//
//	ReservedWord:
//	        TOKEN
//	|       LEFT        // Case 1
//	|       RIGHT       // Case 2
//	|       NONASSOC    // Case 3
//	|       TYPE        // Case 4
//	|       PRECEDENCE  // Case 5
type ReservedWord struct {
	Case  int
	Token *Token
}

func (n *ReservedWord) fragment() interface{} { return n }

// String implements fmt.Stringer.
func (n *ReservedWord) String() string {
	return prettyString(n)
}

// Pos reports the position of the first component of n or zero if it's empty.
func (n *ReservedWord) Pos() token.Pos {
	return n.Token.Pos()
}

// Rule represents data reduced by productions:
//
//	Rule:
//	        C_IDENTIFIER RuleItemList Precedence
//	|       '|' RuleItemList Precedence           // Case 1
type Rule struct {
	Body         []interface{} // For backward compatibility.
	Name         *Token
	Case         int
	Precedence   *Precedence
	RuleItemList *RuleItemList
	Token        *Token
}

func (n *Rule) fragment() interface{} { return n }

// String implements fmt.Stringer.
func (n *Rule) String() string {
	return prettyString(n)
}

// Pos reports the position of the first component of n or zero if it's empty.
func (n *Rule) Pos() token.Pos {
	return n.Token.Pos()
}

// RuleItemList represents data reduced by productions:
//
//	RuleItemList:
//	        /* empty */
//	|       RuleItemList IDENTIFIER      // Case 1
//	|       RuleItemList Action          // Case 2
//	|       RuleItemList STRING_LITERAL  // Case 3
type RuleItemList struct {
	Action       *Action
	Case         int
	RuleItemList *RuleItemList
	Token        *Token
}

func (n *RuleItemList) reverse() *RuleItemList {
	if n == nil {
		return nil
	}

	na := n
	nb := na.RuleItemList
	for nb != nil {
		nc := nb.RuleItemList
		nb.RuleItemList = na
		na = nb
		nb = nc
	}
	n.RuleItemList = nil
	return na
}

func (n *RuleItemList) fragment() interface{} { return n.reverse() }

// String implements fmt.Stringer.
func (n *RuleItemList) String() string {
	return prettyString(n)
}

// Pos reports the position of the first component of n or zero if it's empty.
func (n *RuleItemList) Pos() token.Pos {
	switch n.Case {
	case 0:
		return 0
	case 2:
		if p := n.RuleItemList.Pos(); p != 0 {
			return p
		}

		return n.Action.Pos()
	case 1, 3:
		if p := n.RuleItemList.Pos(); p != 0 {
			return p
		}

		return n.Token.Pos()
	default:
		panic("internal error")
	}
}

// RuleList represents data reduced by productions:
//
//	RuleList:
//	        C_IDENTIFIER RuleItemList Precedence
//	|       RuleList Rule                         // Case 1
type RuleList struct {
	Case         int
	Precedence   *Precedence
	Rule         *Rule
	RuleItemList *RuleItemList
	RuleList     *RuleList
	Token        *Token
}

func (n *RuleList) reverse() *RuleList {
	if n == nil {
		return nil
	}

	na := n
	nb := na.RuleList
	for nb != nil {
		nc := nb.RuleList
		nb.RuleList = na
		na = nb
		nb = nc
	}
	n.RuleList = nil
	return na
}

func (n *RuleList) fragment() interface{} { return n.reverse() }

// String implements fmt.Stringer.
func (n *RuleList) String() string {
	return prettyString(n)
}

// Pos reports the position of the first component of n or zero if it's empty.
func (n *RuleList) Pos() token.Pos {
	switch n.Case {
	case 1:
		return n.RuleList.Pos()
	case 0:
		return n.Token.Pos()
	default:
		panic("internal error")
	}
}

// Specification represents data reduced by production:
//
//	Specification:
//	        DefinitionList "%%" RuleList Tail
type Specification struct {
	Defs           []*Definition // For backward compatibility.
	Rules          []*Rule       // For backward compatibility.
	DefinitionList *DefinitionList
	RuleList       *RuleList
	Tail           *Tail
	Token          *Token
}

func (n *Specification) fragment() interface{} { return n }

// String implements fmt.Stringer.
func (n *Specification) String() string {
	return prettyString(n)
}

// Pos reports the position of the first component of n or zero if it's empty.
func (n *Specification) Pos() token.Pos {
	if p := n.DefinitionList.Pos(); p != 0 {
		return p
	}

	return n.Token.Pos()
}

// Tag represents data reduced by productions:
//
//	Tag:
//	        /* empty */
//	|       '<' IDENTIFIER '>'  // Case 1
type Tag struct {
	Token  *Token
	Token2 *Token
	Token3 *Token
}

func (n *Tag) fragment() interface{} { return n }

// String implements fmt.Stringer.
func (n *Tag) String() string {
	return prettyString(n)
}

// Pos reports the position of the first component of n or zero if it's empty.
func (n *Tag) Pos() token.Pos {
	if n == nil {
		return 0
	}

	return n.Token.Pos()
}

// Tail represents data reduced by productions:
//
//	Tail:
//	        "%%"
//	|       /* empty */  // Case 1
type Tail struct {
	Value string
	Token *Token
}

func (n *Tail) fragment() interface{} { return n }

// String implements fmt.Stringer.
func (n *Tail) String() string {
	return prettyString(n)
}

// Pos reports the position of the first component of n or zero if it's empty.
func (n *Tail) Pos() token.Pos {
	if n == nil {
		return 0
	}

	return n.Token.Pos()
}
