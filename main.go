package main

func main() {
	handlers, err := load("config.json")
	if err != nil {
		panic(err)
	}

	serve(handlers)
}
