// Package vreme vsebuje funkcije ter tip za delo z vremenskimi postajami in meritvami.
package vreme

import "fmt"

// Postaja predstavlja eno instanco vremenske postaje, ki vsebuje tudi njene meritve.
type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

// sprintaj_rezino je pomagalska funkcija za printanje meritev v ukazno vrstico.
func sprintaj_rezino(meritve []float64) {
	fmt.Printf("[")

	for _, c := range meritve {
		fmt.Printf("%f ", c)
	}

	fmt.Println("]")
}

// DodajMeritev doda ene meritve v podan slovar z postajami.
// Funkcija doda meritev z temperaturo temperatura v struct postaje z indeksom oznaka v slovarju postaje.
// Funkcija pred dodajanjem preveri tudi, ali je meritev v pravilnem intervalu.
func DodajMeritev(postaje map[string]Postaja, oznaka string, temperatura float64, mintemp float64, maxtemp float64) {
	postaja, obstaja := postaje[oznaka]

	if !obstaja {
		fmt.Println("Napaka! Postaja z oznako " + oznaka + " ne obstaja!")
		return
	}

	if temperatura > maxtemp || temperatura < mintemp || !obstaja {
		fmt.Println("Napaka! Temperatura je izven veljavnih območij!")
		return
	}

	postaja.meritve = append(postaja.meritve, temperatura)
	postaje[oznaka] = postaja
}

// statistika vrne minimalno, maksimalno ter povprečno meritev postaje z oznako oznaka, ter boolean vrednost.
// Če je št meritev neke postaje premajhno, funkcija vrne boolean vrednost false, sicer vrne true.
func statistika(postaje map[string]Postaja, oznaka string, stmeritev int) (najnizja, najvisja, povprecna float64, ok bool) {

	postaja, obstaja := postaje[oznaka]

	var dolzina float64 = float64(len(postaja.meritve))

	if !obstaja || len(postaja.meritve) < stmeritev {

		return 0, 0, 0, false

	}

	min, max := postaja.meritve[0], postaja.meritve[0]
	var sum float64 = 0

	for _, c := range postaja.meritve {

		if c > max {
			max = c
		}

		if c < min {
			min = c

		}

		sum = sum + c

	}

	var avg float64 = sum / dolzina

	return min, max, avg, true
}

// IzpisMeritev izpiše meritve vseh postaj v podanem slovarju.
func IzpisMeritev(postaje map[string]Postaja) {
	fmt.Println("MERITVE:")
	for kljuc, postaja := range postaje {
		fmt.Printf("%s - %s (%s): ", kljuc, postaja.ime, postaja.kraj)
		sprintaj_rezino(postaja.meritve)
	}

}

// IzpisPorocila izpiše poročilo za vsako postajo v podanem slovarju.
// Funkciija izpiše obcutek temperatre, če pa je št meritev premalo, funkcija to zazna in izpiše primerno sporočilo.
func IzpisPorocila(postaje map[string]Postaja, stmeritev int, min float64, max float64) {
	for kljuc, postaja := range postaje {

		var avg float64 = 0
		var ok bool = true
		var obcutek string = ""
		min, max, avg, ok = statistika(postaje, kljuc, stmeritev)

		switch {
		case !ok:
			fmt.Println(postaja.ime + " : premalo meritev!")
			continue
		case avg >= 25:
			obcutek = "Vroče"
		case avg >= 15 && avg < 25:
			obcutek = "Toplo"

		case avg >= 5 && avg < 15:
			obcutek = "Hladno"
		case avg < 5:
			obcutek = "Mrzlo"
		}
		fmt.Printf("%s: min %f C, max %f C, povprečje %f C -> %s\n", postaja.ime, min, max, avg, obcutek)
	}
}

// DodajPostajo doda postajo z podanimi atributi v podan slovar.
func DodajPostajo(ime string, kraj string, meritve []float64, key string, slovar map[string]Postaja) {
	postaja1 := Postaja{ime, kraj, meritve}

	slovar[key] = postaja1
}
