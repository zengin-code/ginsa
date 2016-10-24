package main

import (
	"fmt"
	"github.com/briandowns/spinner"
	"github.com/urfave/cli"
	"github.com/zengin-code/ginsa"
	"io"
	"os"
	"time"
)

var tagsCmd = cli.Command{
	Name:    "tags",
	Aliases: []string{"d"},
	Usage:   "Show tags",
	Flags:   []cli.Flag{},
	Action:  tagsAction,
}

func tagsAction(c *cli.Context) {
	s := spinner.New(spinner.CharSets[21], 100*time.Millisecond)
	s.Suffix = "  Fetch tags by GitHub..."
	s.Writer = os.Stderr
	s.Start()

	data, err := ginsa.FetchAllSourceData()
	if err != nil {
		panic(err)
	}

	s.Stop()
	io.WriteString(os.Stderr, "\n")

	for _, d := range data {
		fmt.Println(d.Tag)
	}
}
