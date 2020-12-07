package scan

import "unicode/utf8"

// dropCR drops a terminal \r from the data.
// Code from bufio package.
func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}

func Groups(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	// Scan until single blank line, marking end of a group.
	for width, i := 0, 0; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])
		if d := dropCR(data[:i]); r == '\n' && d[i-1] == '\n' {
			return i + width, dropCR(data[0 : i-1]), nil
		}
	}

	if atEOF {
		return len(data), dropCR(data), nil
	}

	// Request more data.
	return 0, nil, nil
}
