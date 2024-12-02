package year_2024_test

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
	"testing"
)

type Report struct {
	Levels []int
}

func (r Report) Candidates() []Report {
	var candidates []Report
	for i := 0; i < len(r.Levels); i++ {
		candidate := Report{Levels: append([]int{}, r.Levels[:i]...)}
		candidate.Levels = append(candidate.Levels, r.Levels[i+1:]...)
		candidates = append(candidates, candidate)
	}
	return candidates
}

func isSafe(report Report) bool {
	increasing := report.Levels[0] < report.Levels[1]
	diff := int(math.Abs(float64(report.Levels[0] - report.Levels[1])))
	if diff < 1 || diff > 3 {
		return false
	}
	for i := 1; i < len(report.Levels)-1; i++ {
		diff = int(math.Abs(float64(report.Levels[i+1] - report.Levels[i])))
		localIncreasing := report.Levels[i] < report.Levels[i+1]
		if localIncreasing != increasing {
			return false
		}
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

func isAnyCandidateSafe(report Report) bool {
	candidates := report.Candidates()
	for _, candidate := range candidates {
		if isSafe(candidate) {
			return true
		}
	}
	return false
}

func CountSafeReports(reports []Report) int {
	count := 0
	for _, report := range reports {
		if isSafe(report) {
			count++
		}
	}
	return count
}

func CountSafeReportsWithCandidates(reports []Report) int {
	count := 0
	for _, report := range reports {
		if isSafe(report) || isAnyCandidateSafe(report) {
			count++
		}
	}
	return count
}

func TestCountSafeReports(t *testing.T) {
	tests := map[string]struct {
		reports  []Report
		expected int
	}{
		"advent_example": {
			reports: []Report{
				{Levels: []int{7, 6, 4, 2, 1}},
				{Levels: []int{1, 2, 7, 8, 9}},
				{Levels: []int{9, 7, 6, 2, 1}},
				{Levels: []int{1, 3, 2, 4, 5}},
				{Levels: []int{8, 6, 4, 4, 1}},
				{Levels: []int{1, 3, 6, 7, 9}},
			},
			expected: 2,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actual := CountSafeReports(tc.reports)
			if actual != tc.expected {
				t.Errorf("expected: %v, got: %v", tc.expected, actual)
			}
		})
	}
}

func TestCountSafeReportsWithCandidates(t *testing.T) {
	tests := map[string]struct {
		reports  []Report
		expected int
	}{
		"advent_example": {
			reports: []Report{
				{Levels: []int{7, 6, 4, 2, 1}},
				{Levels: []int{1, 2, 7, 8, 9}},
				{Levels: []int{9, 7, 6, 2, 1}},
				{Levels: []int{1, 3, 2, 4, 5}},
				{Levels: []int{8, 6, 4, 4, 1}},
				{Levels: []int{1, 3, 6, 7, 9}},
			},
			expected: 4,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actual := CountSafeReportsWithCandidates(tc.reports)
			if actual != tc.expected {
				t.Errorf("expected: %v, got: %v", tc.expected, actual)
			}
		})
	}
}

func TestCountSafeReportsInput(t *testing.T) {
	reports := readDayTwoInput()
	t.Logf("Safe reports: %d", CountSafeReports(reports))
}

func TestCountSafeReportsWithCandidatesInput(t *testing.T) {
	reports := readDayTwoInput()
	t.Logf("Safe reports with candidates: %d", CountSafeReportsWithCandidates(reports))
}

func readDayTwoInput() []Report {
	cwd, _ := os.Getwd()
	file, _ := os.Open(cwd + "/02_input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var reports []Report
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		var levels []int
		for _, part := range parts {
			input, _ := strconv.Atoi(part)
			levels = append(levels, input)
		}
		reports = append(reports, Report{levels})
	}

	return reports
}
