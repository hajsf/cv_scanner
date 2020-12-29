//https://www.devdungeon.com/content/working-files-go//
/*
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	p := []string{}
	p = append(p, "engineer")
	p = append(p, "doctor")
	p = append(p, "chemical (permit)")
	skills := strings.Join(p, "|")

	fmt.Println(skills)
	re := regexp.MustCompile(`(?i)` + skills)

	for _, file := range files {
		if strings.ToLower(filepath.Ext(file.Name())) == ".txt" {
			fmt.Println(file.Name())

			b, err := ioutil.ReadFile(file.Name()) // just pass the file name
			if err != nil {
				fmt.Print(err)
			}
			//fmt.Println(b) // print the content as 'bytes'

			str := string(b) // convert content to a 'string'
			matches := re.FindAllString(str, -1)

			uniqueSlice := unique(matches)
			fmt.Println(uniqueSlice, len(uniqueSlice))
			for i, j := range uniqueSlice {
				fmt.Println(i, j)
			}
		}
	}
}
*/

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/ledongthuc/pdf"
)

func main() {
	pdf.DebugOn = true

	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	p := []string{"import", "shipping documents", "ERP", "supply chain", "(b/l|(bill of loading))", "chemical (permit)",
		"permit", "HS (code)", "Police (permit)"} //  to be scanned in the given CV
	skills := strings.Join(p, "|")

	re := regexp.MustCompile(`(?i)` + skills)

	for _, file := range files {
		if strings.ToLower(filepath.Ext(file.Name())) == ".pdf" {
			println(file.Name())
			println("=============================")
			content, err := readPdf(file.Name()) // Read local pdf file
			if err != nil {
				panic(err)
			}
			matches := re.FindAllString(content, -1)

			uniqueSlice := unique(matches)
			fmt.Println("skills found:", len(uniqueSlice), "/", len(p))
			fmt.Println("'" + strings.Join(uniqueSlice, `','`) + `'`)
			//for i, j := range uniqueSlice {
			//	fmt.Println(i, j)
			//}
		}
	}
	return
}

func readPdf(path string) (string, error) {
	f, r, err := pdf.Open(path)
	// remember close file
	defer f.Close()
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	b, err := r.GetPlainText()
	if err != nil {
		return "", err
	}
	buf.ReadFrom(b)
	return buf.String(), nil
}

func unique(stringSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
