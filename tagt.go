package tagt

import (
	"fmt"
	"testing"
)

// T prints tags when test funcs are called.
type T struct {
	t    *testing.T
	tags []string
}

// New T from testing.T and tags.
func New(t *testing.T, tags ...string) T {
	return T{
		t:    t,
		tags: tags,
	}
}

// Error prints tags with original function.
func (t T) Error(args ...interface{}) {
	t.t.Helper()
	t.t.Error(t.newArgs(args)...)
}

// Errorf prints tags with original function.
func (t T) Errorf(format string, args ...interface{}) {
	t.t.Helper()
	if len(args) == 0 {
		t.Error(format)
	} else {
		t.t.Error(t.tag(), fmt.Sprintf(format, args))
	}
}

// Fail delegates original function.
func (t T) Fail() {
	t.t.Helper()
	t.t.Fail()
}

// FailNow delegates original function.
func (t T) FailNow() {
	t.t.Helper()
	t.t.FailNow()
}

// Failed delegates original function.
func (t T) Failed() bool {
	return t.t.Failed()
}

// Fatal prints tags with original function.
func (t T) Fatal(args ...interface{}) {
	t.t.Helper()
	t.t.Fatal(t.newArgs(args)...)
}

// Fatalf prints tags with original function.
func (t T) Fatalf(format string, args ...interface{}) {
	t.t.Helper()
	if len(args) == 0 {
		t.Fatal(format)
	} else {
		t.t.Fatal(t.tag(), fmt.Sprintf(format, args))
	}
}

// Log prints tags with original function.
func (t T) Log(args ...interface{}) {
	t.t.Helper()
	t.t.Log(t.newArgs(args)...)
}

// Logf prints tags with original function.
func (t T) Logf(format string, args ...interface{}) {
	t.t.Helper()
	if len(args) == 0 {
		t.Log(format)
	} else {
		t.t.Log(t.tag(), fmt.Sprintf(format, args))
	}
}

// Name delegates original function.
func (t T) Name() string {
	return t.t.Name()
}

// Skip prints tags with original function.
func (t T) Skip(args ...interface{}) {
	t.t.Helper()
	t.t.Skip(t.newArgs(args)...)
}

// SkipNow delegates original function.
func (t T) SkipNow() {
	t.t.Helper()
	t.SkipNow()
}

// Skipf prints tags with original function.
func (t T) Skipf(format string, args ...interface{}) {
	t.t.Helper()
	if len(args) == 0 {
		t.Skip(format)
	} else {
		t.t.Skip(t.tag(), fmt.Sprintf(format, args))
	}
}

// Skipped delegates original function.
func (t T) Skipped() bool {
	return t.t.Skipped()
}

func (t T) tag() string {
	tg := ""
	for _, v := range t.tags {
		tg += "[" + v + "]"
	}
	return tg
}

func (t T) newArgs(args []interface{}) []interface{} {
	ret := make([]interface{}, len(args)+1)
	ret[0] = t.tag()
	for i, v := range args {
		ret[i+1] = v
	}
	return ret
}
