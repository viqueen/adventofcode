package year_2024_test

import (
	"bufio"
	"os"
	"strconv"
	"testing"
)

type TokenKind string

const (
	OpenParentheses  TokenKind = "("
	CloseParentheses TokenKind = ")"
	Comma            TokenKind = ","
	Multiply         TokenKind = "mul"
	Integer          TokenKind = "integer"
	DontMultiply     TokenKind = "don't()"
	DoMultiply       TokenKind = "do()"
)

type Token struct {
	kind       TokenKind
	value      string
	intValue   int
	nextOffset int
}

func SumMultiplyInstructions(input string) int {
	sum := 0
	for index := 0; index < len(input); index++ {
		multiplyToken := consumeToken(input, index, Multiply)
		if multiplyToken == nil {
			continue
		}
		openParenthesesToken := consumeToken(input, multiplyToken.nextOffset, OpenParentheses)
		if openParenthesesToken == nil {
			continue
		}
		firstIntegerToken := consumeToken(input, openParenthesesToken.nextOffset, Integer)
		if firstIntegerToken == nil {
			continue
		}
		commaToken := consumeToken(input, firstIntegerToken.nextOffset, Comma)
		if commaToken == nil {
			continue
		}
		secondIntegerToken := consumeToken(input, commaToken.nextOffset, Integer)
		if secondIntegerToken == nil {
			continue
		}
		closeParenthesesToken := consumeToken(input, secondIntegerToken.nextOffset, CloseParentheses)
		if closeParenthesesToken == nil {
			continue
		}
		sum += firstIntegerToken.intValue * secondIntegerToken.intValue
	}
	return sum
}

func SumMultiplyInstructionsWhenEnabled(input string) int {
	sum := 0
	multiplyEnabled := true
	for index := 0; index < len(input); index++ {
		dontMultiplyToken := consumeToken(input, index, DontMultiply)
		if dontMultiplyToken != nil {
			multiplyEnabled = false
			continue
		}

		if !multiplyEnabled {
			doMultiplyToken := consumeToken(input, index, DoMultiply)
			if doMultiplyToken != nil {
				multiplyEnabled = true
				continue
			}
		} else {
			multiplyToken := consumeToken(input, index, Multiply)
			if multiplyToken == nil {
				continue
			}
			openParenthesesToken := consumeToken(input, multiplyToken.nextOffset, OpenParentheses)
			if openParenthesesToken == nil {
				continue
			}
			firstIntegerToken := consumeToken(input, openParenthesesToken.nextOffset, Integer)
			if firstIntegerToken == nil {
				continue
			}
			commaToken := consumeToken(input, firstIntegerToken.nextOffset, Comma)
			if commaToken == nil {
				continue
			}
			secondIntegerToken := consumeToken(input, commaToken.nextOffset, Integer)
			if secondIntegerToken == nil {
				continue
			}
			closeParenthesesToken := consumeToken(input, secondIntegerToken.nextOffset, CloseParentheses)
			if closeParenthesesToken == nil {
				continue
			}
			sum += firstIntegerToken.intValue * secondIntegerToken.intValue
		}
	}

	return sum
}

func consumeToken(input string, offset int, kind TokenKind) *Token {
	switch kind {
	case DontMultiply:
		if offset+7 < len(input) && input[offset:offset+7] == "don't()" {
			return &Token{kind: DontMultiply, value: "don't()", nextOffset: offset + 7}
		}
	case DoMultiply:
		if offset+4 < len(input) && input[offset:offset+4] == "do()" {
			return &Token{kind: DoMultiply, value: "do()", nextOffset: offset + 4}
		}
	case OpenParentheses:
		if input[offset] == '(' {
			return &Token{kind: OpenParentheses, value: "(", nextOffset: offset + 1}
		}
	case CloseParentheses:
		if input[offset] == ')' {
			return &Token{kind: CloseParentheses, value: ")", nextOffset: offset + 1}
		}
	case Comma:
		if input[offset] == ',' {
			return &Token{kind: Comma, value: ",", nextOffset: offset + 1}
		}
	case Multiply:
		if offset+3 < len(input) && input[offset:offset+3] == "mul" {
			return &Token{kind: Multiply, value: "mul", nextOffset: offset + 3}
		}
	case Integer:
		parsed := ""
		for i := offset; i < len(input) && i < offset+3; i++ {
			if input[i] >= '0' && input[i] <= '9' {
				parsed += string(input[i])
			} else {
				break
			}
		}
		if parsed != "" {
			intValue, _ := strconv.Atoi(parsed)
			return &Token{kind: Integer, value: parsed, intValue: intValue, nextOffset: offset + len(parsed)}
		}
	}
	return nil
}

