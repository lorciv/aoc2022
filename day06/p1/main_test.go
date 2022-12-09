package main

import "testing"

func TestDetect(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5},
		{"nppdvjthqldpwncqszvftbrmjlhg", 6},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11},
	}

	for i, test := range tests {
		if got := detect(test.input); got != test.want {
			t.Errorf("detect(tests[%d]) = %d, want %d", i, got, test.want)
		}
	}
}
