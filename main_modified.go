package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/gittoks/magic"
)

var hands = map[int]string{0: "tas", 1: "spock", 2: "kagit", 3: "kertenkele", 4: "makas"}

func index(c *magic.Context) {
	c.SendString("Hosgeldin. /:hand {tas, kagit, makas, kertenkele, spock} seklinde hamleni yapabilirsin.")
}

func play(c *magic.Context) {
	var result string = "Gecersiz hamle. /:hand {tas, kagit, makas, kertenkele, spock}"
	userHandID := -1
	handUser := strings.ToLower(c.Params["hand"])

	for k, v := range hands {
		if v == handUser {
			userHandID = k
			break
		}
	}
	if userHandID != -1 {
		goHandID := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(hands))
		handGo := hands[goHandID]
		whoWon := decide(userHandID, goHandID)
		result = "Senin hamlen:     " + handUser +
			"\nGo'nun hamlesi:   " + handGo +
			"\n" + whoWon + "\n"
	}
	c.SendString(result)
}

func decide(u int, g int) string {
	var msg string

	switch (u - g + 5) % 5 {
	case 0:
		msg = "Berabere!"
	case 1, 2:
		msg = "Kazandin!!"
	case 4, 3:
		msg = "Kaybettin."
	default:
		msg = "Bir seyler ters gitti."
	}

	return msg
}

func main() {
	m := magic.NewMagic(8081)
	m.GET("/", index)
	m.GET("/:hand", play)
	fmt.Println("Server started!")
	if err := m.ListenAndServe(); err != nil {
		fmt.Println("Failure:", err, "\nServer closed.")
	}
}
