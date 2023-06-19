package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Juan"

	// la expresión regular buscará una coincidencia exacta con el nombre
	want := regexp.MustCompile(`\b` + name + `\b`)

	// le pasamos tal cual tenemos en name
	msg, err := Hello("Juan")

	// si no coincide ...
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Juan") = %q, %v, quiere coincidencia para %#q, nil`, msg, err, want)
	}
}

// si name esta vacio
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, quiere "", error`, msg, err)
	}
}
