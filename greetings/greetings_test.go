package greetings 

import (
	"testing"
	"regexp"
)

func TestHelloName(t *testing.T) {
	name := "Vinayaka"
	want := regexp.MustCompile(`\b` + name + `\b`)
	message, err := Hello(name)

	if want.MatchString(message) == false || err != nil {
		t.Errorf("Hello(%q) = %q, %v, want match for %q and no error", name, message, err, want)
	} 
}

func TestHelloEmpty(t *testing.T) {
	message, err := Hello("")
	if message != "" || err == nil {
		t.Errorf("Hello(\"\") = %q, %v, want \"\", error", message, err)
	}
}