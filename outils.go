package outils

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func Dejatester(ucl string, s [30]string) bool {
	for _, v := range s {
		if v == ucl {
			return true
		}
	}
	return false
}
func Motrandom(s []string) string {
	rand.Seed(time.Now().UnixNano())
	return s[rand.Intn(len(s))]
}
func Changedemot(s string, news string, nb int) string {
	nnews := ""
	for i := 0; i <= len(s)-1; i++ {
		if i != nb {
			nnews += string(news[i])
		}
		if i == nb {
			nnews += string(s[i])
		}
	}
	return nnews
}
func Nouveaumot(s string, nb int) string {
	var newmsg string
	for range s {
		newmsg += "_"
	}
	if nb != 0 {
		for i := 0; i < nb; i++ {
			rand.Seed(time.Now().UnixNano())
			indexts := rand.Intn(len(s) - 1)
			newmsg = Changedemot(s, newmsg, indexts)
		}
	}

	return newmsg
}
func hangman(hp int) {
	f, _ := os.Open("hangman.txt")
	scanner := bufio.NewScanner(f)
	var line int
	a := 71 - hp*8
	b := 78 - hp*8
	for scanner.Scan() {
		if line >= a && line <= b {
			fmt.Println(scanner.Text())
		}
		line++
	}
}
func Tcheking(Char string, word string) bool { // ici ce passe la verif
	for _, v := range word {
		if Char == string(v) {
			return true
		}
	}
	return false
}

func Bonnelettre(ucl string, word string, hword string) string { //ici il viens recuperer la lettre
	newhword := ""
	for i, v := range word {
		if ucl == string(v) {
			newhword += string(v)
		} else {
			newhword += string(hword[i])
		}
	}
	return newhword
}
func Pendu(hp int) {
	f, _ := os.Open("hangman.txt")
	scanner := bufio.NewScanner(f)
	var line int
	a := 71 - hp*8
	b := 78 - hp*8
	for scanner.Scan() {
		if line >= a && line <= b {
			fmt.Println(scanner.Text())
		}
		line++
	}
}
