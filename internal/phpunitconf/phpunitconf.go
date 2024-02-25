package phpunitconf

import (
	"fmt"
	"os"
)

func Parse(filepath string) (any, error) {
	fmt.Println("Hello there", filepath)

	s, err := os.Stat(filepath)
	if err != nil {
		return nil, err
	}

	if s.IsDir() {
		return nil, fmt.Errorf("config should be a file, dir <%s> given", filepath)
	}

	// TODO read phpunit config
	// Validate against his XSD

	return nil, err
}
