package year_2024_test

import (
	"bufio"
	"errors"
	"math"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
	"testing"
)

func TotalDistance(left []int, right []int) (int, error) {
	if len(left) != len(right) {
		return 0, errors.New("slices must be the same length")
	}
	slices.Sort(left)
	slices.Sort(right)
	sum := 0
	for i := 0; i < len(left); i++ {
		diff := float64(left[i] - right[i])
		sum += int(math.Abs(diff))
	}
	return sum, nil
}

func TestTotalDistance(t *testing.T) {
	tests := map[string]struct {
		left        []int
		right       []int
		expected    int
		expectedErr error
	}{
		"test_1": {left: []int{1, 2, 3}, right: []int{1, 2, 3}, expected: 0},
		"test_2": {left: []int{1, 2, 3}, right: []int{3, 2, 1}, expected: 0},
		"test_3": {left: []int{1, 2, 3}, right: []int{1, 2, 4}, expected: 1},
		"test_4": {left: []int{1, 2, 3}, right: []int{1, 2, 5}, expected: 2},
		"test_5": {left: []int{1, 2, 3}, right: []int{1, 2, 6}, expected: 3},
		"test_6": {left: []int{1, 2, 3}, right: []int{1, 2, 7}, expected: 4},
		"test_7": {left: []int{1, 2, 3}, right: []int{1, 2, 8}, expected: 5},
		"test_8": {left: []int{1, 2, 3}, right: []int{1, 2, 9}, expected: 6},
		"advent_example": {
			left:     []int{3, 4, 2, 1, 3, 3},
			right:    []int{4, 3, 5, 3, 9, 3},
			expected: 11,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := TotalDistance(tc.left, tc.right)
			if err != nil {
				if err.Error() != tc.expectedErr.Error() {
					t.Errorf("expected: %v, got: %v", tc.expectedErr, err)
				}
			}
			if actual != tc.expected {
				t.Errorf("expected: %v, got: %v", tc.expected, actual)
			}
		})
	}
}

func TestTotalDistanceInput(t *testing.T) {
	leftList, rightList := readInput(t)
	sum, err := TotalDistance(leftList, rightList)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Total distance: %v", sum)
}

func SimilarityScore(left []int, right []int) (int, error) {
	if len(left) != len(right) {
		return 0, errors.New("slices must be the same length")
	}
	rightMap := make(map[int]int)
	for _, v := range right {
		rightMap[v]++
	}
	score := 0
	for _, v := range left {
		occurrences, ok := rightMap[v]
		if ok {
			score += v * occurrences
		}
	}
	return score, nil
}

func TestSimilarityScore(t *testing.T) {
	tests := map[string]struct {
		left        []int
		right       []int
		expected    int
		expectedErr error
	}{
		"advent_example": {
			left:     []int{3, 4, 2, 1, 3, 3},
			right:    []int{4, 3, 5, 3, 9, 3},
			expected: 31,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := SimilarityScore(tc.left, tc.right)
			if err != nil {
				if err.Error() != tc.expectedErr.Error() {
					t.Errorf("expected: %v, got: %v", tc.expectedErr, err)
				}
			}
			if actual != tc.expected {
				t.Errorf("expected: %v, got: %v", tc.expected, actual)
			}
		})
	}
}

func TestSimilarityScoreInput(t *testing.T) {
	leftList, rightList := readInput(t)
	score, err := SimilarityScore(leftList, rightList)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Similarity score: %v", score)
}

func readInput(t *testing.T) ([]int, []int) {
	cwd, _ := os.Getwd()
	file, err := os.Open(filepath.Join(cwd, "01_input.txt"))
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var leftList []int
	var rightList []int
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		leftInput, _ := strconv.Atoi(parts[0])
		rightInput, _ := strconv.Atoi(parts[1])
		leftList = append(leftList, leftInput)
		rightList = append(rightList, rightInput)
	}
	return leftList, rightList
}
