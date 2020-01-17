package base

import "fmt"

const prefix = "userDetail"

var (
	ErrFormParamInValid = NewErr(1, "form param invalid", prefix)
)

// Error translate
type Error struct {
	Code   int32
	Detail string
	Prefix string
	//
	vars  map[string][]interface{}
	isExt bool
}

func (e *Error) GetPrefix() string {
	return e.Prefix
}

func (e *Error) GetCode() int32 {
	return e.Code
}

func (e *Error) GetKey() string {
	return fmt.Sprintf("%s%d", e.GetPrefix(), e.GetCode())
}

func (e *Error) GetDetail() string {
	return e.Detail
}

func (e *Error) GetVars() []interface{} {
	return e.GetVarsBy("")
}

func (e *Error) GetVarsBy(lang string) []interface{} {
	return e.vars[lang]
}

func (e *Error) SetDetail(c string) *Error {
	e.Detail = c
	return e
}

func (e *Error) SetVars(con ...interface{}) *Error {
	return e.SetVarsBy("", con...)
}

func (e *Error) SetVarsBy(lang string, con ...interface{}) *Error {
	var e2 *Error
	if e.isExt {
		e2 = e
	} else {
		e2 = NewErr(e.Code, e.Detail, e.Prefix)
		e2.isExt = true
	}
	e2.vars[lang] = con
	return e2
}

func (e *Error) Error() string {
	return e.GetKey()
}

func NewErr(c int32, t, m string) *Error {
	e := &Error{Code: c, Detail: t, Prefix: m}
	e.vars = make(map[string][]interface{})
	return e
}

// Text translate
type Text struct {
	Code   int32
	Detail string
	Prefix string
	vars   []interface{}
}

func (e *Text) GetPrefix() string {
	return e.Prefix
}

func (e *Text) GetCode() int32 {
	return e.Code
}

func (e *Text) GetKey() string {
	return fmt.Sprintf("%s%d", e.GetPrefix(), e.GetCode())
}

func (e *Text) GetDetail() string {
	return e.Detail
}

func (e *Text) GetVars() []interface{} {
	return e.vars
}

func (e *Text) SetDetail(c string) *Text {
	e.Detail = c
	return e
}

func (e *Text) SetVars(con ...interface{}) *Text {
	e2 := &Text{}
	*e2 = *e
	e2.vars = con
	return e2
}

func (e *Text) String() string {
	return fmt.Sprintf("%s%d|%s", e.Prefix, e.Code, e.Detail)
}
