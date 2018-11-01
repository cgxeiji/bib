package bib

import (
	"bytes"
	"io"
	"regexp"
	"strings"
)

// Unmarshal parses a bibtex/biblatex file into an array of map[string]string with
// the values for each entry.
func Unmarshal(bib io.Reader) ([]map[string]string, error) {
	entries := []map[string]string{}

	// search for "@type{key"
	atMark := regexp.MustCompile(`@[[:alpha:]]+?{.*?,`)
	// search for "@type{key,...}"
	block := regexp.MustCompile(`@[[:alpha:]]+?{.*?}[^,][^@]*`)
	// search for "field = {content}"
	field := regexp.MustCompile(`[[:alpha:]]+[\s]*=[\s]*{.*?}(,|\z)`)
	// search for "field"
	name := regexp.MustCompile(`[[:alpha:]]*`)
	// search for "{content}"
	content := regexp.MustCompile(`{.*}`)

	buffer := new(bytes.Buffer)
	if _, err := buffer.ReadFrom(bib); err != nil {
		return entries, err
	}

	bibString := buffer.String()
	bibString = strings.Replace(bibString, "\n", "", -1)

	for _, s := range block.FindAllString(bibString, -1) {
		e := map[string]string{}

		a := atMark.FindString(s)
		a = a[1 : len(a)-1]
		as := strings.Split(a, "{")

		e["type"] = as[0]
		e["key"] = as[1]

		for _, f := range field.FindAllString(s, -1) {
			f = f[:len(f)-1]
			n := name.FindString(f)
			c := content.FindString(f)
			c = strings.Replace(c, "{", "", -1)
			c = strings.Replace(c, "}", "", -1)
			e[n] = c
		}

		entries = append(entries, e)
	}

	return entries, nil
}
