package main

import (
	"fmt"
	"github.com/menymo/is-105-mod1/rivercrossing"
)

var kryssing rivercrossing.Situasjon

func main() {
	// HS begynner med å ta med kylling over elven, og drar tilbake med tom båt
	fmt.Println(kryssing.Kryss(&kryssing.HomoSapiens, &kryssing.Kylling))
	fmt.Println(kryssing.Kryss(&kryssing.HomoSapiens))
	// HS tar med korn over elven, og tar med seg kyllingen tilbake
	fmt.Println(kryssing.Kryss(&kryssing.Korn, &kryssing.HomoSapiens))
	fmt.Println(kryssing.Kryss(&kryssing.HomoSapiens, &kryssing.Kylling))
	// HS Tar med seg rev over elven, og drar tilbake med tom båt
	fmt.Println(kryssing.Kryss(&kryssing.HomoSapiens, &kryssing.Rev))
	fmt.Println(kryssing.Kryss(&kryssing.HomoSapiens))
	// HS tar med seg kylling over elven, og er deretter ferdig med jobben
	fmt.Println(kryssing.Kryss(&kryssing.HomoSapiens, &kryssing.Kylling))

	fmt.Println(kryssing)
}
