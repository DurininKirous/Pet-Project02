package scanner

import (
	"strings"
	"os"
	"io/fs"
	"path/filepath"
)

type File struct {
	Path string
	Size int64
}

func ScanDir(dir string, maxDepth int) *File {
	maxFile := &File{Path: "null", Size: 0}
	filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rootAbs, _ := filepath.Abs(dir)
		pathAbs, _ := filepath.Abs(path)
		dirPathSeparators := strings.Count(rootAbs, string(os.PathSeparator))
		filePathPathSeparators := strings.Count(pathAbs, string(os.PathSeparator))

		if !d.IsDir() && ( filePathPathSeparators - dirPathSeparators <= maxDepth || maxDepth == 0 ) {
			info, err := d.Info()
			if err != nil {
				return err
			}
			if info.Size() > maxFile.Size {
				maxFile.Path = path
				maxFile.Size = info.Size()
			}
		}
		return nil
	})

	return maxFile
}


