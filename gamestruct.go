package main

type Game struct {
	Wordtested string
	Health     int
	Fullword   string
	Hiddenword string
}
type Lettreut struct {
	lettre [30]string
}
type Gamesave struct {
	Wordtested string `json:"Wordtested"`
	Health     int    `json:"Health"`
	Fullword   string `json:"Fullword"`
	Hiddenword string `json:"Hiddenword"`
	Flag       bool
	Tab        [30]string `json:"Tab"`
}
