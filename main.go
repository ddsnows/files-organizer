package main

import (
	"fmt"
	"os"
)

func main() {
	dir := "C:\\Users\\douglas.neves\\OneDrive\\Devs\\go\\src\\filename\\sample_files"

	list, err := os.ReadDir(dir)

	if err != nil {
		panic(err)
	}

	for _, each := range list {
		name := each.Name()
		newName := "EF_" + name
		fmt.Println(newName)
		//os.Rename(filepath.Join(dir, name), filepath.Join(dir, newName))
	}
}
