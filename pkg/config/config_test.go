package config

import (
	"os"
	"testing"
)

func TestGetBooleanConfiguration(t *testing.T) {
	result := GetBooleanConfiguration("NOVALUE_CONFIG")
	if result {
		t.Fatal("NOVALUE_CONFIG is true, should be false")
	}

	os.Setenv("NOTTRUE_CONFIG", "1")
	result = GetBooleanConfiguration("NOTTRUE_CONFIG")
	if result {
		t.Fatal("NOTTRUE_CONFIG is true, should be false")
	}

	os.Setenv("ISTRUE_CONFIG", "true")
	result = GetBooleanConfiguration("ISTRUE_CONFIG")
	if !result {
		t.Fatal("ISTRUE_CONFIG is false, should be true")
	}
}

func TestGetStringConfiguration(t *testing.T) {
	result := GetStringConfiguration("NOVALUE_CONFIG")
	if result != "" {
		t.Fatal("NOVALUE_CONFIG is true, should be false")
	}

	os.Setenv("NOTTRUE_CONFIG", "somevalue")
	result = GetStringConfiguration("NOTTRUE_CONFIG")
	if result != "somevalue" {
		t.Fatal("NOTTRUE_CONFIG is true, should be false")
	}

	os.Setenv("ISTRUE_CONFIG", "true")
	result = GetStringConfiguration("ISTRUE_CONFIG")
	if result != "true" {
		t.Fatal("ISTRUE_CONFIG is false, should be true")
	}
}
