package main

func main() {
	baseUrl := "https://pokeapi.co/api/v2/location-area/"
	config := cfg{
		&baseUrl,
		nil,
	}

	startRepl(&config)
}
