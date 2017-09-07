package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// Add a Reader([]byte)(int, error) method to MyReader.

func (reader MyReader) Read(array []byte) (int, error) {
	size := len(array)
	for i := 0; i < size; i++ {
		array[i] = 'A'
	}
	return size, nil
}

func main() {
	reader.Validate(MyReader{})
}
