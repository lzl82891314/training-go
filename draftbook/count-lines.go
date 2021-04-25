package draftbook

import (
	"bufio"
	"io"
)

func CountLines(r io.Reader) (int, error) {
	br := bufio.NewReader(r)
	var lines int
	var err error

	for {
		_, err = br.ReadString('\n')
		lines++
		if err != nil {
			break
		}
	}
	if err != io.EOF {
		return 0, err
	}
	return lines, nil
}

func CountLInes_v2(r io.Reader) (int, error) {
	sc := bufio.NewScanner(r)
	lines := 0
	for sc.Scan() {
		lines++
	}
	return lines, sc.Err()
}
