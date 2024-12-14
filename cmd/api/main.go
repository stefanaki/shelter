package main

import "log"

func main() {
	cfg := config{
		address: ":8080",
	}
	app := &api{
		config: cfg,
	}

	mux := app.mount()

	if err := app.serve(mux); err != nil {
		log.Fatalf("error serving api: %v", err)
	}
}
