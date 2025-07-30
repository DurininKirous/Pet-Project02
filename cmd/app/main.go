package main

import (
	"project02/internal/validate"
	"project02/internal/scanner"
	"project02/internal/flags"
	"project02/internal/out"
)

func main() {
	flags.InitFlags()
	validate.ValidateInputDir(flags.Dir)
	file := scanner.ScanDir(flags.Dir, flags.MaxDepth)
	out.Print(flags.JSONOut, flags.HumanRead, file)
}
