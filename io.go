package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func tryIO(tttt int) {
	if tttt == 0 {
		return
	}

	tryReadAndWrite()

	tryOSfile()

	tryFromAndTo()

	trySeeker()

	tryCloser()

	tryPipe()
}

func tryPipe() {
	pipeReader, pipeWriter := io.Pipe()
	go PipeWrite(pipeWriter)
	go PipeRead(pipeReader)
	time.Sleep(1e7)
}

func PipeWrite(pipeWriter *io.PipeWriter) {
	var (
		i   = 0
		err error
		n   int
	)
	data := []byte("golang test")
	for _, err = pipeWriter.Write(data); err == nil; n, err = pipeWriter.Write(data) {
		i++
		if i == 3 {
			err = pipeWriter.CloseWithError(errors.New("输出三次停止"))
		}
	}
	fmt.Println("close后输出的字节数： ", n, " error = ", err)
}

func PipeRead(pipeReader *io.PipeReader) {
	var (
		err error
		n   int
	)
	data := make([]byte, 1024)
	for n, err = pipeReader.Read(data); err == nil; n, err = pipeReader.Read(data) {
		fmt.Printf("%s\n", data[:n])
	}
	fmt.Println("writer端 closewitherror后：", err)
}

func tryCloser() {
	file, err := os.Open("receivor.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
}

func trySeeker() {
	reader := strings.NewReader("goLangTest")
	reader.Seek(-1, os.SEEK_END)
	r, _, _ := reader.ReadRune()
	fmt.Printf("%c\n", r)
}

func tryFromAndTo() {
	file, err := os.Open("writeAt.txt")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	anotherFile, err := os.Create("receivor.txt")
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			panic(err)
		}
	}(anotherFile)

	writer := bufio.NewWriter(anotherFile)
	_, _ = writer.ReadFrom(file)
	_ = writer.Flush()

	reader := bytes.NewReader([]byte("this is a func\n"))
	_, _ = reader.WriteTo(os.Stdout)
}

func tryReadAndWrite() {
	// io.Reader
	/*
		type Reader interface {
			Read(p []byte) (n int, err error)
		}
	*/
	data, _ := ReadFrom(os.Stdin, 0) // change "0" to another number if you really want to read sth
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
}

func tryOSfile() {
	file, err := os.Create("writeAt.txt")
	if err != nil {
		panic(err)
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	_, _ = file.WriteString("this is my demo")

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
