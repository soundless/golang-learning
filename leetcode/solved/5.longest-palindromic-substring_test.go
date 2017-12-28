package gogo

import "testing"

func Test_5_longestPalindrome(t *testing.T) {
	tests := []string{
		"a",
		"aaaaa",
		"ababababab",
		"abcfcdfgfdc",
		"babad",
		"cbbd",
	}
	results := []string{
		"a",
		"aaaaa",
		"ababababa",
		"cdfgfdc",
		"bab",
		"bb",
	}
	for i := 0; i < len(tests); i++ {
		if result := longestPalindrome(tests[i]); result != results[i] {
			t.Fatalf("case %d failed\nactual: %s, expect: %s\n", i, result, results[i])
		}
	}
}

func Benchmark_longestPalindrome(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			tests := []string{
				"a",
				"aaaaa",
				"ababababab",
				"abcfcdfgfdc",
				"babad",
				"cbbd",
			}
			for i := 0; i < len(tests); i++ {
				longestPalindrome(tests[i])
			}
		}
	})
}
