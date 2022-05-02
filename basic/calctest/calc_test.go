package calctest

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// "$go test" to run tests in folder
// "$go test ./..." to run all tests of a project
// "go test -c" to compile tests without running it

// Some useful flags
// "$go test -cover" display coverage data analysis
// "$go test -coverprofile profile" generates coverprofile file
// "$go tool cover -html=profile" same functionality as above but displayed in HTML file
// "-covermode set/count/atomic" add this flags to choose specific covermode

// Basic and simple unit tests
func TestCalcSum(t *testing.T) {
	expected := 10
	actual := CalcSum(5, 5)
	if expected != actual {
		t.Errorf("Expected %d do not match actual %d", expected, actual)
	}
}

func TestCalcSubtract(t *testing.T) {
	expected := 5
	actual := CalcSubtract(10, 5)
	if expected != actual {
		t.Errorf("Expected %d do not match actual %d", expected, actual)
	}
}

// Unit tests which uses assertion library "testify"
func TestCalcMultiply(t *testing.T) {
	assert.Equal(t, 25, CalcMultiply(5, 5), "they should be equal")
}

func TestCalcDivide(t *testing.T) {
	assert.Equal(t, 5, CalcDivide(25, 5), "they should be equal")
}

// Table test for CalcSum function with few test cases
func TestTableCalcSum(t *testing.T) {
	type params struct {
		a int
		b int
	}

	type testCase struct {
		name string
		args params
		want int
	}

	tests := []testCase {
		{
			name: "test 10 / 2",
			args: params{a: 10, b: 2},
			want: 5,
		},
		{
			name: "test 100 / 2",
			args: params{a: 100, b: 2},
			want: 50,
		},
		{
			name: "test 22 / 2",
			args: params{a: 22, b: 2},
			want: 11,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcDivide(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("CalcDivide() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Running tests in parallel
func TestTimer1(t *testing.T) {
	t.Parallel()
	time.Sleep(500 * time.Microsecond)
}

func TestTimer2(t *testing.T) {
	t.Parallel()
	time.Sleep(500 * time.Microsecond)
}

func TestTimer3(t *testing.T) {
	t.Parallel()
	time.Sleep(500 * time.Microsecond)
}
