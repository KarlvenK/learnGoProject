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
	data, _ := ReadFrom(os.Stdin, 3)
	fmt.Println(data)
	data, _ = ReadFrom(strings.NewReader("from string"), 12)
	fmt.Println(data)

	//io.Writer
	/*
		type Writer interface {
			Write(p []byte) (n int, err error)
		}
	*/
	_, _ = Println(1, 2, 3, 4, 5)

	tryOSfile()

}

func tryOSfile() {
	file, err := os.Create("writeAt.txt")
	if err != nil {
		panic(err)
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	file.WriteString("this is my demo")

	n, err := file.WriteAt([]byte("golang is cool"), 10)
	if err != nil {
		panic(err)
	}
	_, _ = Println(n)
}

func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}
	return p, err
}

func Println(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(os.Stdout, a...)
}
