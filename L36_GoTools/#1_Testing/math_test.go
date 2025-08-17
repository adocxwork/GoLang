package math

import "testing"

func Add(a, b int) int {
    return a + b
}

// Run the test
// go test <file name>
func TestAdd(t *testing.T) {
    result := Add(2, 3)
    if result != 5 {
        t.Errorf("Add(2,3) failed. Expected %d, got %d", 5, result)
    }
}

func TestAddTable(t *testing.T) { // best practice
    tests := []struct {
        a, b, expected int
    }{
        {2, 3, 5},
        {10, 20, 30},
        {-1, 1, 0},
    }

    for _, tt := range tests {
        result := Add(tt.a, tt.b)
        if result != tt.expected {
            t.Errorf("Add(%d,%d): expected %d, got %d", tt.a, tt.b, tt.expected, result)
        }
    }
}



// Go supports performance benchmarking in tests!
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(2, 3)
    }
}
// run benchmark
// go test -bench=.