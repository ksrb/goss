//http://tour.golang.org/methods/11
package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (m MyReader) Read(input []byte) (int, error) {
	for index := range input {
		input[index] = 'A'
	}
	return len(input), nil
}

func main() {
	reader.Validate(MyReader{})
}
