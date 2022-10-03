package main

import "auth/internal"

func main() {
	r := internal.NewRouter()
	r.Run()
}
