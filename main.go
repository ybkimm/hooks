package main

import (
	"context"
	"os"

	"github.com/charmbracelet/fang"
	"github.com/ybkimm/hooks/cmd"

	_ "github.com/ybkimm/hooks/cmd/posttoolgo"
)

func main() {
	err := fang.Execute(context.Background(), cmd.Get())
	if err != nil {
		os.Exit(1)
	}
}
