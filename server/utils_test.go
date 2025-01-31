package main

import "testing"

func TestGenerateURLID(t *testing.T) {
	type testInput struct {
		Length int
	}

	type testOutput struct {
		ErrorMessage string
	}

	cases := []struct {
		Name   string
		Input  testInput
		Output testOutput
	}{
		{
			Name: "Happy path",
			Input: testInput{
				Length: 5,
			},
			Output: testOutput{
				ErrorMessage: "",
			},
		},
		{
			Name: "Length of 0",
			Input: testInput{
				Length: 0,
			},
			Output: testOutput{
				ErrorMessage: "length must be greater than 0",
			},
		},
		{
			Name: "Length less than 0",
			Input: testInput{
				Length: -5,
			},
			Output: testOutput{
				ErrorMessage: "length must be greater than 0",
			},
		},
	}

	for _, c := range cases {
		result, err := generateURLID(c.Input.Length)
		if err != nil {
			if err.Error() != c.Output.ErrorMessage {
				t.Errorf("Unexpected error: %v\n", err.Error())
			}
		} else {
			if len(result) != c.Input.Length {
				t.Errorf("oh no!\n")
			}
		}
		t.Logf("Test case %v: Success!", c.Name)
	}
}
