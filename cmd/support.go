package cmd

// Check returned error value. Takes into account verbosity settings.
func check(e error) {
	if e != nil {
		panic(e) // TODO: should swap to log.Fatal(err)
	}
}