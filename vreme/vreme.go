package vreme

import (
	"fmt"
	"slices"
)

type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

func DodajPostajo(postaje map[string]Postaja, oznaka, ime, kraj string) {
	postaje[oznaka] = Postaja{ime: ime, kraj: kraj, meritve: []float64{}}
}

func DodajMeritev(postaje map[string]Postaja, oznaka string, temperatura float64, minTemp, maxTemp float64) {
	p, ok := postaje[oznaka]
	if ok == false {
		fmt.Println("Napaka: Postaja s tem imeno ne obstaja")
		return
	}
	if temperatura < minTemp || temperatura > maxTemp {
		fmt.Println("Temperatura ni ustrezna")
		return
	}
	p.meritve = append(p.meritve, temperatura)
	postaje[oznaka] = p
}

// skrita
func statistika(postaje map[string]Postaja, oznaka string, stMeritev int) (najnizja, najvisja, povprecna float64, ok bool) {
	p, oki := postaje[oznaka]
	if len(p.meritve) < stMeritev || oki == false {
		return
	}
	najnizja = slices.Min(p.meritve)
	najvisja = slices.Max(p.meritve)
	var vsota float64
	for _, x := range p.meritve {
		vsota += x
	}
	povprecna = vsota / float64(len(p.meritve))
	ok = true
	return najnizja, najvisja, povprecna, ok
}

func IzpisMeritev(postaje map[string]Postaja) {
	fmt.Println("MERITVE:")
	for oznaka, p := range postaje {
		fmt.Printf("%s\n%s (%s): %v\n", oznaka, p.ime, p.kraj, p.meritve)
	}
}

func IzpisPorocila(postaje map[string]Postaja, stMeritev int) {
	for oznaka, p := range postaje {
		mini, maxi, avgi, ok := statistika(postaje, oznaka, stMeritev)
		if !ok {
			fmt.Printf("%s: premalo meritev\n", p.ime)
			continue
		}
		var opis string
		switch {
		case avgi >= 25:
			opis = "Vroče"
		case avgi >= 15:
			opis = "Toplo"
		case avgi >= 5:
			opis = "Hladno"
		default:
			opis = "Mrzlo"
		}
		fmt.Printf("%s: min %.1f °C, max %.1f °C, povprečje %.1f °C -> %s\n", p.ime, mini, maxi, avgi, opis)
	}
}
