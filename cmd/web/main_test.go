package main

import "testing"

func TestRun(t *testing.T) {
	//t.Parallel()
	_, err := run()
	if err != nil {
		t.Error("failed run()")
	}
}
