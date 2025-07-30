package flags

import "flag"

var (
	Dir string
	//MinSize int
	MaxDepth int
	JSONOut bool
)

func InitFlags() {
	flag.StringVar(&Dir, "dir", ".", "Directory to scan")
	//flag.IntVar(&MinSize, "min-size", 10, "Minimum file size in MB")
	flag.IntVar(&MaxDepth, "max-depth", 1, "Minimum file size in MB")
	flag.BoolVar(&JSONOut, "json", false, "Output in JSON format")
	flag.Parse()
}
