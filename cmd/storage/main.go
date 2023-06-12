package main

import (
	"fmt"
	"log"

	"github.com/HenRok1/internal/storage"
)

func main() {
	st := storage.NewStorage()

	file, err := st.Upload("test.txt", []byte("Hello world"))
	if err != nil {
		log.Fatal(err)
	}

	restoredFile, err := st.GetByID(file.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("it work", restoredFile)
}
