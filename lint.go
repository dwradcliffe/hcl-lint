package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

        "github.com/hashicorp/hcl/hclparse"
        "github.com/hashicorp/hcl"
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
			        colorstring.Printf("[red]File Content: %s", file)
				break
			}
			parser := hclparse.NewParser()
			_, diags := parser.ParseHCLFile(filename)
                        wr := hcl.NewDiagnosticTextWriter(
                            os.Stdout,      // writer to send messages to
                            parser.Files(), // the parser's file cache, for source snippets
                            78,             // wrapping width
                            true,           // generate colored/highlighted output
                        )
			if diags.HasErrors() {
			        fmt.Printf("\n")
			        wr.WriteDiagnostics(diags)
				break
			}
			colorstring.Printf("[green]OK!\n")
		}
	}
}
