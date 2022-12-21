package testtime

import (
	"testing"
	"time"

	"github.com/tenntenn/testtime"
)

func Test(t *testing.T) {

	t.Run("A", func(t *testing.T) {
		// set zero value
		testtime.SetTime(t, time.Time{})
		// true
		if time.Now().IsZero() {
			t.Error("error")
		}
	})

	t.Run("B", func(t *testing.T) {
		// set func which return zero value
		f := func() time.Time {
			return time.Time{}
		}
		testtime.SetFunc(t, f)
		// true
		if time.Now().IsZero() {
			t.Error("error")
		}
	})

	// false
	if !time.Now().IsZero() {
		t.Error("error")
	}
}
