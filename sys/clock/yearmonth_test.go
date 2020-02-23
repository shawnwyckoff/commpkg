package clock

import (
	"testing"
)

func TestParseYearMonthInt(t *testing.T) {
	type testitem struct {
		input          int
		expectedOutput string
	}

	testitems := []testitem{
		testitem{input: 200007, expectedOutput: "200007"},
		testitem{input: -200007, expectedOutput: "-200007"},
	}

	for _, v := range testitems {
		output, err := ParseYearMonthInt(v.input)
		if err != nil {
			t.Error(err)
			return
		}
		if output.StringYYYYMM() != v.expectedOutput {
			t.Errorf("output %s, expected output %s", output, v.expectedOutput)
			return
		}
	}
}
