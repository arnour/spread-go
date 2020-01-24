package spread

import (
	"crypto/sha256"
	"hash"
	"math"
	"unsafe"
)

type (
	// Spread is the top-level library instance.
	Spread struct {
		hash hash.Hash
	}
)

const (
	minValue   = float64(math.MinInt32)
	rangeValue = float64(math.MaxInt32) - minValue
	highBound  = float64(1)
	lowBound   = float64(0)
)

// New creates an instance of Spread with default sha256 hash implementation if none provided.
func New(h hash.Hash) (s *Spread) {
	if h == nil {
		h = sha256.New()
	}
	return new(h)
}

// Key returns deterministic float64 number in the interval 0.0 to 1.0 that represents its partition.
func (s *Spread) Key(k string) float64 {
	bytes := s.bytes(k)
	i := s.makeInt(bytes[0], bytes[1], bytes[2], bytes[3])
	f := s.fraction(i)
	return s.bound(f)
}

func new(h hash.Hash) (s *Spread) {
	s = &Spread{
		hash: h,
	}
	return s
}

func (s *Spread) bytes(k string) []int8 {
	s.hash.Write([]byte(k))
	arr := s.hash.Sum(nil)
	signed := *(*[]int8)(unsafe.Pointer(&arr))
	s.hash.Reset()
	return signed
}

func (s *Spread) makeInt(b3 int8, b2 int8, b1 int8, b0 int8) int32 {
	return ((int32(b3) << 24) | ((int32(b2) & 0xff) << 16) | ((int32(b1) & 0xff) << 8) | (int32(b0) & 0xff))
}

func (s *Spread) fraction(i int32) float64 {
	return (float64(i) - minValue) / rangeValue
}

func (s *Spread) bound(fraction float64) float64 {
	return math.Max(math.Min(fraction, highBound), lowBound)
}
