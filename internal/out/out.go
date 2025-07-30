package out

import (
	"fmt"
	"encoding/json"
	"project02/internal/scanner"
	"log"
	"os"
)

func Print(jsonOut bool, humanRead bool, file *scanner.File) {
	if jsonOut {
		data, err := json.MarshalIndent(file, "", " ")
		if err != nil {
			log.Fatal("failed to encode json", err)
		}
		os.Stdout.Write(data)
		fmt.Println()
	} else {
		if humanRead {
			fmt.Printf("%s  %d  (%s)\n", file.Path, file.Size, file.SizeHuman)
		} else {
			fmt.Printf("%s  %d\n", file.Path, file.Size)
		}
	}
}
