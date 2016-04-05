package parser

import (
    "testing"
    "bufio"
    "bytes"
)

// test scanning an integer
func TestParseInt(t *testing.T) {

    s := bufio.NewScanner(bytes.NewReader([]byte{' ', '1', '2', '3', '\t'}))
    s.Split(bufio.ScanWords)

    if n, err := parseInt(s); err != nil {
        panic(err)
    } else if n != 123 {
        t.Errorf("Expected 123, got %d.", n)
    }
}

// test failing to scan an integer
func TestParseIntError(t *testing.T) {

    s := bufio.NewScanner(bytes.NewReader([]byte{' ', 'h', 'i', '\n', '\t'}))
    s.Split(bufio.ScanWords)

    expect := "strconv.ParseInt: parsing \"hi\": invalid syntax"
    if _, err := parseInt(s); err == nil {
        t.Error("Expected error, got nil.")
    } else if err.Error() != expect {
        t.Errorf("Expected error \"%s\", got \"%s\".", expect, err.Error())
    }

    if _, err := parseInt(s); err == nil {
        t.Error("Expected error, got nil.")
    } else if err.Error() != "EOF" {
        t.Errorf("Expected error \"EOF\", got \"%s\".", err.Error())
    }
}

// test scanning an integer
func TestParseFloat(t *testing.T) {

    s := bufio.NewScanner(bytes.NewReader([]byte{' ', '1', '.', '2', '3', '\t'}))
    s.Split(bufio.ScanWords)

    if n, err := parseFloat(s); err != nil {
        panic(err)
    } else if n != 1.23 {
        t.Errorf("Expected 1.23, got %f.", n)
    }
}

// test failing to scan an integer
func TestParseFloatError(t *testing.T) {

    s := bufio.NewScanner(bytes.NewReader([]byte{' ', 'h', 'i', '\n', '\t'}))
    s.Split(bufio.ScanWords)

    expect := "strconv.ParseFloat: parsing \"hi\": invalid syntax"
    if _, err := parseFloat(s); err == nil {
        t.Error("Expected error, got nil.")
    } else if err.Error() != expect {
        t.Errorf("Expected error \"%s\", got \"%s\".", expect, err.Error())
    }

    if _, err := parseFloat(s); err == nil {
        t.Error("Expected error, got nil.")
    } else if err.Error() != "EOF" {
        t.Errorf("Expected error \"EOF\", got \"%s\".", err.Error())
    }
}