package ginsa

import (
	"github.com/k0kubun/pp"
	"testing"
)

func TestFetchAllSourceData(t *testing.T) {
	old := &SourceData{Tag: "2016-09-01"}
	now := &SourceData{Tag: "2016-09-15"}

	err := old.Load()
	if err != nil {
		pp.Println(err)
	}
	err = now.Load()
	if err != nil {
		pp.Println(err)
	}

	diffs := DiffSourceData(old, now)
	diffs.Out()
}
