package spread

import (
	"bufio"
	"crypto/sha1"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/montanaflynn/stats"

	"github.com/google/uuid"
)

func TestSpreadDeterministicBehavior(t *testing.T) {
	key := "my-special-key-51f4c00e-7e77-4da8-8fbe-28c676cd179c"
	wanted := str(0.8422691961386)
	s := New(nil)
	for i := 0; i < 100; i++ {
		got := str(s.Key(key))
		if got != wanted {
			t.Errorf("Generated fraction for key '%s' was incorrect, got: %s, want: %s", key, got, wanted)
		}
	}
}

func TestSpreadSHA1DistributionBehavior(t *testing.T) {
	s := New(sha1.New())
	const size = 100000
	data := make([]float64, size)
	for i := 0; i < size; i++ {
		key, _ := uuid.NewRandom()
		data[i] = s.Key(key.String())
	}
	q, _ := stats.Quartile(data)
	q1 := fmt.Sprintf("%.2f", q.Q1)
	if q1 != "0.25" {
		t.Errorf("Generated distribution is not even. Q1 is not around 0.25, got: %s", q1)
	}
	q2 := fmt.Sprintf("%.2f", q.Q2)
	if q2 != "0.50" {
		t.Errorf("Generated distribution is not even. Q2 is not around 0.50, got: %s", q2)
	}
	q3 := fmt.Sprintf("%.2f", q.Q3)
	if q3 != "0.75" {
		t.Errorf("Generated distribution is not even. Q3 is not around 0.75, got: %s", q3)
	}
}

func TestSpreadSHA256DistributionBehavior(t *testing.T) {
	s := New(nil)
	const size = 100000
	data := make([]float64, size)
	for i := 0; i < size; i++ {
		key, _ := uuid.NewRandom()
		data[i] = s.Key(key.String())
	}
	q, _ := stats.Quartile(data)
	q1 := fmt.Sprintf("%.2f", q.Q1)
	if q1 != "0.25" {
		t.Errorf("Generated distribution is not even. Q1 is not around 0.25, got: %s", q1)
	}
	q2 := fmt.Sprintf("%.2f", q.Q2)
	if q2 != "0.50" {
		t.Errorf("Generated distribution is not even. Q2 is not around 0.50, got: %s", q2)
	}
	q3 := fmt.Sprintf("%.2f", q.Q3)
	if q3 != "0.75" {
		t.Errorf("Generated distribution is not even. Q3 is not around 0.75, got: %s", q3)
	}
}

func TestSpreadSHA1JVMCompliance(t *testing.T) {
	s := New(sha1.New())
	f, _ := os.Open("internal/sha1.txt")
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ";")
		key := parts[0]
		wanted := parts[1]
		got := str(s.Key(key))
		if got != wanted {
			t.Errorf("SHA1 Generated fraction for key '%s' was incorrect, got: %s, want: %s", key, got, wanted)
		}
	}
}

func TestSpreadSHA256JVMCompliance(t *testing.T) {
	s := New(nil)
	f, _ := os.Open("internal/sha256.txt")
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ";")
		key := parts[0]
		wanted := parts[1]
		got := str(s.Key(key))
		if got != wanted {
			t.Errorf("SHA256 Generated fraction for key '%s' was incorrect, got: %s, want: %s", key, got, wanted)
		}
	}
}

func BenchmarkSHA1(b *testing.B) {
	s := New(sha1.New())
	key, _ := uuid.NewRandom()
	for i := 0; i < b.N; i++ {
		s.Key(key.String())
	}
}

func BenchmarkSHA256(b *testing.B) {
	s := New(nil)
	key, _ := uuid.NewRandom()
	for i := 0; i < b.N; i++ {
		s.Key(key.String())
	}
}

func str(f float64) string {
	return fmt.Sprintf("%.32f", f)[:14]
}
