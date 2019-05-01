package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/shurcooL/vfsgen"
)

func main() {
	var cwd, _ = os.Getwd()
	directories := []string{"responses"}
	for _, dir := range directories {
		testdata := http.Dir(filepath.Join(cwd, "pkg", "rundeck", dir, "testdata"))
		if err := vfsgen.Generate(testdata, vfsgen.Options{
			Filename:     "pkg/rundeck/"+dir+"/testdata.go",
			PackageName:  dir,
			VariableName: "assets",
		}); err != nil {
			log.Fatalln(err)
		}
	}
}