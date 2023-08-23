package numeric

import (
	"unicode/utf8"
)

func readnum(n int, s string) (int, int, string) {
	count := 0
	for len(s) > 0 {
		c, siz := utf8.DecodeRuneInString(s)
		if c < '0' || '9' < c {
			return n, count, s
		}
		n = n*10 + int(c-'0')
		s = s[siz:]
		count++
	}
	return n, count, s
}

func Compare(s1, s2 string) int {
	for {
		if len(s1) <= 0 {
			if len(s1) <= 0 {
				return 0
			}
			return -1
		}
		if len(s2) <= 0 {
			return +1
		}

		c1, siz := utf8.DecodeRuneInString(s1)
		s1 = s1[siz:]

		c2, siz := utf8.DecodeRuneInString(s2)
		s2 = s2[siz:]

		if '0' <= c1 && c1 <= '9' && '0' <= c2 && c2 <= '9' {
			var cnt1, cnt2 int
			n1 := int(c1 - '0')
			n2 := int(c2 - '0')
			n1, cnt1, s1 = readnum(n1, s1)
			n2, cnt2, s2 = readnum(n2, s2)
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
		} else if c1 < c2 {
			return -1
		} else if c1 > c2 {
			return +1
		}
	}
}

func LessThan(s1, s2 string) bool {
	return Compare(s1, s2) < 0
}
