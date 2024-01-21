package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Valentin"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Valentin")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Valentin") = %q, %v, quiere coindicendia para %#q, nil `, msg, err, want)
	}
}

func TestHelloEmptyName(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, quiere "", error`, msg, err)
	}

}
