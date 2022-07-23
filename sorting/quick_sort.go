package sorting

import (
	"crypto/rand"
	"math/big"
)

func QuickSort(a []int) {
	if len(a) < 2 {
		return
	}

	left, right := 0, len(a)-1

	random_pivot, _ := rand.Int(rand.Reader, big.NewInt(int64(len(a)))) // select random pivot
	pivot := random_pivot.Int64()

	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	QuickSort(a[:left])
	QuickSort(a[left+1:])
}
