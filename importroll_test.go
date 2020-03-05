package importroll

import (
	"testing"

	"github.com/gobwas/glob"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	rule = "./testdata/src/github.com/wreulicke/sample/importroll.yaml"
	analysistest.Run(t, testdata, Analyzer, "github.com/wreulicke/sample/...")
}

func TestGlob(t *testing.T) {
	if !glob.MustCompile("*/controller*").Match("github.com/wreulicke/sample/controller") {
		t.Error("expected match")
	}
}
