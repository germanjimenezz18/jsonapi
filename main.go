package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	seed := flag.Bool("seed", false, "seed DB")
	flag.Parse()

	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	//seed stuff
	if *seed {
		fmt.Println("seeding the DB")
		seedAccounts(store)
	}

	server := NewAPIServer(":3000", store)
	server.Run()
}

func seedAccount(store Storage, fname, lname, pw string) *Account {
	acc, err := NewAccount(fname, lname, pw)
	if err != nil {
		log.Fatal(err)
	}
	if err := store.CreateAccount(acc); err != nil {
		log.Fatal(err)
	}

	fmt.Println("new account number =>", acc.Number)

	return acc

}

func seedAccounts(s Storage) {
	seedAccount(s, "germanciopass", "jimenez", "huntermada")
}
