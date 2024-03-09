package phpunitconf

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"

	"github.com/ilario-pierbattista/parallel-phpunit-executor/internal/common"
)

type Testsuite struct {
	Name        string   `xml:"name,attr"`
	Directories []string `xml:"directory"`
	Files       []string `xml:"file"`
}

func MakeTestSuite(chunk common.Chunk) Testsuite {
	return Testsuite{
		Name:        chunk.Uuid.String(),
		Directories: []string{},
		Files:       chunk.Files,
	}
}

type PHPUnitPartialConfig struct {
	XMLName    xml.Name `xml:"phpunit"`
	Inner      []byte   `xml:",innerxml"`
	Testsuites struct {
		Testsuite `xml:"testsuite"`
	} `xml:"testsuites"`
}

func Parse(filepath string) (any, error) {
	fmt.Println("Hello there", filepath)

	s, err := os.Stat(filepath)
	if err != nil {
		return nil, err
	}

	if s.IsDir() {
		return nil, fmt.Errorf("config should be a file, dir <%s> given", filepath)
	}

	xmlFile, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer xmlFile.Close()

	byteValue, err := io.ReadAll(xmlFile)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(byteValue))

	data := PHPUnitPartialConfig{}
	err = xml.Unmarshal(byteValue, &data)
	if err != nil {
		return nil, err
	}

	fmt.Println(data)

	bytesDone, err := xml.MarshalIndent(data, "", "  ")
	if err != nil {
		return nil, err
	}
	fmt.Println(string(bytesDone))

	var bytesStripped []byte = []byte{}

	fmt.Println("Another approach")
	buffer := bytes.NewBuffer(bytesStripped)
	fmt.Println(buffer.Len())

	RemoveNodeByName(bytes.NewBuffer(byteValue), buffer, "testsuites")

	fmt.Println(buffer.String())
	fmt.Println("Fine approccio")

	// TODO read phpunit config
	// Validate against his XSD

	return nil, err
}
