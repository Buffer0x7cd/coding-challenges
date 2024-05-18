package main

import (
	"bytes"
	"testing"

)

type countFileByBytesTest struct {
	arg1 []byte
}

type countFileByWordsTest struct {
	arg1 []byte
	want int
}

type countFileByLinesTest struct {
	arg1 []byte
	want int
}

var countFileByBytesTests = []countFileByBytesTest{
	{[]byte("Hello, World!")},
	{[]byte("This is a test line")},
	{[]byte("this is another test line")},
}

var countFileByWordsTests = []countFileByWordsTest{
	{[]byte("Hello, World!"), 2},
	{[]byte("Hello"), 1},
	{[]byte("This line is 4 world long"), 6},
}

var countFileByLinesTests = []countFileByLinesTest{
	{[]byte("Line1"), 1},
	{[]byte("Line1 \n Line2"), 2},
	{[]byte(""), 0},

}

func TestCountFileByBytes(t *testing.T) {
	for _, test := range countFileByBytesTests {
		reader := bytes.NewReader(test.arg1)
		got, _ := CountFileByBytes(reader)
		if got != len(test.arg1) {
			t.Errorf("countFileByBytes(%s) = %d; want %d", test.arg1, got, len(test.arg1))
		}
	}
}


func TestCountFileByWords(t *testing.T) {
	for _, test := range countFileByWordsTests {
		reader := bytes.NewReader(test.arg1)
		got, _ := CountFileByWords(reader)
		if got != test.want{
			t.Errorf("countFileByWords(%s) = %d; want %d", test.arg1, got, test.want)	
		}
	}
}



func TestCountFileByLines(t *testing.T) {
	for _, test := range countFileByLinesTests{
		reader := bytes.NewReader(test.arg1)
		got, _ := CountFileByLines(reader)
		if got != test.want{
			t.Errorf("countFileByLines(%s) = %d; want %d", test.arg1, got, test.want)	
		}
	}
}