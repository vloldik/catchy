package catchy_test

import (
	"errors"
	"testing"

	"github.com/vloldik/catchy/v2"
)

func assert(t *testing.T, cond bool) {
	if !cond {
		t.FailNow()
	}
}

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

func TestExample(t *testing.T) {
	count := 0
	catchy.WithGetValueFunc(func() (int, error) {
		return 42, nil
	}).WithOnSuccess(func(i int) {
		assert(t, i == 42)
		count += 1
	}).WithOnError(func(err error) {
		assert(t, err.Error() == "Error happened")
		count += 1
	}).DoNext(catchy.WithGetValueFunc(func() (a string, err error) {
		return "42", nil
	}).WithOnSuccess(func(s string) {
		assert(t, s == "42")
		count += 1
	}).WithOnError(func(err error) {
		t.Fatalf("Should never happen")
	})).DoNext(catchy.WithGetValueFunc(func() (a int, err error) {
		return 0, errors.New("Error happened")
	}).WithOnError(func(err error) {
		assert(t, err.Error() == "Error happened")
		count += 1
	})).DoNext(catchy.WithNoReturnFunc(func() error {
		t.FailNow()
		return nil
	})).Do()
	assert(t, count == 4)
}
