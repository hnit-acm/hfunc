package utils

type LessFunc func(i, j int) bool
type SwapFunc func(i, j int)
type LenFunc func() int
type SortFuncs func() (LenFunc, LessFunc, SwapFunc)

// maxDepth returns a threshold at which quicksort should switch
// to heapsort. It returns 2*ceil(lg(n+1)).
func maxDepth(n int) int {
	var depth int
	for i := n; i > 0; i >>= 1 {
		depth++
	}
	return depth * 2
}

func SortFunc(sortFuncs SortFuncs) {
	lenth, _, _ := sortFuncs()
	quickSort(sortFuncs, 0, lenth(), maxDepth(lenth()))
}

func insertionSort(sortFuncs SortFuncs, a, b int) {
	_, less, swap := sortFuncs()
	for i := a + 1; i < b; i++ {
		for j := i; j > a && less(j, j-1); j-- {
			swap(j, j-1)
		}
	}
}

func quickSort(sortFuncs SortFuncs, a, b, maxDepth int) {
	_, less, swap := sortFuncs()
	for b-a > 12 { // Use ShellSort for slices <= 12 elements
		if maxDepth == 0 {
			heapSort(sortFuncs, a, b)
			return
		}
		maxDepth--
		mlo, mhi := doPivot(sortFuncs, a, b)
		// Avoiding recursion on the larger subproblem guarantees
		// a stack depth of at most lg(b-a).
		if mlo-a < b-mhi {
			quickSort(sortFuncs, a, mlo, maxDepth)
			a = mhi // i.e., quickSort(data, mhi, b)
		} else {
			quickSort(sortFuncs, mhi, b, maxDepth)
			b = mlo // i.e., quickSort(data, a, mlo)
		}
	}
	if b-a > 1 {
		// Do ShellSort pass with gap 6
		// It could be written in this simplified form cause b-a <= 12
		for i := a + 6; i < b; i++ {
			if less(i, i-6) {
				swap(i, i-6)
			}
		}
		insertionSort(sortFuncs, a, b)
	}
}

// first is an offset into the array where the root of the heap lies.
func siftDown(sortFuncs SortFuncs, lo, hi, first int) {
	root := lo
	_, less, swap := sortFuncs()
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		if child+1 < hi && less(first+child, first+child+1) {
			child++
		}
		if !less(first+root, first+child) {
			return
		}
		swap(first+root, first+child)
		root = child
	}
}

func heapSort(sortFuncs SortFuncs, a, b int) {
	first := a
	lo := 0
	hi := b - a
	_, _, swap := sortFuncs()

	// Build heap with greatest element at top.
	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDown(sortFuncs, i, hi, first)
	}

	// Pop elements, largest first, into end of data.
	for i := hi - 1; i >= 0; i-- {
		swap(first, first+i)
		siftDown(sortFuncs, lo, i, first)
	}
}

func medianOfThree(sortFuncs SortFuncs, m1, m0, m2 int) {
	// sort 3 elements
	_, less, swap := sortFuncs()

	if less(m1, m0) {
		swap(m1, m0)
	}
	// data[m0] <= data[m1]
	if less(m2, m1) {
		swap(m2, m1)
		// data[m0] <= data[m2] && data[m1] < data[m2]
		if less(m1, m0) {
			swap(m1, m0)
		}
	}
	// now data[m0] <= data[m1] <= data[m2]
}

func doPivot(sortFuncs SortFuncs, lo, hi int) (midlo, midhi int) {
	_, less, swap := sortFuncs()
	m := int(uint(lo+hi) >> 1) // Written like this to avoid integer overflow.
	if hi-lo > 40 {
		// Tukey's ``Ninther,'' median of three medians of three.
		s := (hi - lo) / 8
		medianOfThree(sortFuncs, lo, lo+s, lo+2*s)
		medianOfThree(sortFuncs, m, m-s, m+s)
		medianOfThree(sortFuncs, hi-1, hi-1-s, hi-1-2*s)
	}
	medianOfThree(sortFuncs, lo, m, hi-1)

	// Invariants are:
	//	data[lo] = pivot (set up by ChoosePivot)
	//	data[lo < i < a] < pivot
	//	data[a <= i < b] <= pivot
	//	data[b <= i < c] unexamined
	//	data[c <= i < hi-1] > pivot
	//	data[hi-1] >= pivot
	pivot := lo
	a, c := lo+1, hi-1

	for ; a < c && less(a, pivot); a++ {
	}
	b := a
	for {
		for ; b < c && !less(pivot, b); b++ { // data[b] <= pivot
		}
		for ; b < c && less(pivot, c-1); c-- { // data[c-1] > pivot
		}
		if b >= c {
			break
		}
		// data[b] > pivot; data[c-1] <= pivot
		swap(b, c-1)
		b++
		c--
	}
	// If hi-c<3 then there are duplicates (by property of median of nine).
	// Let's be a bit more conservative, and set border to 5.
	protect := hi-c < 5
	if !protect && hi-c < (hi-lo)/4 {
		// Lets test some points for equality to pivot
		dups := 0
		if !less(pivot, hi-1) { // data[hi-1] = pivot
			swap(c, hi-1)
			c++
			dups++
		}
		if !less(b-1, pivot) { // data[b-1] = pivot
			b--
			dups++
		}
		// m-lo = (hi-lo)/2 > 6
		// b-lo > (hi-lo)*3/4-1 > 8
		// ==> m < b ==> data[m] <= pivot
		if !less(m, pivot) { // data[m] = pivot
			swap(m, b-1)
			b--
			dups++
		}
		// if at least 2 points are equal to pivot, assume skewed distribution
		protect = dups > 1
	}
	if protect {
		// Protect against a lot of duplicates
		// Add invariant:
		//	data[a <= i < b] unexamined
		//	data[b <= i < c] = pivot
		for {
			for ; a < b && !less(b-1, pivot); b-- { // data[b] == pivot
			}
			for ; a < b && less(a, pivot); a++ { // data[a] < pivot
			}
			if a >= b {
				break
			}
			// data[a] == pivot; data[b-1] < pivot
			swap(a, b-1)
			a++
			b--
		}
	}
	// Swap pivot into middle
	swap(pivot, b-1)
	return b - 1, c
}
