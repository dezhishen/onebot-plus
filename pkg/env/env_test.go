package env

import "testing"

func TestSetEnv(t *testing.T) {
	SetEnv("foo", "bar")
}
