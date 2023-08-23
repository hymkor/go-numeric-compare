package numeric

import (
	"testing"
)

var ope = map[int]string{
	-1: "<",
	0:  "--",
	+1: ">",
}

func expect(s1, s2 string, result int, t *testing.T) bool {
	diff := Compare(s1, s2)
	if diff != result {
		t.Fatalf("failed: %s %s %s", s1, ope[result], s2)
		return false
	}
	return true
}

func TestCompare(t *testing.T) {
	_ = (expect("a100", "a2", +1, t) &&
		expect("a2", "a100", -1, t) &&
		expect("a2", "a3", -1, t) &&
		expect("a3", "a2", +1, t) &&
		expect("a2", "b1", -1, t) &&
		expect("a", "b", -1, t) &&
		expect("a001", "a1", -1, t) &&
		expect("a1", "a0001", +1, t) &&
		expect("v80", "v100", -1, t) &&
		expect("v100", "v80", +1, t) &&
		expect("v100", "v120", -1, t) &&
		expect("v120", "v100", +1, t) &&
		expect("v100", "v100", 0, t) &&
		expect("0200", "v100", -1, t))
}

func BenchmarkCompare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Compare("a100", "a2")
		Compare("a2", "a100")
		Compare("a2", "a3")
		Compare("a3", "a2")
		Compare("a2", "b1")
		Compare("a", "b")
		Compare("a001", "a1")
		Compare("a1", "a0001")
		Compare("v80", "v100")
		Compare("v100", "v80")
		Compare("v100", "v120")
		Compare("v120", "v100")
		Compare("v100", "v100")
		Compare("0200", "v100")
	}
}
