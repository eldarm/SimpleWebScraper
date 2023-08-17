package scrape

import (
	//"log"
	"testing"
)

func TestEmptyQueuePush(t *testing.T) {
	tests := []LUrl{
		LUrl{
			url:      "https://coo.com/abc/def/ghi.html?q=aaa&b=bbb",
			protocol: "https",
			host:     "coo.com",
			path:     "/abc/def",
			name:     "ghi.html",
			args:     "q=aaa&b=bbb",
			err:      nil,
		},
		LUrl{
			url:      "http://coo.com/abc/def/ghi.html#foo?qq=bcd",
			protocol: "http",
			host:     "coo.com",
			path:     "/abc/def",
			name:     "ghi.html",
			args:     "qq=bcd",
			err:      nil,
		},
		LUrl{
			url:      "/abc/def/ghi.html#foo?qq=bcd",
			protocol: "",
			host:     "",
			path:     "/abc/def",
			name:     "ghi.html",
			args:     "qq=bcd",
			err:      nil,
		},
		LUrl{
			url:      "abc/def/ghi.html#foo?qq=bcd",
			protocol: "",
			host:     "",
			path:     "abc/def",
			name:     "ghi.html",
			args:     "qq=bcd",
			err:      nil,
		},
		LUrl{
			url:      "https://management4cannibals.blogspot.com",
			protocol: "https",
			host:     "management4cannibals.blogspot.com",
			path:     "",
			name:     "",
			args:     "",
			err:      nil,
		},
		LUrl{
			url:      "https://management4cannibals.blogspot.com/",
			protocol: "https",
			host:     "management4cannibals.blogspot.com",
			path:     "",
			name:     "",
			args:     "",
			err:      nil,
		},
		LUrl{
			url:      "//blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEgLcLAk7mDanogb2rGgJ-6QgDeRUHJ3hjFBLFynCpD_KrbdYo2Wk6a7hsTSHAIu7mQ2sctE3LKbVx_bC8p-dBKr-ynMqGhgykeqOYBrnjKugAbC32PF6grqN2JaPHKmfQ/s220/150px-Scrat.jpg",
			protocol: "http",
			host:     "blogger.googleusercontent.com",
			path:     "/img/b/R29vZ2xl/AVvXsEgLcLAk7mDanogb2rGgJ-6QgDeRUHJ3hjFBLFynCpD_KrbdYo2Wk6a7hsTSHAIu7mQ2sctE3LKbVx_bC8p-dBKr-ynMqGhgykeqOYBrnjKugAbC32PF6grqN2JaPHKmfQ/s220",
			name:     "150px-Scrat.jpg",
			args:     "",
			err:      nil,
		},
	}
	for _, test := range tests {
		r := ParseUrl(test.url)
		if r.String() != test.String() {
			t.Errorf("Parse error:\ne: %v\nr: %v", test, *r)
		}
	}
}
