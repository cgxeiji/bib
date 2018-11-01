package bib

import (
	"strings"
	"testing"
)

type pUnmarshal struct {
	bib    string
	title  []string
	length int
	fieldN []int
}

var testsUnmarshal = []pUnmarshal{
	{
		bib: `
@article{einstein1922,
  author = {Einstein, Albert},
  title = {The General Theory of Relativity},
  journaltitle = {The Meaning of Relativity},
  date = {1922},
  doi = {10.1007/978-94-011-6022-3_3}
}
`,
		title:  []string{"The General Theory of Relativity"},
		length: 1,
		fieldN: []int{5},
	},
	{
		bib: `
@article{einstein1922,
  author = {Einstein, Albert},
  title = {The General Theory of Relativity},
  journaltitle = {The Meaning of Relativity},
  date = {1922},
  doi = {10.1007/978-94-011-6022-3_3}
}

@article{bailly2010,
  title = {Gaze, conversational @ agents and face-to-face communication},
  author = {Bailly, Gérard and Raidt, Stephan and Elisei, Frédéric},
  date = {2010-6},
  journaltitle = {Speech Communication},
  number = {6},
  volume = {52},
  doi = {10.1016/j.specom.2010.02.015}
}

@article{meltzoff2010,
  date = {2010-10},
  journaltitle = {Neural Networks},
  title = {“Social” robots are psychological agents for infants: A test of gaze following},
  author = {Meltzoff, Andrew N. and Brooks, Rechele and Shon, Aaron P. and Rao, Rajesh P.N.},
  doi = {10.1016/j.neunet.2010.09.005},
  number = {8-9},
  volume = {23}
}

@inproceedings{thomaz2006,
  author = {{Thomaz, Andrea L.} and Hoffman, Guy and Breazeal, Cynthia},
  booktitle = {Proceeding of the 1st ACM SIGCHI/SIGART conference on Human-robot interaction  - HRI \'06},
  date = {2006},
  title = {Experiments in {{socially}} guided machine learning},
  doi = {10.1145/1121241.1121315}
}

`,
		title: []string{
			"The General Theory of Relativity",
			"Gaze, conversational @ agents and face-to-face communication",
			"“Social” robots are psychological agents for infants: A test of gaze following",
			"Experiments in {{socially}} guided machine learning",
		},
		length: 4,
		fieldN: []int{
			5,
			7,
			7,
			5,
		},
	},
}

func TestUnmarshal(t *testing.T) {
	for _, pair := range testsUnmarshal {
		entries, err := Unmarshal(strings.NewReader(pair.bib))
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
		for i, entry := range entries {
			if entry["title"] != pair.title[i] {
				t.Errorf(
					"\n%v\nfor:\n%v\nwant: %v\nhave: %v\n",
					pair.bib,
					entry,
					pair.title[i],
					entry["title"],
				)
			}
			if len(entry)-2 != pair.fieldN[i] {
				t.Errorf(
					"\n%v\nfor:\n%v\nwant: %v\nhave: %v\n",
					pair.bib,
					entry,
					pair.fieldN[i],
					len(entry)-2,
				)
			}
		}
	}
}
