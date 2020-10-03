package helpers

import "net/url"

func AddQ(url *url.URL, key string, value string) {
	q := url.Query()
	q.Set(key, value)
	url.RawQuery = q.Encode()
}