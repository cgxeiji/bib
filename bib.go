package bib

import (
	"bufio"
	"io"
	"strings"
)

// Parse parses a bibtex/biblatex file into an array of map[string]string with
// the values for each entry.
func Parse(bib io.Reader) ([]map[string]string, error) {
	entries := []map[string]string{}

	scanner := bufio.NewScanner(bib)

	entry := map[string]string{}
	for scanner.Scan() {
		w, ts := parseLine(scanner.Text())
		switch w {
		case '@':
			entry = map[string]string{}
			entry["type"] = ts[0]
			entry["key"] = ts[1]
		case '}':
			entries = append(entries, entry)
		case 'v':
			entry[ts[0]] = strings.TrimSuffix(strings.TrimPrefix(ts[1], "{"), "}")
		}
	}

	return entries, nil
}

func parseLine(line string) (byte, []string) {
	line = strings.TrimSpace(line)
	line = strings.TrimSuffix(line, ",")
	things := []string{}
	var which byte

	if len(line) > 0 {
		which = line[0]
		switch which {
		case '@':
			things = strings.Split(line[1:], "{")
		case '}':
			break
		default:
			things = strings.Split(line, " = ")
			which = 'v'
		}
	}

	return which, things
}
