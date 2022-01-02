/*
Copyright 2021 The terraform docs Author.
Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.
You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package reader

import (
	"bufio"
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

// Lines represents line reader in a given 'FileName' immediately
// before the given 'LineNum'. Extraction happens when 'Condition'
// is met and being processed by 'Parser' function.
type Lines struct {
	FileName  string
	LineNum   int // value -1 means scan the whole file and break after finding what we were looking for
	Condition func(line Line, lines []Line) bool
	Parser    func(line Line) ([]string, bool)
}

type Line struct {
	Parent  string
	Value   string
	LineNum int
}

// Extract extracts lines in given file and based on the provided
// condition. returns empty if nothing found.
func (l *Lines) ExtractFromText() []Line {
	f, err := os.Open(l.FileName)
	if err != nil {
		log.Error(err)
		return nil
	}
	stat, err := f.Stat()
	if err != nil {
		log.Error(err)
		return nil
	}
	if stat.Size() == 0 {
		return []Line{}
	}
	defer f.Close()
	return l.extract(f)
}

func (l *Lines) extract(r io.Reader) []Line {
	bf := bufio.NewScanner(r)
	bf.Split(bufio.ScanLines)

	var lines = make([]Line, 0)
	i := 0
	lineP := ""
	for bf.Scan() {
		var line Line
		line.Value = bf.Text()
		line.Parent = lineP
		line.LineNum = i
		if l.Condition(line, lines) {
			if extracted, capture := l.Parser(line); capture {
				for _, word := range extracted {
					line.Value = word
					lines = append(lines, line)
				}
			}
		}

		lineP = line.Value
		i++
	}

	return lines
}
