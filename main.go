package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/olekukonko/tablewriter"
)

type table struct {
	Header []string   `json:"header"`
	Data   [][]string `json:"data"`
}

func exitError(msg string) {
	fmt.Fprintf(os.Stderr, "%s\n", msg)
	os.Exit(1)
}

func main() {
	args := os.Args
	argsSize := len(args) - 1
	if argsSize != 1 {
		exitError(fmt.Sprintf("unexpexcted arguments size, expected 1, but got %d", argsSize))
	}

	path := args[1]
	b, err := ioutil.ReadFile(path)
	if err != nil {
		exitError(err.Error())
	}

	t := &table{}
	err = json.Unmarshal(b, t)
	if err != nil {
		exitError(fmt.Sprintf("failed to parse %s as JSON, %s", path, err))
	}

	tw := tablewriter.NewWriter(os.Stdout)
	tw.SetHeader(t.Header)
	tw.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	tw.SetCenterSeparator("|")
	tw.AppendBulk(t.Data)
	tw.Render()
}
