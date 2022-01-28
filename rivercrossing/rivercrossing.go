package rivercrossing

import (
	"errors"
	"fmt"
	"strings"
)

// Elvekryssing-situasjon. True vil si at de har krysset elven
type Situasjon struct {
	Rev         bool
	Korn        bool
	Kylling     bool
	HomoSapiens bool
}

func (r Situasjon) SjekkRegler() error {
	// Rev kan ikke være alene med kylling
	if (r.Rev == r.Kylling) && (r.HomoSapiens != r.Rev) {
		return errors.New("Rev kan ikke være alene med kylling")
	}
	// Kylling kan ikke være alene med korn
	if (r.Kylling == r.Korn) && (r.HomoSapiens != r.Kylling) {
		return errors.New("Kylling kan ikke være alene med korn")
	}
	return nil
}

func (r *Situasjon) Kryss(args ...*bool) error {
	// Hvis det kun er ett argument (HomoSapiens) kan båten fortsatt krysse elven
	if len(args) == 1 {
		if args[0] == &r.HomoSapiens {
			*args[0] = !*args[0]

			err := r.SjekkRegler()

			err = r.MakeASCIIArt(args...)
			if err != nil {
				return err
			}

			return err
		}
		return errors.New("HomoSapiens må kontrollere båten")
	}
	// Sjekk for duplikater
	if args[0] == args[1] {
		return errors.New("Kan ikke ha to like argumenter")
	}
	// Begge gjenstandene må være på samme side av elven
	if *args[0] != *args[1] {
		return errors.New("Begge gjenstandene må være på samme side av elven")
	}
	// En av gjenstandene i båten MÅ være HomoSapiens
	if (args[0] == &r.HomoSapiens) || (args[1] == &r.HomoSapiens) {
		*args[0] = !*args[0]
		*args[1] = !*args[1]

		err := r.SjekkRegler()
		if err != nil {
			return err
		}

		err = r.MakeASCIIArt(args...)
		if err != nil {
			return err
		}

		return nil

	}
	return errors.New("Unknown error occured")
}

// Lag ASCII art for å illustrere kryssingen, tenk ViewState() fra eksempelkoden
// Sjekker ikke om argumentene er gyldige, men dette er ingen stor risiko da Kryss() allerede har sjekket sin egen input
func (r *Situasjon) MakeASCIIArt(args ...*bool) error {
	var høyre string
	var venstre string

	r.PlasserVenstreHøyre(&r.HomoSapiens, &venstre, &høyre)
	r.PlasserVenstreHøyre(&r.Rev, &venstre, &høyre)
	r.PlasserVenstreHøyre(&r.Kylling, &venstre, &høyre)
	r.PlasserVenstreHøyre(&r.Korn, &venstre, &høyre)

	fmt.Println("høyre", høyre)
	fmt.Println("venstre", venstre)

	var båtfører *bool
	var passasjer *bool
	var sete1 string = ""
	var sete2 string = ""

	// Sjekk hvilket argument er båtfører eller passager
	if len(args) < 1 {
		if *args[0] == r.HomoSapiens {
			båtfører = args[0]
			passasjer = args[1]
		} else if *args[1] == r.HomoSapiens {
			båtfører = args[1]
			passasjer = args[0]
		} else {
			return errors.New("Kunne ikke finne HS i argument 1 eller 2")
		}
		fmt.Println(båtfører, passasjer)
	} else {
		båtfører = args[0]
	}

	if *båtfører {
		fmt.Printf("[%s ---V_____________________ \\_____%s__%s_____/ Ø--- %s]\n", venstre, sete1, sete2, høyre)
		sete1 = r.GiNavn(args[0])
		sete2 = r.GiNavn(args[1])
		høyre = strings.Replace(høyre, " "+sete1, "", 1)
		høyre = strings.Replace(høyre, " "+sete2, "", 1)
		fmt.Printf("[%s ---V_____________________ \\_____%s__%s_____/ Ø--- %s]\n", venstre, sete1, sete2, høyre)
		fmt.Printf("[%s ---V \\_____%s__%s_____/ _____________________Ø--- %s]\n", venstre, sete1, sete2, høyre)
		venstre = venstre + sete1 + " " + sete2 + " "
		sete1 = ""
		sete2 = ""
		fmt.Printf("[%s ---V \\_____%s__%s_____/ _____________________Ø--- %s]\n", venstre, sete1, sete2, høyre)
	} else {
		fmt.Printf("[%s ---V \\_____%s__%s_____/ _____________________Ø--- %s]\n", venstre, sete1, sete2, høyre)
		sete1 = r.GiNavn(args[0])
		if len(args) > 1 {
			sete2 = r.GiNavn(args[1])
		} else {
			sete2 = ""
		}
		venstre = strings.Replace(høyre, " "+sete1, "", 1)
		venstre = strings.Replace(høyre, " "+sete2, "", 1)
		fmt.Printf("[%s ---V \\_____%s__%s_____/ _____________________Ø--- %s]\n", venstre, sete1, sete2, høyre)
		fmt.Printf("[%s ---V_____________________ \\_____%s__%s_____/ Ø--- %s]\n", venstre, sete1, sete2, høyre)
		høyre = venstre + sete1 + " " + sete2 + " "
		sete1 = ""
		sete2 = ""
		fmt.Printf("[%s ---V_____________________ \\_____%s__%s_____/ Ø--- %s]\n", venstre, sete1, sete2, høyre)
	}
	return nil
}

func (r *Situasjon) PlasserVenstreHøyre(gjenstand *bool, venstre *string, høyre *string) {
	if *gjenstand {
		*høyre = " " + r.GiNavn(gjenstand) + *høyre
	} else {
		*venstre = *venstre + r.GiNavn(gjenstand) + " "
	}
}

func (r *Situasjon) GiNavn(gjenstand *bool) string {
	var navn string
	switch gjenstand {
	case &r.Rev:
		navn = "🦊"
	case &r.Kylling:
		navn = "🐓"
	case &r.Korn:
		navn = "🌾"
	case &r.HomoSapiens:
		navn = "👨"
	default:
		navn = "err"
	}
	return navn
}

// [Kylling Rev Korn HS ---V \____________/ _____________________Ø---]
