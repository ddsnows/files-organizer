package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	dir := "C:\\Users\\douglas.neves\\OneDrive\\Devs\\go\\src\\files_organizer\\sample_files"
	//newDir := ""
	//prevFolderName := ""

	list, err := os.ReadDir(dir)

	if err != nil {
		panic(err)
	}

	for _, each := range list {

		newName := "EF_"
		name := each.Name()

		// Process only files doc type
		if !strings.HasSuffix(name, ".doc") && !strings.HasSuffix(name, ".docx") {
			continue
		}

		// Create standardization to filename start with 'EF_XX.999.9 - '
		if !strings.HasPrefix(name, "EF_") {

			if name[2:3] == " " || name[2:3] == "_" {
				name = name[0:2] + "." + name[3:len(name)]
			}

			if name[6:7] == "_" {
				name = name[0:6] + " - " + name[7:len(name)]

			} else if name[6:7] == " " {

				if x, err := strconv.Atoi(name[7:8]); err == nil {

					if x == 0 {
						name = name[0:6] + "." + name[8:len(name)]
					} else {
						name = name[0:6] + "." + name[7:len(name)]
					}
				}
			}

			x := name[8:9]
			if x == "-" || name[8:9] == "_" || name[8:9] == " " {
				name = name[0:8] + " - " + name[9:len(name)]
			}

			// Keep tag _v99 always lower case
			v_tag := strings.LastIndex(strings.ToLower(name), "_v")
			if v_tag > 0 {
				name = name[0:v_tag] + strings.ToLower(name[v_tag:len(name)])
			}

			newName = newName + name
			fmt.Println("(Modified Name) - ", newName)

		} else {
			newName = name
			fmt.Println("(Same Name - ", newName)
			continue
		}

		/*
			// Eliminate last part of string refering to file version before define new directory
			folderName := newName[0:strings.LastIndex(strings.ToLower(newName), "_v")]
			fmt.Println(folderName)

			// Create new directory once using documenct name
			if prevFolderName != folderName {
				newDir = filepath.Join(dir, folderName)

				err := os.Mkdir(newDir, 0777)
				if err != nil {
					log.Fatal(err)

				} else {
					prevFolderName = folderName
				}

				// Move file to the new directory
				os.Rename(filepath.Join(dir, name), filepath.Join(newDir, newName))

			} else {
		*/
		os.Rename(filepath.Join(dir, name), filepath.Join(dir, newName))

		//}
	}
}
