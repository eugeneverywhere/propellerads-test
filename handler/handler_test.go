package handler

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"reflect"
	"sort"
	"testing"
)

type testCase struct {
	data     []string
	expected [][]string
}

func getTestCases() []*testCase {
	return []*testCase{
		{
			data: []string{
				"1",
				"2",
			},
			expected: [][]string{},
		},
		{
			data: []string{
				"1",
				"2,3",
			},
			expected: [][]string{
				{"3", "2"},
			},
		},
		{
			data: []string{
				"9",
				"1,5",
				"2,3",
			},
			expected: [][]string{
				{"5", "1"},
				{"2", "3"},
			},
		},
		{
			data: []string{
				"1,5",
				"9",
				"2,3",
			},
			expected: [][]string{
				{"1", "5"},
				{"2", "3"},
			},
		},
		{
			data: []string{
				"1,2",
				"3,4",
				"2,3",
			},
			expected: [][]string{
				{"1", "2", "3", "4"},
			},
		},
		{
			data: []string{
				"1,2,3",
				"11",
				"4,5,6",
				"13",
				"7,8,9",
				"14",
				"1,4,7",
				"101,102",
			},
			expected: [][]string{
				{"1", "2", "3", "4", "5", "6", "7", "8", "9"},
				{"101", "102"},
			},
		},
	}
}

func TestDataHandler_GetGroups(t *testing.T) {
	h := New()
	for j, tcase := range getTestCases() {
		groupedAccounts := h.GetGroups(tcase.data)
		checkResult(t, groupedAccounts, tcase.expected, j)
	}
}

func checkResult(t *testing.T, groupedAccounts [][]string, expected [][]string, caseNumber int) {
	assert.Equal(t, len(groupedAccounts), len(expected), fmt.Sprintf("Test case %v", caseNumber))
	for i := 0; i < len(groupedAccounts); i++ {
		found := false
		sort.Strings(groupedAccounts[i])
		for _, exp := range expected {
			sort.Strings(exp)
			found = reflect.DeepEqual(exp, groupedAccounts[i])
			if found {
				break
			}
		}
		assert.Equal(t, found, true, fmt.Sprintf("Test case %v: Not expected %v ", caseNumber, groupedAccounts[i]))
	}
}
