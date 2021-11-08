package main

import (
		"archive/tar"
		"compress/gzip"
		"flag"
		"fmt"
		"io"
		"os"
		"strings"
		"path/filepath"
)

func checkerror(err error) {

		if err != nil {
				fmt.Println(err)
				os.Exit(1)
		}
}

func main() {

		flag.Parse()

		destinationfile := flag.Arg(0)

		if destinationfile == "" {
				fmt.Println("Usage : ./compres filename.tar.gz ./destinationfile")
				os.Exit(1)
		}

		sourcedir := flag.Arg(1)

		if sourcedir == "" {
				fmt.Println("Usage : ./compres filename.tar.gz ./destinationfile")
				os.Exit(1)
		}

		dir, err := os.Open(sourcedir)

		checkerror(err)

		defer dir.Close()

		files, err := dir.Readdir(0)

		checkerror(err)

		tarfile, err := os.Create(destinationfile)

		checkerror(err)

		defer tarfile.Close()
		var fileWriter io.WriteCloser = tarfile

		if strings.HasSuffix(destinationfile, ".gz") {
				fileWriter = gzip.NewWriter(tarfile) 
				defer fileWriter.Close()            
		}

		tarfileWriter := tar.NewWriter(fileWriter)
		defer tarfileWriter.Close()

		for _, fileInfo := range files {

				if fileInfo.IsDir() {
				   continue
				}

				file, err := os.Open(dir.Name() + string(filepath.Separator) + fileInfo.Name())

				checkerror(err)

				defer file.Close()

				header := new(tar.Header)
				header.Name = file.Name()
				header.Size = fileInfo.Size()
				header.Mode = int64(fileInfo.Mode())
				header.ModTime = fileInfo.ModTime()


				err = tarfileWriter.WriteHeader(header)

				checkerror(err)

				_, err = io.Copy(tarfileWriter, file)

				checkerror(err)
		}

}