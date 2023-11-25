package comicinfo

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	file, err := os.Open("Example.xml") // For read access.
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	v := ComicInfo{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Println(v)
}

func TestMarshal(t *testing.T) {
	file, err := os.Open("Example.xml") // For read access.
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	v := ComicInfo{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	// Due to XML unknown problem
	v.Xsi = "http://www.w3.org/2001/XMLSchema-instance"
	v.Xsd = "http://www.w3.org/2001/XMLSchema"

	output, err := xml.MarshalIndent(v, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	os.Stdout.Write([]byte("<?xml version=\"1.0\"?>\n"))

	os.Stdout.Write(output)
}
