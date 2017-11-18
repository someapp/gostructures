package gostructures

type MaxHeap struct {
	tree []int
}

func NewMaxHeap() MaxHeap {
	return MaxHeap{
		tree: make([]int, 0),
	}
}

func (mh *MaxHeap) Insert(value int) {
	index := len(mh.tree)
	mh.tree = append(mh.tree, value)

	// bubble
	for index > 0 {
		up := ((index + 1) / 2) - 1
		if mh.tree[index] < mh.tree[up] {
			break
		}
		tmp := mh.tree[index]
		mh.tree[index] = mh.tree[up]
		mh.tree[up] = tmp

		// next
		index = up
	}
}

func (mh *MaxHeap) Max() int {
	return mh.tree[0]
}

func (mh *MaxHeap) Pop() {
	// pop off the last element
	end := mh.tree[len(mh.tree)-1]
	mh.tree[0] = end
	mh.tree = mh.tree[:len(mh.tree)-1]

	i := 0
	// bubble
	for i < len(mh.tree) {
		left := ((i + 1) * 2) - 1
		right := left + 1

		if left >= len(mh.tree) {
			break
		}

		if right < len(mh.tree) && mh.tree[right] > mh.tree[left] {
			left = right
		}

		if mh.tree[left] > mh.tree[i] {
			tmp := mh.tree[left]
			mh.tree[left] = mh.tree[i]
			mh.tree[i] = tmp
		}

		i = left
	}
}
