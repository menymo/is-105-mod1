package myquote

import (
	"rsc.io/quote"
	"testing"
)

func TestGlass(t *testing.T) {
	wanted := quote.Glass()
	got := Glass()

	if got != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", got, wanted)
	}
}

func TestGo(t *testing.T) {
	wanted := quote.Go()
	got := Go()

	if got != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", got, wanted)
	}
}

func TestHello(t *testing.T) {
	wanted := quote.Hello()
	got := Hello()

	if got != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", got, wanted)
	}
}

func TestOpt(t *testing.T) {
	wanted := quote.Opt()
	got := Opt()

	if got != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", got, wanted)
	}
}
