package config

import (
	"os"
	"strconv"
	"testing"
)

func TestGetEnvIntValid(t *testing.T) {
	testNumber := 450
	os.Setenv("TEST_INT", strconv.Itoa(testNumber))
	testValue := getEnvInt("TEST_INT", 0)
	if testNumber != testValue {
		t.Errorf("Environment number returned %d does not match the set value %d", testValue, testNumber)
	}
}

func TestGetEnvIntNotSet(t *testing.T) {
	os.Unsetenv("TEST_INT")
	testValue := getEnvInt("TEST_INT", 0)
	if 0 != testValue {
		t.Errorf("Environment number returned %d does not match the default value %d", testValue, 0)
	}
}

func TestGetEnvStringValid(t *testing.T) {
	testString := "test environment variable"
	os.Setenv("TEST_STRING", testString)
	testValue := getEnvString("TEST_STRING", "")
	if testString != testValue {
		t.Errorf("Environment number returned %s does not match the set value %s", testValue, testString)
	}
}

func TestGetEnvStringNotSet(t *testing.T) {
	os.Unsetenv("TEST_STRING")
	defaultString := "default string"
	testValue := getEnvString("TEST_STRING", defaultString)
	if defaultString != testValue {
		t.Errorf("Environment number returned %s does not match the set value %s", testValue, defaultString)
	}
}
