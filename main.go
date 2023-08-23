package numeric

import (
	"io"
	"strings"
)

func readnum(n int, sc io.RuneScanner) (int, int, error) {
	count := 0
	for {
		c, _, err := sc.ReadRune()
		if err != nil {
			return n, count, err
		}
		if c < '0' || '9' < c {
			sc.UnreadRune()
			return n, count, nil
		}
		n = n*10 + int(c-'0')
		count++
	}
}

func CompareReader(r1, r2 io.RuneScanner) int {
	for {
		c1, _, err1 := r1.ReadRune()
		c2, _, err2 := r2.ReadRune()

		if err1 != nil {
			if err2 != nil {
				return 0
			}
			return -1
		}
		if err2 != nil {
			return +1
		}

		if '0' <= c1 && c1 <= '9' && '0' <= c2 && c2 <= '9' {
			var cnt1, cnt2 int
			n1 := int(c1 - '0')
			n2 := int(c2 - '0')
			n1, cnt1, err1 = readnum(n1, r1)
			n2, cnt2, err2 = readnum(n2, r2)
			if n1 < n2 {
				return -1
			} else if n1 > n2 {
				return +1
			}
			if cnt1 < cnt2 {
				return +1
			} else if cnt1 > cnt2 {
				return -1
			}
			if err1 != nil {
				if err2 != nil {
					return 0
				}
				return -1
			}
			if err2 != nil {
				return +1
			}
			continue
		}
		if c1 < c2 {
			return -1
		} else if c1 > c2 {
			return +1
		}
	}
}

func Compare(s1, s2 string) int {
	return CompareReader(strings.NewReader(s1), strings.NewReader(s2))
}

func LessThan(s1, s2 string) bool {
	return CompareReader(strings.NewReader(s1), strings.NewReader(s2)) < 0
}
