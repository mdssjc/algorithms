package two_pointers

import (
	"slices"
	"testing"
)

func TestPairSumSortedBruteForce(t *testing.T) {
	t.Run("Tests an empty array", func(t *testing.T) {
		input := []int{}
		target := 0
		expected := []int{}

		result := PairSumSortedBruteForce(input, target)

		assertResultEquals(t, expected, result)
	})

	t.Run("Tests an array with just one element", func(t *testing.T) {
		input := []int{1}
		target := 1
		expected := []int{}

		result := PairSumSortedBruteForce(input, target)

		assertResultEquals(t, expected, result)
	})

	t.Run("Tests a two-element array that contains a pair that sums to the target", func(t *testing.T) {
		input := []int{2, 3}
		target := 5
		expected := []int{0, 1}

		result := PairSumSortedBruteForce(input, target)

		assertResultEquals(t, expected, result)
	})

	t.Run("Tests a two-element array that doesn't contain a pair that sums to the target", func(t *testing.T) {
		input := []int{2, 4}
		target := 5
		expected := []int{}

		result := PairSumSortedBruteForce(input, target)

		assertResultEquals(t, expected, result)
	})

	t.Run("Tests an array with duplicate values", func(t *testing.T) {
		input := []int{2, 2, 3}
		target := 5
		expected := []int{0, 2} // or [1, 2]

		result := PairSumSortedBruteForce(input, target)

		assertResultEquals(t, expected, result)
	})

	t.Run("Tests if the algorithm works with a negative number in the target pair", func(t *testing.T) {
		input := []int{-1, 2, 3}
		target := 2
		expected := []int{0, 2}

		result := PairSumSortedBruteForce(input, target)

		assertResultEquals(t, expected, result)
	})

	t.Run("Tests if the algorithm works with both numbers of the target pair being negative", func(t *testing.T) {
		input := []int{-3, -2, -1}
		target := -5
		expected := []int{0, 1}

		result := PairSumSortedBruteForce(input, target)

		assertResultEquals(t, expected, result)
	})
}

func TestPairSumSorted(t *testing.T) {
	t.Run("Tests an empty array", func(t *testing.T) {
		input := []int{}
		target := 0
		expected := []int{}

		result := PairSumSorted(input, target)

		assertResultEquals(t, expected, result)
	})

	t.Run("Tests an array with just one element", func(t *testing.T) {
		input := []int{1}
		target := 1
		expected := []int{}

		result := PairSumSorted(input, target)

		assertResultEquals(t, expected, result)
	})

	t.Run("Tests a two-element array that contains a pair that sums to the target", func(t *testing.T) {
		input := []int{2, 3}
		target := 5
		expected := []int{0, 1}

		result := PairSumSorted(input, target)

		assertResultEquals(t, expected, result)
	})

	t.Run("Tests a two-element array that doesn't contain a pair that sums to the target", func(t *testing.T) {
		input := []int{2, 4}
		target := 5
		expected := []int{}

		result := PairSumSorted(input, target)

		assertResultEquals(t, expected, result)
	})

	t.Run("Tests an array with duplicate values", func(t *testing.T) {
		input := []int{2, 2, 3}
		target := 5
		expected := []int{0, 2} // or [1, 2]

		result := PairSumSorted(input, target)

		assertResultEquals(t, expected, result)
	})

	t.Run("Tests if the algorithm works with a negative number in the target pair", func(t *testing.T) {
		input := []int{-1, 2, 3}
		target := 2
		expected := []int{0, 2}

		result := PairSumSorted(input, target)

		assertResultEquals(t, expected, result)
	})

	t.Run("Tests if the algorithm works with both numbers of the target pair being negative", func(t *testing.T) {
		input := []int{-3, -2, -1}
		target := -5
		expected := []int{0, 1}

		result := PairSumSorted(input, target)

		assertResultEquals(t, expected, result)
	})

}

func BenchmarkPairSumSorted(b *testing.B) {
	b.Run("Benchmark - Brute Force", func(b *testing.B) {
		for b.Loop() {
			PairSumSortedBruteForce([]int{-5, -2, 3, 4, 6}, 7)
		}
	})

	b.Run("Benchmark - Two Pointers", func(b *testing.B) {
		for b.Loop() {
			PairSumSorted([]int{-5, -2, 3, 4, 6}, 7)
		}
	})
}

func assertResultEquals(t testing.TB, got, want []int) {
	t.Helper()
	if !slices.Equal(got, want) {
		t.Errorf("Expected %d, but got %d", got, want)
	}
}
