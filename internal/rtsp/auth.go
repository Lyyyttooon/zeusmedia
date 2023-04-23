package rtsp

import (
	"encoding/base64"
	"fmt"

	"github.com/Lyyyttooon/zeusmedia/utils"
)

const (
	AuthTypeDigest = "Digest"
	AuthTypeBasic  = "Basic"
)

type Auth struct {
	Username string
	Password string

	Typ       string
	Realm     string
	Nonce     string
	Algorithm string
	Uri       string
	Response  string
	Opaque    string
	Stale     string
}

func (a *Auth) MakeAuthorization(method, uri string) string {
	if a.Username == "" {
		return ""
	}
	switch a.Typ {
	case AuthTypeBasic:
		base1 := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf(`%s:%s`, a.Username, a.Password)))
		return fmt.Sprintf(`%s %s`, a.Typ, base1)
	case AuthTypeDigest:
		ha1 := utils.Md5([]byte(fmt.Sprintf("%s:%s:%s", a.Username, a.Realm, a.Password)))
		ha2 := utils.Md5([]byte(fmt.Sprintf("%s:%s", method, uri)))
		response := utils.Md5([]byte(fmt.Sprintf("%s:%s:%s", ha1, a.Nonce, ha2)))
		return fmt.Sprintf(`%s username="%s", realm="%s", nonce="%s", uri="%s", response="%s", algorithm="%s"`, a.Typ, a.Username, a.Realm, a.Nonce, uri, response, a.Algorithm)
	}
	return ""
}
