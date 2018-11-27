package ex04

import (
	"math/rand"
	"testing"
	"time"
)

var punctuations = []rune{' ', ',', '.', '、', '。'}

func randomPalindrome(rng *rand.Rand, punctuationNum int) string {
	n := rng.Intn(25)
	runes := make([]rune, n+1)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000))
		runes[i] = r
		runes[n-1-i] = r
	}

	for i := 0; i < punctuationNum; i++ {
		n := rng.Intn(len(punctuations))
		p := rng.Intn(len(runes))
		runes = append(runes[:p+1], runes[p:]...)
		runes[p] = punctuations[n]
	}
	return string(runes)
}

func randomNotPalindrome(rng *rand.Rand) string {
	runes := make([]rune, 25)
	for i := 0; i < 25; i++ {
		r := rune(rng.Intn(0x1000))
		runes[i] = r
	}
	return string(runes)
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{
			"", true,
		},
		{
			"a", true,
		},
		{
			"aa", true,
		},
		{
			"ab", false,
		},
	}

	for _, test := range tests {
		if got := IsPalindrome(test.input); got != test.want {
			t.Errorf("IsPalidrome(%q) = %v", test.input, got)

		}
	}
}

func TestIsPalindromeRandom(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng, 0)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}

func TestIsPalindromeRandom_withPunctuation(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng, 4)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}

func TestIsNotPalindromeRandom(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 1000; i++ {
		p := randomNotPalindrome(rng)
		if IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = true, %r", p, p)
		}
	}
}
