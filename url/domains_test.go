package url

import (
	"testing"
)

type testCompose struct {
	domains Domains
	res     string
}

var tCompose []testCompose = []testCompose{
	testCompose{Domains{"", "ru", "site", "sub", "www"}, "www.sub.site.ru"},
	testCompose{Domains{"", "ru", "site", "www"}, "www.site.ru"},
	testCompose{Domains{"", "ru", "www"}, "www.ru"},
	testCompose{Domains{"", "ru"}, "ru"},
}

type testEqual struct {
	s1  string
	s2  string
	res bool
}

var tEqual []testEqual = []testEqual{
	testEqual{"www.site.ru", "site.ru", true},
	testEqual{"site.ru", "www.site.ru", true},
	testEqual{"site.ru", "site.ru", true},
	testEqual{"www1.site.ru", "www.site.ru", false},
	testEqual{"www.site.ru", "www1.site.ru", false},
	testEqual{"www.site.ru", "www.site.ru", true},
	testEqual{"site.ru", "siter.ru", false},
	testEqual{"www.sitey.ru", "site.ru", false},
	testEqual{"www.site.ru", "sitey.ru", false},
	testEqual{"www.site.net", "www.site.ru", false},
	testEqual{"www.sitey.ru", "www.site.ru", false},
	testEqual{"www.site.ru", "site.net", false},
	testEqual{"www.sub.site.ru", "sub.site.ru", true},
}

var tFill []string = []string{
	"www.site.ru",
	"site.ru",
	"www.sub.sub.site.ru",
	"www.site.ru",
	"ru",
}

func TestCompose(t *testing.T) {
	for i, test := range tCompose {
		if test.domains.String() != test.res {
			t.Error("Error test #", i, ". "+test.domains.String())
		}
	}
}

func TestFill(t *testing.T) {
	d := Domains{}
	for i, test := range tFill {
		d.Fill(test)
		if d.String() != test {
			t.Error("Error test #", i, " ", d.String())
		}
		//t.Log(d)
	}
}
func TestFill2(t *testing.T) {
	d := Domains{}
	d.Fill("www.site.ru:8080")
	t.Log(d.String())

}

func TestEqual(t *testing.T) {
	for i, test := range tEqual {
		d1 := Domains{}
		d1.Fill(test.s1)
		d2 := Domains{}
		d2.Fill(test.s2)
		if d1.Equal(&d2) != test.res {
			t.Error("Error #", i)
		}
	}
}
