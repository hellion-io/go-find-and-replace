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
		//pass in the string you want to find as '--find STRNGTOFNDANDREPLACE'
                //the third parameter in stringsReplace is what will replace the found string
		var find string
                var replace string
		flag.StringVar(&find, "find", "foo", "what you want to find")
		flag.StringVar(&replace, "replace", "bar", "what to replace it with")
		flag.Parse()

		//hardcoded replacement string (for now)
		newContents := strings.Replace(string(read), find, replace, -1)

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
