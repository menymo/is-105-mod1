package rivercrossing

import (
	"testing"
)

func TestSituasjon_Kryss(t *testing.T) {
	var situasjon Situasjon

	// HS begynner med å ta med kylling over elven, og drar tilbake med tom båt
	err := situasjon.Kryss(&situasjon.HomoSapiens, &situasjon.Kylling)
	if err != nil {
		t.Error(err)
	}
	err = situasjon.Kryss(&situasjon.HomoSapiens)
	if err != nil {
		t.Error(err)
	}
	// HS tar med korn over elven, og tar med seg kyllingen tilbake
	err = situasjon.Kryss(&situasjon.HomoSapiens, &situasjon.Korn)
	if err != nil {
		t.Error(err)
	}
	err = situasjon.Kryss(&situasjon.HomoSapiens, &situasjon.Kylling)
	if err != nil {
		t.Error(err)
	}
	// HS Tar med seg rev over elven, og drar tilbake med tom båt
	err = situasjon.Kryss(&situasjon.HomoSapiens, &situasjon.Rev)
	if err != nil {
		t.Error(err)
	}
	err = situasjon.Kryss(&situasjon.HomoSapiens)
	if err != nil {
		t.Error(err)
	}
	// HS tar med seg kylling over elven, og er deretter ferdig med jobben
	err = situasjon.Kryss(&situasjon.HomoSapiens, &situasjon.Kylling)
	if err != nil {
		t.Error(err)
	}
}

// Rev kan ikke være alene med Kylling
func TestSituasjon_SjekkRegler1(t *testing.T) {
	var situasjon Situasjon

	situasjon.Rev = true
	situasjon.HomoSapiens = true

	err := situasjon.SjekkRegler()
	if err != nil {
		t.Logf("Regel 1 pass: %s", err)
	} else {
		t.Error("oof")
	}
}

// Kylling kan ikke være alene med Korn
func TestSituasjon_SjekkRegler2(t *testing.T) {
	var situasjon Situasjon
	situasjon.Kylling = true
	situasjon.Korn = true

	err := situasjon.SjekkRegler()
	if err != nil {
		t.Logf("Regel 2 pass: %s", err)
	} else {
		t.Fail()
	}
}
