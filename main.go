package main

//go build -ldflags "-H=windowsgui -s -w"

import (
	"github.com/ncruces/zenity"
	"math/rand"
	"time"
)

func tony() {
	for {
		rand.Seed(time.Now().Unix())
		msg := phrases[rand.Intn(len(phrases))]
		go msgBox(msg)
		time.Sleep(time.Duration(rand.Intn(4)) * time.Second)
		msg = phrases[rand.Intn(len(phrases))]
		go msgBox(msg)
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
		msg = phrases[rand.Intn(len(phrases))]
		msgBox(msg)
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	}
}

func msgBox(msg string) {
	zenity.Question(msg,
		zenity.OKLabel("D'accord"),
		zenity.CancelLabel("STOP !"),
		zenity.Width(1000),
		zenity.Height(1000),
		zenity.Title("Salut c'est Anthony !"),
		zenity.Icon("N:\\COMMUN\\ELEVE\\INFO\\SI-T1b\\GRP1\\Jan - Luca\\Antony\\test.ico"))
}

func main() {
	tony()
}
