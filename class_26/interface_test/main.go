package main

import (
    "fmt"
    "io"
    "net/http"
)

type BodyEOFSignal struct {
    reader io.ReadCloser
    isEOF  bool
}

func (b *BodyEOFSignal) Read(p []byte) (n int, err error) {
    if b.isEOF {
        return 0, io.EOF
    }

    n, err = b.reader.Read(p)
    if err != nil {
        if err != io.EOF {
            return n, err
        }

        b.isEOF = true
    }

    return n, nil
}

func (b *BodyEOFSignal) Close() error {
    if b.reader != nil {
        return b.reader.Close()
    }

    return nil
}

func (b *BodyEOFSignal) IsEOF() bool {
    return b.isEOF
}

func main() {
    resp, err := http.Get("https://www.example.com")
    if err != nil {
        fmt.Println(err)
        return
    }

    body := &BodyEOFSignal{
        reader: resp.Body,
    }

    for {
        buf := make([]byte, 1024)
        n, err := body.Read(buf)
        if err != nil {
            if err == io.EOF {
                fmt.Println("End of body reached")
                break
            }
            fmt.Println(err)
            return
        }

        fmt.Println(string(buf[:n]))
    }
}