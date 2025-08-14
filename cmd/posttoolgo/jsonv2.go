//go:build goexperiment.jsonv2

package posttoolgo

import "os"

func init() {
	os.Setenv("GOEXPERIMENT", "jsonv2")
}
