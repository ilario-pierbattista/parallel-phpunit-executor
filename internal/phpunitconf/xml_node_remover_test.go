package phpunitconf

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXmlNodeRemover(t *testing.T) {

	t.Run("Remove node by name", func(t *testing.T) {

		inputXml := strings.TrimSpace(`
<?xml version="1.0"?>
<customers>
<customer id="55000">
<name>Charter Group</name>
<address>
<street>100 Main</street>
<city>Framingham</city>
<state>MA</state>
<zip>01701</zip>
</address>
<address>
<street>720 Prospect</street>
<city>Framingham</city>
<state>MA</state>
<zip>01701</zip>
</address>
<address>
<street>120 Ridge</street>
<state>MA</state>
<zip>01760</zip>
</address>
</customer>
</customers>
`)

		outBuffer := bytes.NewBufferString("")

		err := RemoveNodeByName(
			strings.NewReader(inputXml),
			outBuffer,
			"address",
		)
		if err != nil {
			t.Fail()
		}

		expectedOutput := strings.TrimSpace(`
<?xml version="1.0"?>
<customers>
<customer id="55000">
<name>Charter Group</name>



</customer>
</customers>
`)

		assert.Equal(t, expectedOutput, outBuffer.String())
	})

}
