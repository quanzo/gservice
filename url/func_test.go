package url

import (
	"testing"
)

type testResolve struct {
	base   string
	target string
	result string
}

var tResolve []testResolve = []testResolve{
	testResolve{"/testq/testw/teste/testr/", "../bbb1/index.php", "/testq/testw/teste/bbb1/index.php"},
	testResolve{"/testq/testw/teste/testr/", "bbb1/index.php", "/testq/testw/teste/testr/bbb1/index.php"},
	testResolve{"/testq/testw/teste/testr/", ".././../bbb1/index.php", "/testq/testw/bbb1/index.php"},
	testResolve{"/testq/testw/teste/testr/", "./bbb1/index.php", "/testq/testw/teste/testr/bbb1/index.php"},
}

func TestResolve(t *testing.T) {
	for i, test := range tResolve {
		if ResolvePath(test.base, test.target) != test.result {
			t.Error("Error #", i, " ", ResolvePath(test.base, test.target))
		}
	}
}
