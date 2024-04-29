package mocking

import (
	"bytes"
	"testing"
)

type SpySleep struct{
	Calls int
}

func (s *SpySleep)Sleep(){
	s.Calls++
}

func TestCountDown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleep{}

	CountDown(buffer, spySleeper)
	got := buffer.String()
	want := "3\n2\n1\nGo!\n"

	if got != want{
		t.Errorf("got %q, want %q", got, want)
	}

	if spySleeper.Calls != 3{
		t.Errorf("not enough calls to sleeper, want 3 got %d", spySleeper.Calls)
	}
}