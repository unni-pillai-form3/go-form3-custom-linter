package test

import (
	"github.com/go-form3-custom-linter/pkg/src/constantarg"
	"github.com/go-form3-custom-linter/pkg/src/timeusage"
	"os"
	"path/filepath"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestConstantArgumentAnalyzer(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get wd: %s", err)
	}

	testdata := filepath.Join(filepath.Dir(filepath.Dir(wd)), "testdata/constantarg")
	analysistest.Run(t, testdata, constantarg.ConstantArgumentAnalyzer)
}

func TestTimeUsageAnalyzer(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get wd: %s", err)
	}

	testdata := filepath.Join(filepath.Dir(filepath.Dir(wd)), "testdata/timeusage")
	analysistest.Run(t, testdata, timeusage.TimeAnalyzer)

}
