package ini

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"unicode"
)

func removeComments(value string) string {
	pos := strings.LastIndexAny(value, ";#")

	//if no inline comments
	if pos == -1 || !unicode.IsSpace(rune(value[pos-1])) {
		return value
	}
	return strings.TrimSpace(value[0:pos])
}

func isOctChar(ch byte) bool {
	return ch >= '0' && ch <= '7'
}

func isHexChar(ch byte) bool {
	return ch >= '0' && ch <= '9' ||
		ch >= 'a' && ch <= 'f' ||
		ch >= 'A' && ch <= 'F'
}

func fromEscape(value string) string {
	if strings.Index(value, "\\") == -1 {
		return value
	}

	r := ""
	n := len(value)
	for i := 0; i < n; i++ {
		if value[i] == '\\' {
			if i+1 < n {
				i++
				//if is it oct
				if i+2 < n && isOctChar(value[i]) && isOctChar(value[i+1]) && isOctChar(value[i+2]) {
					t, err := strconv.ParseInt(value[i:i+3], 8, 32)
					if err == nil {
						r = r + string(rune(t))
					}
					i += 2
					continue
				}
				switch value[i] {
				case '0':
					r = r + string(byte(0))
				case 'a':
					r = r + "\a"
				case 'b':
					r = r + "\b"
				case 'f':
					r = r + "\f"
				case 't':
					r = r + "\t"
				case 'r':
					r = r + "\r"
				case 'n':
					r = r + "\n"
				case 'v':
					r = r + "\v"
				case 'x':
					if i+3 < n && isHexChar(value[i]) &&
						isHexChar(value[i+1]) &&
						isHexChar(value[i+2]) &&
						isHexChar(value[i+3]) {

						t, err := strconv.ParseInt(value[i:i+4], 16, 32)
						if err == nil {
							r = r + string(rune(t))
						}
						i += 3
					}
				default:
					r = fmt.Sprintf("%s%c", r, value[i])
				}
			}
		} else {
			r = fmt.Sprintf("%s%c", r, value[i])
		}
	}
	return r
}

func removeContinuationSuffix(value string) (string, bool) {
	pos := strings.LastIndex(value, "\\")
	n := len(value)
	if pos == -1 || pos != n-1 {
		return "", false
	}
	for pos >= 0 {
		if value[pos] != '\\' {
			return "", false
		}
		pos--
		if pos < 0 || value[pos] != '\\' {
			return value[0 : n-1], true
		}
		pos--
	}
	return "", false
}

type LineReader struct {
	reader *bufio.Scanner
}

func NewLineReader(reader io.Reader) *LineReader {
	return &LineReader{reader: bufio.NewScanner(reader)}
}

func (lr *LineReader) ReadLine() (string, error) {
	if lr.reader.Scan() {
		return lr.reader.Text(), nil
	}
	return "", errors.New("No data")

}

func readLinesUntilSuffix(lineReader *LineReader, suffix string) string {
	r := ""
	for {
		line, err := lineReader.ReadLine()
		if err != nil {
			break
		}
		t := strings.TrimRightFunc(line, unicode.IsSpace)
		if strings.HasSuffix(t, suffix) {
			r = r + t[0:len(t)-len(suffix)]
			break
		} else {
			r = r + line + "\n"
		}
	}
	return r
}

func readContinuationLines(lineReader *LineReader) string {
	r := ""
	for {
		line, err := lineReader.ReadLine()
		if err != nil {
			break
		}
		line = strings.TrimRightFunc(line, unicode.IsSpace)
		if t, continuation := removeContinuationSuffix(line); continuation {
			r = r + t
		} else {
			r = r + line
			break
		}
	}
	return r
}

func (ini *Ini) Load(reader io.Reader) {
	lineReader := NewLineReader(reader)
	var curSection *Section = nil
	for {
		line, err := lineReader.ReadLine()
		if err != nil {
			break
		}
		line = strings.TrimSpace(line)

		//empty line or comments line
		if len(line) <= 0 || line[0] == ';' || line[0] == '#' {
			continue
		}
		//if it is a section
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			if len(sectionName) > 0 {
				curSection = ini.NewSection(sectionName)
			}
			continue
		}
		pos := strings.Index(line, "=")
		if pos != -1 {
			key := strings.TrimSpace(line[0:pos])
			value := strings.TrimLeftFunc(line[pos+1:], unicode.IsSpace)
			if strings.HasPrefix(value, "\"\"\"") {
				t := strings.TrimRightFunc(value, unicode.IsSpace)
				if strings.HasSuffix(t, "\"\"\"") {
					value = t[3 : len(t)-3]
				} else {
					value = value[3:] + "\n" + readLinesUntilSuffix(lineReader, "\"\"\"")
				}
			} else {
				value = strings.TrimRightFunc(value, unicode.IsSpace)
				if t, continuation := removeContinuationSuffix(value); continuation {
					value = t + readContinuationLines(lineReader)
				}
			}

			if len(key) > 0 && curSection != nil {
				curSection.Add(key, fromEscape(strings.TrimSpace(value)))
			}
		}
	}
}
