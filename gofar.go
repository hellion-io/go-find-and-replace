package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func gofar(path string, fi os.FileInfo, err error) error {

	if err != nil {
		return err
	}

	if !!fi.IsDir() {
		return nil //
	}

	matched, err := filepath.Match("*.yml", fi.Name())

	if err != nil {
		panic(err)
		return err
	}

	if matched {
		read, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}
		fmt.Println(path)

		//capture the flag
		//pass in the string you want to find as '--svar STRNGTOFNDANDREPLACE'
                //the third parameter in stringsReplace is what will replace the found string
		var finder string
		flag.StringVar(&svar, "svar", "foo", "a command line flag")
		flag.Parse()

		//hardcoded replacement string (for now)
		newContents := strings.Replace(string(read), svar, "REPLACEME", -1)

		fmt.Println(newContents)

		err = ioutil.WriteFile(path, []byte(newContents), 0)
		if err != nil {
			panic(err)
		}

	}

	return nil
}

func main() {
	err := filepath.Walk(".", gofar)
	if err != nil {
		panic(err)
	}
}
