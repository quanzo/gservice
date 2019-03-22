package url

import (
	"strings"
	"unicode/utf8"
)

type Domains []string

func (this Domains) String() string {
	var (
		size int = 0
		i    int
	)
	count := len(this)
	for i = 0; i < count; i++ {
		size += len(this[i])
	}
	if count > 2 {
		size += count - 2
	}

	if size > 0 {
		b := make([]byte, size)
		size = 0
		for i = count - 1; i > 0; i-- {
			copy(b[size:], this[i])
			size += len(this[i])
			if i > 1 {
				copy(b[size:], ".")
				size++
			}

		}
		return string(b)
	} else {
		return ""
	}
} // end func

func (this *Domains) Fill(host string) {
	size := utf8.RuneCountInString(host)
	if len(*this) > 0 {
		(*this)[0] = ""
	} else {
		(*this) = append((*this), "")
	}
	counterHostDomain := 1
	punktePos := 1
	endIndex := size
	domain := ""

	for punktePos != -1 {
		punktePos = strings.LastIndex(host[0:endIndex], ".")
		domain = strings.ToLower(host[punktePos+1 : endIndex])
		if len(*this) > counterHostDomain {
			(*this)[counterHostDomain] = domain
		} else {
			(*this) = append((*this), domain)
		}
		endIndex = punktePos
		counterHostDomain++
	}
	// корректировка
	if len(*this) > counterHostDomain {
		(*this) = (*this)[0:counterHostDomain]
	}
}

func (this *Domains) Equal(d *Domains) bool {
	s1 := len(*this)
	s2 := len(*d)
	ds := s1 - s2
	if ds < 0 {
		ds *= -1
	}
	if ds > 1 {
		return false
	}
	if ds == 1 {
		if s1 > s2 && (*this)[s1-1] == "www" {
			s1--
		}
		if s1 < s2 && (*d)[s2-1] == "www" {
			s2--
		}
	}
	if s1 == s2 {
		for i := 0; i < s1; i++ {
			if (*this)[i] != (*d)[i] {
				return false
			}
		}
		return true
	}
	return false
}
