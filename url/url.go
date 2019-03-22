package url

/*



 */
import (
	"crypto/md5"
	//"fmt"
	nurl "net/url"
	//	"strings"
)

type Url struct {
	parent     *nurl.URL
	base       *Url
	external   bool    // ссылка за пределы сайта
	self       bool    // если страница совпадает с base
	text       string  // текст ссылки
	level      int     // уровень страницы в структуре сайта
	rank       float64 // ранг url
	hostDomain Domains // домены из Host
	param      nurl.Values
}

// Задать базовую страницу для ссылки. Относительно базовой страницы определяются следующие параметры этой ссылки: внешняя ссылка (IsExternal), ссылка на текущую страницу (IsSelf) и уровень. Так же рассчитываются относительные ссылки ссылки.
func (this *Url) setBaseUrl(base *Url) *Url {
	this.base = base
	this.external = true
	this.self = false
	this.level = 0
	if this.base != nil {
		if this.parent.Host == "" {
			this.parent.Host = base.GetHost()

		}
		this.hostDomain.Fill(this.parent.Host)
		if this.parent.Opaque == "" {
			this.parent.Opaque = this.base.GetOpaque()
		}
		if this.parent.Scheme == "" {
			this.parent.Scheme = this.base.GetScheme()
		}

		this.checkExternalSelf()
		this.SetPath(this.GetPath())
		if this.self {
			this.level = this.base.GetLevel()
		} else {
			this.level = this.base.GetLevel() + 1
		}
	} else {
		this.hostDomain.Fill(this.parent.Host)
		this.parent.Path = ResolvePath("", this.parent.Path)
		this.checkExternalSelf()
	}
	return this
}
func (this *Url) checkExternalSelf() *Url {
	this.self = false
	this.external = true
	if this.base != nil {
		// определим внешнюю ссылку и ссылку на эту же страницу
		arHost2 := this.base.GetHostDomains()
		if this.hostDomain.Equal(&arHost2) {
			this.external = false
		}
		if !this.external {
			// определение ссылки на base страницу
			if this.GetPath() == this.base.GetPath() && this.GetQuery().Encode() == this.base.GetQuery().Encode() {
				this.self = true
			}
		}
	}
	return this
}

//******************************************************************************

func (this *Url) SetPath(v string) *Url {
	if this.base != nil {
		this.parent.Path = ResolvePath(this.base.GetPath(), v)
	} else {
		this.parent.Path = ResolvePath("", v)
	}
	this.checkExternalSelf()
	return this
}

func (this *Url) GetPath() string {
	return this.parent.Path
}

func (this *Url) GetHost() string {
	return this.parent.Host
}

func (this *Url) GetHostDomains() Domains {
	return this.hostDomain
}

func (this *Url) SetScheme(v string) *Url {
	this.parent.Scheme = v
	return this
}

func (this *Url) GetScheme() string {
	return this.parent.Scheme
}

func (this *Url) SetOpaque(v string) *Url {
	this.parent.Opaque = v
	return this
}

func (this *Url) GetOpaque() string {
	return this.parent.Opaque
}

func (this *Url) SetFragment(v string) *Url {
	this.parent.Fragment = v
	return this
}

func (this *Url) GetFragment() string {
	return this.parent.Fragment
}

func (this *Url) SetQuery(v string) *Url {
	this.parent.RawQuery = v
	return this
}

func (this *Url) GetQuery() nurl.Values {
	return this.parent.Query()
}
func (this *Url) SetUser(v *nurl.Userinfo) *Url {
	this.parent.User = v
	return this
}

func (this *Url) GetUser() *nurl.Userinfo {
	return this.parent.User
}

func (this *Url) GetURL() nurl.URL {
	duplicate := *this.parent
	return duplicate
}

func (this *Url) String() string {
	return this.parent.String()
}

// *****************************************************************************

// Внешняя ссылка.
func (this *Url) IsExternal() bool {
	return this.external
}

// Url совпадает с базаовой страницей.
func (this *Url) IsSelf() bool {
	return this.self
}

// Вернуть хэш ссылки.
func (this *Url) Hash() []byte {
	h := md5.Sum([]byte(this.String()))
	return h[:]
}

// Задать текст ссылки.
func (this *Url) SetText(t string) *Url {
	this.text = t
	return this
}

// Текст ссылки.
func (this *Url) GetText() string {
	return this.text
}

// Вернуть уровень ссылки.
func (this *Url) GetLevel() int {
	return this.level
}

// Вернуть вес ссылки.
func (this *Url) GetRank() float64 {
	return this.rank
}

// Установить вес ссылки.
func (this *Url) SetRank(r float64) *Url {
	this.rank = r
	return this
}

// сравнивает 2 url
func (this *Url) Equal(u *Url) bool {
	if u != nil {
		return (this.String() == u.String())
	}
	return false
}

// возвращает домен определенного уровня из Host. уровень 0 - пустая строка всегда. Возврат начинается с "конца", т.е. сначала ru -> wocar -> www
func (this *Url) GetHostDomain(level int) string {
	if level > 0 {
		if level > len(this.hostDomain) {
			return ""
		}
		return this.hostDomain[level]
	}
	return ""
}

// возвращает кол-во доменных уровней прим: www.host.ru = 3
func (this *Url) GetHostDomainLevel() int {
	return len(this.hostDomain)
}
