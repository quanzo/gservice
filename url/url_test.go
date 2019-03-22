package url

import (
	"testing"
)

type testURL struct {
	base string
	ref  string
	res  string
}

type testURLadv struct {
	base         string
	ref          string
	res          string
	res_external bool
	res_self     bool
}

var tURL []testURL = []testURL{
	testURL{"http://www.site.ru", "http://www.site.ru/dir0/../dir1/index.php", "http://www.site.ru/dir1/index.php"},
	testURL{"http://www.site.ru/subdir/", "http://www.site.ru/dir0/../dir1/index.php", "http://www.site.ru/dir1/index.php"},
	testURL{"http://www.site.ru/subdir/dir000/", "../dir1/index.php", "http://www.site.ru/subdir/dir1/index.php"},
}

var tURLA []testURLadv = []testURLadv{
	testURLadv{ //0
		"http://www.site.ru",
		"http://www.site.ru/dir0/../dir1/index.php",
		"http://www.site.ru/dir1/index.php",
		false,
		false},
	testURLadv{ //1
		"http://www.site.ru/subdir/",
		"http://www.site.ru/dir0/../dir1/index.php",
		"http://www.site.ru/dir1/index.php",
		false,
		false},
	testURLadv{ //2
		"http://www.site.ru/subdir/dir000/",
		"../dir1/index.php",
		"http://www.site.ru/subdir/dir1/index.php",
		false,
		false},
	testURLadv{ //3
		"http://www.site.ru/subdir/dir000/",
		"../dir000/",
		"http://www.site.ru/subdir/dir000/",
		false,
		true},
	testURLadv{ //4
		"http://www.site.ru/subdir/dir000/",
		"http://sub.site.ru/dir1/../index.php",
		"http://sub.site.ru/index.php",
		true,
		false},
	testURLadv{ //5
		"",
		"http://sub.site.ru/dir1/../index.php",
		"http://sub.site.ru/index.php",
		true,
		false},
	testURLadv{ //6
		"http://www.site.ru/subdir/dir000/",
		"http://site.ru/subdir/dir000/dir001/dir002/../../",
		"http://site.ru/subdir/dir000/",
		false,
		true},
}

func TestUrlN1(t *testing.T) {
	for i, test := range tURL {
		base := New(test.base, nil)
		ref := New(test.ref, base)
		if ref.String() != test.res {
			t.Error("Error #", i, " ", ref, ref.IsExternal())
		}
		t.Log(i, ref, ref.external, ref.self, ref.base)
	}
}

func TestUrlN2(t *testing.T) {
	for i, test := range tURLA {
		var base *Url
		if test.base == "" {
			base = nil
		} else {
			base = New(test.base, nil)
		}

		ref := New(test.ref, base)
		if ref.String() != test.res {
			t.Error("Error url #", i, " ", ref)
		}
		if ref.external != test.res_external {
			t.Error("Error external #", i, " ", ref.external, "!=", test.res_external)
		}
		if ref.self != test.res_self {
			t.Error("Error self #", i, " ", ref.self, "!=", test.res_self)
		}
	}
}
