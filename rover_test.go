package mars_rover_problem

import (
	"fmt"
	"strings"
	"testing"
)

func inverse(prog string) string {
	var inv []string
	for i := len(prog) - 1; i >= 0; i-- {
		switch prog[i] {
		case 'L':
			inv = append(inv, "R")
		case 'R':
			inv = append(inv, "L")
		case 'M':
			inv = append(inv, "RRMRR")
		}
	}
	return strings.Join(inv, "")
}

func TestName(t *testing.T) {
	type testCase struct {
		name      string
		rover     rover
		prog      string
		expected  rover
		expectErr bool
	}

	tests := []testCase{
		{
			name:     "case 1",
			rover:    newRover(1, 2, 5, 5, "N"),
			prog:     "LMLMLMLMM",
			expected: newRover(1, 3, 5, 5, "N"),
		},
		{
			name:     "case 2",
			rover:    newRover(3, 3, 5, 5, "E"),
			prog:     "MMRMMRMRRM",
			expected: newRover(5, 1, 5, 5, "E"),
		},
		{
			name:     "spiral test clockwise",
			rover:    newRover(100, 100, 500, 500, "E"),
			prog:     "MRMMRMMRMMMRMMMR",
			expected: newRover(102, 101, 500, 500, "S"),
		},
		{
			name:     "spiral test anticlockwise",
			rover:    newRover(100, 100, 500, 500, "E"),
			prog:     "MLMMLMMLMMMLMMML",
			expected: newRover(102, 99, 500, 500, "N"),
		},
		{
			name:     "inverse spiral test",
			rover:    newRover(102, 101, 500, 500, "S"),
			prog:     inverse("MRMMRMMRMMMRMMMR"),
			expected: newRover(100, 100, 500, 500, "E"),
		},
		{
			name:     "identity",
			rover:    newRover(100, 100, 500, 500, "S"),
			prog:     inverse("MRMMRMMRMMMRMMMR") + "MRMMRMMRMMMRMMMR",
			expected: newRover(100, 100, 500, 500, "S"),
		},
		{
			name:      "too far east",
			rover:     newRover(3, 3, 5, 5, "E"),
			prog:      "MMMMMMMMMMMMM",
			expected:  newRover(5, 3, 5, 5, "E"),
			expectErr: true,
		},
		{
			name:      "too far south",
			rover:     newRover(3, 3, 5, 5, "S"),
			prog:      "MMMMMMMMMMMMM",
			expected:  newRover(3, 0, 5, 5, "S"),
			expectErr: true,
		},
		{
			name:      "too far north",
			rover:     newRover(3, 3, 5, 5, "N"),
			prog:      "MMMMMMMMMMMMM",
			expected:  newRover(3, 5, 5, 5, "N"),
			expectErr: true,
		},
		{
			name:      "too far west",
			rover:     newRover(3, 3, 5, 5, "W"),
			prog:      "MMMMMMMMMMMMM",
			expected:  newRover(0, 3, 5, 5, "W"),
			expectErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r, err := runProg(test.prog, test.rover)

			fmt.Printf("%+v -> %+v, err = %v\n", test.rover, r, err)

			if (err != nil) != test.expectErr {
				t.Fatalf("expected err: %t, got err: %v", test.expectErr, err)
			}

			if r != test.expected {
				t.Fatalf("expected: %v, got: %v", test.expected, r)
			}
		})
	}
}
