package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type ClosingBuffer struct {
	*bytes.Buffer
}

func (cb *ClosingBuffer) Close() (err error) {
	//we don't actually have to do anything here, since the buffer is just some data in memory
	//and the error is initialized to no-error
	return
}

func main() {
	cb := &ClosingBuffer{bytes.NewBufferString("Hi GoCloudMember!")}
	var rc io.ReadCloser
	rc = cb

	buf := new(bytes.Buffer)

	numOfByte, err := buf.ReadFrom(rc)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	content := buf.String()
	fmt.Printf("Read: %d bytes, content is: %q\r\n", numOfByte, content)

	fmt.Printf("Type of cb is %T", cb)
	fmt.Println("")
	fmt.Printf("Type of rc is %T", rc)
	fmt.Println("")
	rc.Close()
}