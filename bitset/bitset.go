/*
BitSet is a fixed-size sequence of n bits.
Also known as bit array or bit vector.

To store whether something is true or false you use a Bool.
But what if you need to remember whether 10,000 things are true or not?

You could make an array of 10,000 booleans
but you can also use 10,000 bits instead.
That's a lot more compact because 10,000 bits fit in less than 1250 bytes(10,000 / 8).
*/
package bitset

const byteLen = 8

type BitSet []byte

func New(size int) BitSet {
	realSize := 1 + (size-1)/byteLen
	return make([]byte, realSize)
}

// Set true at the index
func (bs BitSet) Set(index int) {
	index, mask := bs.getIndexMask(index)
	bs[index] |= mask
}

// Set false at the index
func (bs BitSet) Unset(index int) {
	index, mask := bs.getIndexMask(index)
	bs[index] &= ^mask
}

// Returns the bool value at the index
func (bs BitSet) Get(index int) bool {
	index, mask := bs.getIndexMask(index)
	return bs[index]&mask != 0
}

func (bs BitSet) getIndexMask(index int) (int, byte) {
	return index / byteLen, 1 << (index % byteLen)
}

func Intersection(a, b BitSet) BitSet {
	if len(a) > len(b) {
		a, b = b, a
	}
	res := make([]byte, len(a))
	for i := range res {
		res[i] = a[i] & b[i]
	}
	return res
}

func Union(a, b BitSet) BitSet {
	if len(a) > len(b) {
		a, b = b, a
	}
	res := make([]byte, len(b))
	for i := range a {
		res[i] = a[i] | b[i]
	}
	for i := len(a); i < len(b); i++ {
		res[i] = b[i]
	}
	return res
}

func Difference(a, b BitSet) BitSet {
	if len(a) > len(b) {
		a, b = b, a
	}
	res := make([]byte, len(b))
	for i := range a {
		res[i] = a[i] ^ b[i]
	}
	for i := len(a); i < len(b); i++ {
		res[i] = b[i]
	}
	return res
}
