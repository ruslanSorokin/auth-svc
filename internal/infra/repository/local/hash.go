package local

import (
	"bytes"
	"encoding/gob"
)

func Hash(s interface{}) string {
	var b bytes.Buffer
	gob.NewEncoder(&b).Encode(s)
	return b.String()
}
