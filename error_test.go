package errors

import (
	"testing"
)

func TestNew(t *testing.T) {
	e := New("code", "message")
	if e == nil {
		t.Errorf("Expects e is not nil")
	}
}

func TestFactoryError_Code(t *testing.T) {
	e := &FactoryError{code: "this_is_code"}
	if e.Code() != "this_is_code" {
		t.Errorf("Expects %s. Got %s", e.code, e.Code())
	}
}

func TestFactoryError_WithCode(t *testing.T) {
	e := &FactoryError{}
	e.WithCode("a_code")
	if e.code != "a_code" {
		t.Errorf("Expects %s. Got %s", "a_code", e.code)
	}
}

func TestFactoryError_Message(t *testing.T) {
	e := &FactoryError{message: "this_is_msg"}
	if e.Message() != "this_is_msg" {
		t.Errorf("Expects %s. Got %s", e.message, e.Message())
	}
}

func TestFactoryError_WithMessage(t *testing.T) {
	e := &FactoryError{}
	e.WithMessage("a_message")
	if e.message != "a_message" {
		t.Errorf("Expects %s. Got %s", "a_message", e.message)
	}
}

func TestFactoryError_Error(t *testing.T) {
	e := &FactoryError{code: "code", message: "message", stack: nil}
	if e.Error() != "[code] message" {
		t.Errorf("Expects %s. Got %s", "[code] message", e.Error())
	}
}

func TestFactoryError_TraceString(t *testing.T) {
	e := New("c", "m")
	if e.TraceString() == "" {
		t.Error("Expects e is not empty")
	}
}

func TestFactoryError_Trace(t *testing.T) {
	e := New("c", "m")
	e.Trace()
}

func TestFactoryError_Debug(t *testing.T) {
	e := &FactoryError{debug: "this_is_code"}
	if e.Debug() != "this_is_code" {
		t.Errorf("Expects %s. Got %s", e.debug, e.Debug())
	}
}

func TestFactoryError_WithDebug(t *testing.T) {
	e := &FactoryError{}
	e.WithDebug("a_code")
	if e.debug != "a_code" {
		t.Errorf("Expects %s. Got %s", "a_code", e.debug)
	}
}
