package url

import (
	nurl "net/url"
	"strings"
)

func New(str string, base *Url) *Url {
	var err error

	res := new(Url)
	res.parent, err = nurl.Parse(str)
	if err != nil {
		// не корректный url
		return nil
	}
	res.parent.Host = strings.ToLower(res.parent.Host)
	if res.parent.Path == "" {
		res.parent.Path = "/"
	}
	res.setBaseUrl(base)
	res.param = res.parent.Query()
	return res
}
