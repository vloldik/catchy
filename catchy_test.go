package catchy_test

import (
	"errors"
	"testing"

	"github.com/vloldik/catchy"
)

func SomeFuncThatReturnsError() (a struct{}, b error) {
	return a, errors.New("some error")
}

func SomeFuncThatReturnsValue() (a struct{}, b error) {
	return a, nil
}

func SomeFuncThatReturnsOnlyError() error {
	return errors.New("some error")
}

func SomeFuncThatReturnsOnlyNil() error {
	return nil
}

func TestCatchy(t *testing.T) {
	catchy.WithGetValueFunc(SomeFuncThatReturnsError).WithOnError(func(err error) {
		t.Log(err)
	}).WithOnSuccess(func(s struct{}) {
		t.Error("should not be called")
	}).Do()
	catchy.WithGetValueFunc(SomeFuncThatReturnsValue).WithOnError(func(err error) {
		t.Error("should not be called")
	}).WithOnSuccess(func(s struct{}) {
		t.Log(s)
	}).Do()
	catchy.WithNoReturnFunc(SomeFuncThatReturnsOnlyError).WithOnError(func(err error) {
		t.Log(err)
	}).WithOnSuccess(func(_ catchy.Never) {
		t.Error("should not be called")
	}).Do()
	catchy.WithNoReturnFunc(SomeFuncThatReturnsOnlyNil).WithOnError(func(err error) {
		t.Error("should not be called")
	}).WithOnSuccess(func(_ catchy.Never) {
	}).Do()
}

func RequirePanic(fn func()) {
	defer func() {
		if r := recover(); r == nil {
			panic("did not panic")
		}
	}()
	fn()
}

func TestMust(t *testing.T) {
	RequirePanic(func() { catchy.Must(SomeFuncThatReturnsError()) })
	catchy.Must(SomeFuncThatReturnsValue())
	RequirePanic(func() { catchy.MustNoReturn(SomeFuncThatReturnsOnlyError()) })
	catchy.MustNoReturn(SomeFuncThatReturnsOnlyNil())
}
