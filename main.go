package main

import "codesmell/golang"

func main() {
	golang.Join2files("a.txt", "b.txt")
	golang.Join2filesCleanly("a.txt", "b.txt")
}
