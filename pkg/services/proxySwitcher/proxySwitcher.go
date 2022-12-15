package proxySwitcher

import(
	"net/url"
	"net/http"
	"math/rand"
)

var proxies []*url.URL = []*url.URL{
	&url.URL{Host: "127.0.0.1:8080"},
	&url.URL{Host: "127.0.0.1:8081"},
}

func RandomProxySwitcher(_ *http.Request) (*url.URL, error) {
	return proxies[rand.Intn(len(proxies))], nil
}
