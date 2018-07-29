package config

import (
	"os"
	"testing"
)

func TestGetBooleanConfiguration(t *testing.T) {
	result := GetBooleanConfiguration("WRONG_CONFIG")
	if result {
		t.Fatal("WRONG_CONFIG is true, should be false")
	}

	os.Setenv("FOO", "1")
	result = GetBooleanConfiguration("FOO")
	if result {
		t.Fatal("FOO is true, should be false")
	}

	os.Setenv("FOO", "true")
	result = GetBooleanConfiguration("FOO")
	if !result {
		t.Fatal("FOO is false, should be true")
	}
}
