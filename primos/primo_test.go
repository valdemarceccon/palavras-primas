package primos_test

import (
	"fmt"
	"palavras-primas/primos"
	"testing"
)

func TestPrimoCalculos(t *testing.T) {
	t.Run("dado varios numeros, testa se é primo", func(t *testing.T) {

		cases := []struct {
			given int
			want  bool
		}{
			{given: 2, want: true},
			{given: 3, want: true},
			{given: 4, want: false},
			{given: 5, want: true},
			{given: 6, want: false},
			{given: 7, want: true},
			{given: 20, want: false},
			{given: 21, want: false},
			{given: 23, want: true},
			{given: 99, want: true},
		}

		for _, c := range cases {
			t.Run(fmt.Sprintf("test if %d is prime, and return %v", c.given, c.want), func(t *testing.T) {
				got := primos.IsPrime(c.given)
				if got != c.want {
					t.Errorf("%d deveria ser %v mas é %v", c.given, c.want, got)
				}
			})
		}
	})
}

type CheckWordTestCases struct {
	given string
	want  bool
}

func (c CheckWordTestCases) String() string {
	return fmt.Sprintf("given %v result should be %v", c.given, c.want)
}

func TestCheckWord(t *testing.T) {
	t.Run("given word assert its value", func(t *testing.T) {

		cases := []CheckWordTestCases{
			{given: "b", want: true},
			{given: "a", want: false},
			{given: "dc", want: true},
			{given: "s", want: true},
		}

		for _, testCase := range cases {
			t.Run(testCase.String(), func(t *testing.T) {
				got, _ := primos.CheckWord(testCase.given)

				if got != testCase.want {
					t.Errorf("should be %v but is %v", testCase.want, got)
				}
			})
		}
	})
}
