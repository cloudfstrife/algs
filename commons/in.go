package commons

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"unicode/utf8"
)

// In input struct
type In struct {
	scanner *bufio.Scanner
}

//NewIn Initializes an input from a reader.
func NewIn(reader io.Reader) In {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	return In{scanner: scanner}
}

//IsEmpty return true if input is empty
func (in In) IsEmpty() bool {
	return len(in.scanner.Bytes()) == 0
}

//ReadString read line
func (in In) ReadString() string {
	v := in.scanner.Text()
	_ = in.scanner.Scan()
	return v
}

//ReadAllString read all lines as string slice
func (in In) ReadAllString() []string {
	var result []string
	for !in.IsEmpty() {
		line := in.scanner.Text()
		result = append(result, line)
		_ = in.scanner.Scan()
	}
	return result
}

// ReadInteger read the next line from input stream, parses it as a int.
func (in In) ReadInteger() (int, error) {
	v := strings.TrimSpace(in.ReadString())
	return strconv.Atoi(v)
}

// ReadAllInteger reads all lines from input stream, parses all of those as a integer slice.
func (in In) ReadAllInteger() []int {
	result := []int{}
	for {
		i, err := in.ReadInteger()
		if err != nil {
			break
		}
		result = append(result, i)
	}
	return result
}

// ReadBool read the next line from input stream, parses it as a bool .
// if some line can not parse as bool , will return all error
func (in In) ReadBool() (bool, error) {
	v := strings.TrimSpace(in.ReadString())
	switch v {
	case "0":
		return false, nil
	case "false":
		return false, nil
	case "1":
		return true, nil
	case "true":
		return true, nil
	default:
		return false, fmt.Errorf("attempts to read a 'bool' value , but the next token is %#v ", v)
	}
}

// ReadAllBool read all lines from input stream, parses all of those as a bool slice.
// if some line can not parse as bool , will be ignored
func (in In) ReadAllBool() []bool {
	result := []bool{}
	for {
		b, err := in.ReadBool()
		if err != nil {
			break
		}
		result = append(result, b)
	}
	return result
}

// ReadRune read the next line from input stream, return first rune.
// return a error while decode line as rune slice failed
func (in In) ReadRune() (rune, int, error) {
	r, l := utf8.DecodeRuneInString(in.ReadString())
	if r == utf8.RuneError {
		return 0, 0, errors.New("internal In.ReadRune() error")
	}
	return r, l, nil
}

// ReadAllRune read all line from input stream , return a rune slice contains rune slice.
func (in In) ReadAllRune() []rune {
	result := []rune{}
	for !in.IsEmpty() {
		line := in.scanner.Text()
		result = append(result, []rune(line)...)
		_ = in.scanner.Scan()
	}
	return result
}

// ReadFloat read the next line from input stream, parses it as a float64.
func (in In) ReadFloat() (float64, error) {
	v := strings.TrimSpace(in.ReadString())
	return strconv.ParseFloat(v, 64)
}

// ReadAllFloat reads all lines from input stream, parses all of those as a float64 slice.
func (in In) ReadAllFloat() []float64 {
	result := []float64{}
	for {
		f, err := in.ReadFloat()
		if err != nil {
			break
		}
		result = append(result, f)
	}
	return result
}