func TestConsumeToken(t *testing.T) {
	tests := map[string]struct {
		input    string
		offset   int
		kind     TokenKind
		expected *Token
	}{
		"multiply": {
			input:    "mul(2,4)",
			offset:   0,
			kind:     Multiply,
			expected: &Token{kind: Multiply, value: "mul", nextOffset: 3},
		},
		"open_parentheses": {
			input:    "mul(2,4)",
			offset:   3,
			kind:     OpenParentheses,
			expected: &Token{kind: OpenParentheses, value: "(", nextOffset: 4},
		},
		"integer": {
			input:    "mul(2,4)",
			offset:   4,
			kind:     Integer,
			expected: &Token{kind: Integer, value: "2", intValue: 2, nextOffset: 5},
		},
		"comma": {
			input:    "mul(2,4)",
			offset:   5,
			kind:     Comma,
			expected: &Token{kind: Comma, value: ",", nextOffset: 6},
		},
		"close_parentheses": {
			input:    "mul(2,4)",
			offset:   7,
			kind:     CloseParentheses,
			expected: &Token{kind: CloseParentheses, value: ")", nextOffset: 8},
		},
		"invalid": {
			input:    "mul(2,4)",
			offset:   8,
			kind:     Multiply,
			expected: nil,
		},
		"invalid_integer": {
			input:    "mul(2,4)",
			offset:   2,
			kind:     Integer,
			expected: nil,
		},
		"don't": {
			input:    "don't()",
			offset:   0,
			kind:     DontMultiply,
			expected: &Token{kind: DontMultiply, value: "don't()", nextOffset: 7},
		},
		"do": {
			input:    "do()",
			offset:   0,
			kind:     DoMultiply,
			expected: &Token{kind: DoMultiply, value: "do()", nextOffset: 4},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actual := consumeToken(tc.input, tc.offset, tc.kind)
			if actual != nil && tc.expected != nil {
				if *actual != *tc.expected {
					t.Errorf("expected: %v, got: %v", tc.expected, actual)
				}
			}
		})
	}
}

func TestSumMultiplyInstructions(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected int
	}{
		"advent_example": {
			input:    "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))",
			expected: 161,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actual := SumMultiplyInstructions(tc.input)
			if actual != tc.expected {
				t.Errorf("expected: %v, got: %v", tc.expected, actual)
			}
		})
	}
}

func TestSumMultiplyInstructionsWhenEnabled(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected int
	}{
		"advent_example": {
			input:    "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
			expected: 48,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actual := SumMultiplyInstructionsWhenEnabled(tc.input)
			if actual != tc.expected {
				t.Errorf("expected: %v, got: %v", tc.expected, actual)
			}
		})
	}
}

func TestSumMultiplyInstructionsInput(t *testing.T) {
	input := readDayThreeInput()
	t.Logf("Day 3 part 1 answer: %d", SumMultiplyInstructions(input))
}

func TestSumMultiplyInstructionsWhenEnabledInput(t *testing.T) {
	input := readDayThreeInput()
	t.Logf("Day 3 part 2 answer: %d", SumMultiplyInstructionsWhenEnabled(input))
}

func readDayThreeInput() string {
	cwd, _ := os.Getwd()
	file, _ := os.Open(cwd + "/03_input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	buffer := ""
	for scanner.Scan() {
		buffer += scanner.Text()
	}
	return buffer
}
