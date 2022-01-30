package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
)

type KicadFootprint struct {
	LibraryName string
}

func main() {
	footprintsSrcPath := flag.String("path", "", "Path to look for .pretty footprints")
	flag.Parse()

	if *footprintsSrcPath == "" {
		log.Fatal("-path not specified!")
	}

	var libs []KicadFootprint
	fileInfo, err := ioutil.ReadDir(*footprintsSrcPath)
	if err != nil {
		log.Fatal(err)
	}

	t, err := template.ParseFiles("fp-lib-table.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	outFile, err := os.Create("fp-lib-table")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range fileInfo {
		//files = append(files, file.Name())
		if strings.Contains(file.Name(), ".pretty") {
			libName := strings.Split(file.Name(), ".")[0]
			libs = append(libs, KicadFootprint{LibraryName: libName})
		}
	}
	t.Execute(outFile, libs)

}
