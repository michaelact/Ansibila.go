package helper

import (
	"bufio"
	"io"
	"os"
)

type Lines struct {
	FileName  string
	Condition func(line string) bool
	Parser    func(line string) []string
}

type Line struct {
	Parent string
	Value  string
}

func (l *Lines) Extract() ([]string, error) {
	f, err := os.Open(l.FileName)
	if err != nil {
		return nil, err
	}
	stat, err := f.Stat()
	if err != nil {
		return nil, err
	}
	if stat.Size() == 0 {
		return []string{}, nil
	}
	defer func() {
		_ = f.Close()
	}()

	return l.extract(f), nil
}

func (l *Lines) extract(r io.Reader) []string {
	bf := bufio.NewScanner(r)
	var lines []string
	for bf.Scan() {
		line := bf.Text()
		if l.Condition(line) {
			extracted := l.Parser(line)
			lines = append(lines, extracted...)
		}
	}

	return lines
}
