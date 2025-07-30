package out

import (
	"fmt"
	"encoding/json"
	"project02/internal/scanner"
	"log"
	"os"
)

func Print(jsonOut bool, file *scanner.File) {
	if jsonOut {
		data, err := json.MarshalIndent(file, "", " ")
		if err != nil {
			log.Fatal("failed to encode json", err)
		}
		os.Stdout.Write(data)
		fmt.Println()
	} else {
		fmt.Printf("%s  %d\n", file.Path, file.Size)
	}
}
