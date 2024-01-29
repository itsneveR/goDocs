// First, you declare the name of the package that you want to test here we wanna test math
package math

import "testing"

///////////////////////// regular test ////////////////
/*
	   func TestAdd(t *testing.T) {

	//here you can call methods and functions from your package and see if the result they give you
	//matches your expectaions.

	get := Add(5, 6)
	want := 11

	if get != want {
		t.Errorf("got %q - want %q", get, want)
	}
	}*/

//////////table driven test////////////////////////
/*
	Here you are defining a struct, populating a table of structs that include the arguments and expected results
	for your Add function, and then writing a new TestAdd function.
	In this new function, you iterate over the table, run the arguments, compare the outputs to each expected result
	and then returning any errors, should they occur.
*/
type testStruct struct {
	arg1, arg2, expected int
}

var testCases = []testStruct{
	{2, 3, 5},
	{5, 5, 10},
	{74, 26, 100},
}

func TestAdd(t *testing.T) {
	for _, test := range testCases {
		if output := Add(test.arg1, test.arg2); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}

////////////////benchmarking//////////

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(5, 6)
	}
}

/*
	N is an integer that can be adjusted. During benchmark execution, b.N is adjusted until the benchmark function
	lasts long enough to be timed reliably. The --bench flag accepts its arguments in the form of a regular expression.
*/
