package search

import (
	"fmt"
	"testing"
)

func TestExtractPath(t *testing.T) {
	root := "../../../"
	files, paths, err := ExtractPath(root)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("files:")
	for _, file := range files {
		fmt.Println(file)
	}
	fmt.Println("paths:")
	for _, path := range paths {
		fmt.Println(path)
	}
}
