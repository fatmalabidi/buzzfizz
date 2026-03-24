package fizzbuzz

import (
	"strconv"
)

type FizzBuzzService interface {
	Generate(int1, int2, limit int, str1, str2 string) []string
}

type Service struct{}

func (s *Service) Generate(int1, int2, limit int, str1, str2 string) []string {
	result := make([]string, 0, limit)
	for i := 1; i <= limit; i++ {
		if i%(int1*int2) == 0 {
			result = append(result, str1+str2)
		} else if i%int2 == 0 {
			result = append(result, str2)
		} else if i%int1 == 0 {
			result = append(result, str1)
		} else {
			result = append(result, strconv.Itoa(i))
		}
	}

	return result
}
