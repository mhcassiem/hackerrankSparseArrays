package tests

import (
	"github.com/mhcassiem/hackerrankSparseArrays/pkg"
	"testing"
)

func TestMatchingStrings(t *testing.T) {

	expected := []int32{2, 1, 0}
	inputData := []string{"ab", "ab", "abc"}
	inputQueries := []string{"ab", "abc", "bc"}

	results, valid := runTestCase([]string{"ab", "ab", "abc"}, []string{"ab", "abc", "bc"}, []int32{2, 1, 0})
	if !valid {
		t.Errorf("Expected %d, received %d", expected, results)
	}

	expected = []int32{2, 1, 0}
	inputData = []string{"aba", "baba", "aba", "xzxb"}
	inputQueries = []string{"aba", "xzxb", "ab"}

	results, valid = runTestCase(inputData, inputQueries, expected)
	if !valid {
		t.Errorf("Expected %d, received %d", expected, results)
	}

	expected = []int32{1, 0, 1}
	inputData = []string{"def", "de", "fgh"}
	inputQueries = []string{"de", "lmn", "fgh"}

	results, valid = runTestCase(inputData, inputQueries, expected)
	if !valid {
		t.Errorf("Expected %d, received %d", expected, results)
	}

	expected = []int32{1, 3, 4, 3, 2}
	inputData = []string{"abcde", "sdaklfj", "asdjf", "na", "basdn", "sdaklfj", "asdjf", "na", "asdjf", "na", "basdn", "sdaklfj", "asdjf"}
	inputQueries = []string{"abcde", "sdaklfj", "asdjf", "na", "basdn"}

	results, valid = runTestCase(inputData, inputQueries, expected)
	if !valid {
		t.Errorf("Expected %d, received %d", expected, results)
	}

}

func runTestCase(inputData []string, inputQueries []string, expected []int32) ([]int32, bool) {
	var results = pkg.MatchingStrings(inputData, inputQueries)
	var valid = false

	if len(results) == len(expected) {
		for index, value := range expected {
			if value == results[index] {
				valid = true
			} else {
				valid = false
			}
		}
	} else {
		valid = false
	}
	return results, valid
}
