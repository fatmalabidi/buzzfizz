package fizzbuzz

import (
	"slices"
	"testing"
)

func TestGenerate(t *testing.T) {
	tests := []struct {
		name  string
		int1  int
		int2  int
		limit int
		str1  string
		str2  string
		want  []string
	}{
		{
			name:  "basic fizzbuzz",
			int1:  3,
			int2:  5,
			limit: 20,
			str1:  "fizz",
			str2:  "buzz",
			want:  []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz", "16", "17", "fizz", "19", "buzz"},
		},
		{
			name:  "empty limit",
			int1:  3,
			int2:  5,
			limit: 0,
			str1:  "Fizz",
			str2:  "Buzz",
			want:  []string{},
		},
		{
			name:  "params greater than limit",
			int1:  50,
			int2:  10,
			limit: 5,
			str1:  "Fizz",
			str2:  "Buzz",
			want:  []string{"1", "2", "3", "4", "5"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{}
			got := s.Generate(tt.int1, tt.int2, tt.limit, tt.str1, tt.str2)
			if !slices.Equal(got, tt.want) {
				t.Errorf("Generate() = %v, want %v", got, tt.want)
			}
		})
	}
}
