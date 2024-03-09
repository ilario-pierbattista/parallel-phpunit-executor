package phpunitconf

import (
	"encoding/xml"
	"io"
	"log"
)

func RemoveNodeByName(
	reader io.Reader,
	writer io.Writer,
	field string,
) error {
	decoder := xml.NewDecoder(reader)
	encoder := xml.NewEncoder(writer)
	defer encoder.Close()

	shouldEncode := true
	for {
		token, err := decoder.Token()
		encodeThisToken := true

		if err != nil {
			log.Println(err)
		}
		if token == nil {
			break
		}

		switch element := token.(type) {
		case xml.StartElement:
			if element.Name.Local == field {
				shouldEncode = false
				encodeThisToken = false
			}
		case xml.EndElement:
			if element.Name.Local == field {
				shouldEncode = true
				encodeThisToken = false
			}
		}

		if shouldEncode && encodeThisToken {
			err = encoder.EncodeToken(token)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	err := encoder.Flush()
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
