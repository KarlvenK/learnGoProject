package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func tryIO(tttt int) {
	if tttt == 0 {
		return
	}
	// io.Reader
	/*
		type Reader interface {
			Read(p []byte) (n int, err error)
		}
	*/
	data, _ := ReadFrom(os.Stdin, 11)
	fmt.Println(data)
	data, _ = ReadFrom(strings.NewReader("from string"), 12)
	fmt.Println(data)

	//io.Writer
	/*
		type Writer interface {
			Write(p []byte) (n int, err error)
		}
	*/

}

func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}
	return p, err
}
