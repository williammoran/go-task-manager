package tasks

import "testing"

func TestFixMeDoesNotFire(t *testing.T) {
	FixMe("9999-01-01 Incorrect panic")
}

func TestFixMePanicInvalidDate(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("The code did not panic")
		}
	}()
	FixMe("Jan 20 Should Panic")
}

func TestFixMePanicAfterDate(t *testing.T) {
	const panicStr = "0001-01-01 Should Panic"
	defer func() {
		r := recover()
		if r == nil {
			t.Fatalf("The code did not panic")
		}
		if r != panicStr {
			t.Fatalf("Wrong message: '%s'", r)
		}
	}()
	FixMe("0001-01-01 Should Panic")
}
