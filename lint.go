package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/hashicorp/hcl/hclparse"
	"github.com/mitchellh/colorstring"
)

func main() {
	for i, arg := range os.Args {
		if i == 0 {
			continue
		}
		search := arg
		if info, err := os.Stat(arg); err == nil && info.IsDir() {
			search = fmt.Sprintf("%s/*.tf", arg)
		}
		files, err := filepath.Glob(search)
		if err != nil {
			colorstring.Printf("[red]Error finding files: %s", err)
		}
		for _, filename := range files {
			fmt.Printf("Checking %s ... ", filename)
			file, err := ioutil.ReadFile(filename)
			if err != nil {
				colorstring.Printf("[red]Error reading file: %s\n", err)
				break
			}
//			#_, err = hclparse.NewParser(string(file))
                        parser := hclparse.NewParser()
			_, err = parser.ParseHCLFile(string(file))
			if err != nil {
				colorstring.Printf("[red]Error parsing file: %s\n", err)
				break
			}
			colorstring.Printf("[green]OK!\n")
		}
	}

}
