package main

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")

	for k, v := range commands {
		fmt.Printf("%s: %s\n", k, v.description)
	}

	return nil
}
