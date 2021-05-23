package chat

import "testing"

func TestSay(t *testing.T) {
	output := Say("%s", "testing")

	if output != "testing" {
		t.Logf("output does not have the correct value. Got: %s, Want: testing", output)
		t.Fail()
	}
}
