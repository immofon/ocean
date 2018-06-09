package token

import "fmt"

type Pos int

func (p Pos) String() string {
	return fmt.Sprint(p)
}

type Type int

const (
	T_Unknown Type = 0
	T_Left_P       // (
	T_Right_P      // )
	T_Number       // [0-9]+
	T_String       //
	T_Ident        // others
)

type Token interface {
	Type() Type
	Pos() Pos
	String() string // <token:%pos:%type:%details>
}

type NumberToken struct {
	pos    Pos
	number string
}

func Number(pos Pos, num string) NumberToken {
	return NumberToken{
		pos:    pos,
		number: num,
	}
}

func (t *NumberToken) Type() Type {
	return T_Number
}

func (t *NumberToken) Pos() Pos {
	return t.pos
}

func (t *NumberToken) Number() string {
	return t.number
}

func (t *NumberToken) String() string {
	return fmt.Sprintf("<token:%s:number: %s>", t.Pos, t.number)
}
