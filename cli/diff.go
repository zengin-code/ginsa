package main

import (
	"github.com/briandowns/spinner"
	"github.com/urfave/cli"
	"github.com/zengin-code/ginsa"
	"io"
	"os"
	"sort"
	"time"
)

var diffCmd = cli.Command{
	Name:    "diff",
	Aliases: []string{"d"},
	Usage:   "Show diff",
	Flags:   []cli.Flag{},
	Action:  diffAction,
}

func diffAction(c *cli.Context) {
	tags := c.Args()
	sort.Strings(tags)

	s := spinner.New(spinner.CharSets[21], 100*time.Millisecond)
	s.Suffix = "  Fetch tags by GitHub..."
	s.Writer = os.Stderr
	s.Start()

	dmap := map[string]*ginsa.SourceData{}
	data, err := ginsa.FetchAllSourceData()
	if err != nil {
		panic(err)
	}
	for _, d := range data {
		dmap[d.Tag] = d
	}

	s.Stop()
	io.WriteString(os.Stderr, "\n")

	if len(tags) == 0 {
		tags = []string{data[len(data)-1].Tag}
	}

	if len(tags) == 1 {
		prevTag := ""
		for _, d := range data {
			if d.Tag == tags[0] {
				tags[0] = prevTag
				tags = append(tags, d.Tag)
				break
			}

			prevTag = d.Tag
		}
	}

	s = spinner.New(spinner.CharSets[21], 100*time.Millisecond)
	s.Suffix = "  Fetch data by GitHub..."
	s.Writer = os.Stderr
	s.Start()

	for _, tag := range tags {
		d := dmap[tag]
		err := d.Load()
		if err != nil {
			panic(err)
		}
	}

	s.Stop()
	io.WriteString(os.Stderr, "\n")

	s = spinner.New(spinner.CharSets[21], 100*time.Millisecond)
	s.Suffix = "  Calculate diff..."
	s.Writer = os.Stderr
	s.Start()

	diffs := ginsa.DiffSourceData(dmap[tags[0]], dmap[tags[1]])

	s.Stop()
	io.WriteString(os.Stderr, "\n")

	diffs.Out()
}
