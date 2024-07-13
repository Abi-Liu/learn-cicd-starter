package auth

import (
	"errors"
	"net/http"
	"reflect"
	"testing"
)

func TestAuth(t *testing.T) {
	type res struct {
		key string
		err error
	}

	type test struct {
		name     string
		expected res
		input    http.Header
	}

	badHeader := http.Header{}
	badHeader.Add("Authorization", "Bearer asdf")
	goodHeader := http.Header{}
	goodHeader.Add("Authorization", "ApiKey test")

	tests := []test{
		{name: "missing header", expected: res{key: "", err: errors.New("no authorization header included")}, input: http.Header{}},
		{name: "malfored header", expected: res{key: "", err: errors.New("malformed authorization header")}, input: badHeader},
		{name: "good header", expected: res{key: "test", err: nil}, input: goodHeader},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Logf("%s test", test.name)
			key, err := GetAPIKey(test.input)
			if !reflect.DeepEqual(key, test.expected.key) || !reflect.DeepEqual(err, test.expected.err) {
				t.Fatalf("Expected key to be: %v Received: %v\nExpected error: %v, Received: %v", test.expected.key, key, test.expected.err, err)
			}
		})
	}
}
