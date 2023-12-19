package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {

	var w Writer = ConsoleWriter{}
	_, err := w.Write([]byte("Hello Go!"))
	if err != nil {
		return
	}

	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}

	var wc WriterCloser = NewBufferedWriterCloser()
	_, err = wc.Write([]byte("Hello go developers, I`m still learning"))
	if err != nil {
		return
	}
	err = wc.Close()
	if err != nil {
		return
	}

	// Here we successfully cast the interface to a concrete type
	bwc := wc.(BufferedWriterCloser)
	fmt.Println(bwc)

	r, ok := wc.(io.Reader)
	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("Conversion failed")
	}

	var emptyInterface interface{} = 0
	switch emptyInterface.(type) {
	case int:
		fmt.Println("emptyInterface is an int")
	case string:
		fmt.Println("emptyInterface is a string")
	default:
		fmt.Println("emptyInterface is another type")
	}

}

// Writer In Go we don't need to explicitly say that a type implements an interface
// as long as the type has the methods that the interface requires
// the type is said to implement the interface
type Writer interface {
	Write([]byte) (int, error)
}

type Closer interface {
	Close() error
}

type WriterCloser interface {
	Writer
	Closer
}
type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc BufferedWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}
	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}
	return n, nil
}

func (bwc BufferedWriterCloser) Close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

func NewBufferedWriterCloser() BufferedWriterCloser {
	return BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}

type Incrementer interface {
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}
