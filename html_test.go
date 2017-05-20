package main

import (
	"testing"
)

func TestHTML(t *testing.T) {
	s := modifyUrl("test", "/css/1.css", "(url( \"../file\")) url('/file') url(file)")
	t.Log(s)
}
