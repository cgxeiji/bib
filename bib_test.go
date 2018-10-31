package bib

import (
	"strings"
	"testing"

	"github.com/kr/pretty"
)

type pairParse struct {
	bib    string
	title  string
	length int
}

var testsParse = []pairParse{
	{
		bib: `
@article{einstein1922,
  author = {Einstein, Albert},
  title = {The General Theory of Relativity},
  journaltitle = {The Meaning of Relativity},
  date = {1922},
  doi = {10.1007/978-94-011-6022-3_3},
}
`,
		title:  "The General Theory of Relativity",
		length: 1,
	},
}

func TestParse(t *testing.T) {
	for _, pair := range testsParse {
		entries, err := Parse(strings.NewReader(pair.bib))
		if err != nil {
			t.Error(err)
		}

		l := len(entries)
		if l != pair.length {
			t.Errorf(
				"\nfor:\n%v\nwant: %v\nhave: %v\n",
				pair.bib,
				pair.length,
				l,
			)
		}
		title := entries[0]["title"]
		if title != pair.title {
			t.Errorf(
				"\nfor:\n%v\nwant: %v\nhave: %v\n",
				pair.bib,
				pair.title,
				title,
			)
		}

		pretty.Println(entries)
	}
}
