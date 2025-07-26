package hooks

import (
	"encoding/json"
	"io"
)

func ReadInput[TI any, TR any](r io.Reader) (out Input[TI, TR], err error) {
	err = json.NewDecoder(r).Decode(&out)
	return
}
