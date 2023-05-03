package main

import "os"

func main() {
	a := App{}
	a.initialize(
		os.Getenv("AUT_DATABASE_URL"),
		os.Getenv("AUT_BROKER_URL"),
		os.Getenv("AUT_AUD"),
		os.Getenv("AUT_ISS"),
	)
	a.Run(os.Getenv("AUT_PORT"))
}
