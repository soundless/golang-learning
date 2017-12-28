package gogo

import "testing"

type Pair struct {
    a string
    b int
}

func Test_6_convert(t *testing.T) {
    tests := []Pair{
        {"ABCDEFGHIJKLMNOPQRSTUVWXYZ", 1},
        {"ABCDEFGHIJKLMNOPQRSTUVWXYZ", 2},
        {"ABCDEFGHIJKLMNOPQRSTUVWX", 5},
        {"A", 4},
    }
    results := []string{
        "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
        "ACEGIKMOQSUWYBDFHJLNPRTVXZ",
        "AIQBHJPRXCGKOSWDFLNTVEMU",
        "A",
    }
    for i := 0; i < len(tests); i++ {
        if result := convert(tests[i].a, tests[i].b); result != results[i] {
            t.Fatalf("case %d failed\nactual: %s, expect: %s\n", i, result, results[i])
        }
    }
}
