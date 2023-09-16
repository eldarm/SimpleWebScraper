package scrape

import (
	//"log"
	"testing"
)

func TestParsing(t *testing.T) {
	type test struct {
		eu string
		u  LUrl
	}
	tests := []test{
		// #0
		test{
			eu: "https://coo.com/abc/def/ghi.html",
			u: LUrl{
				url:      "https://coo.com/abc/def/ghi.html?q=aaa&b=bbb",
				protocol: "https",
				host:     "coo.com",
				path:     "/abc/def",
				name:     "ghi.html",
				args:     "",
				err:      nil,
			},
		},
		// #1
		test{
			eu: "http://coo.com/abc/def/ghi.html",
			u: LUrl{
				url:      "http://coo.com/abc/def/ghi.html#foo?qq=bcd",
				protocol: "http",
				host:     "coo.com",
				path:     "/abc/def",
				name:     "ghi.html",
				args:     "",
				err:      nil,
			},
		},
		// #2
		test{
			eu: "/abc/def/ghi.html",
			u: LUrl{
				url:      "/abc/def/ghi.html#foo?qq=bcd",
				protocol: "",
				host:     "",
				path:     "/abc/def",
				name:     "ghi.html",
				args:     "",
				err:      nil,
			},
		},
		// #3
		test{
			eu: "abc/def/ghi.html",
			u: LUrl{
				url:      "abc/def/ghi.html#foo?qq=bcd",
				protocol: "",
				host:     "",
				path:     "abc/def",
				name:     "ghi.html",
				args:     "",
				err:      nil,
			},
		},
		// #4
		test{
			eu: "https://management4cannibals.blogspot.com",
			u: LUrl{
				url:      "https://management4cannibals.blogspot.com",
				protocol: "https",
				host:     "management4cannibals.blogspot.com",
				path:     "",
				name:     "",
				args:     "",
				err:      nil,
			},
		},
		// #5
		test{
			eu: "https://management4cannibals.blogspot.com?page=5", // "/" after host?
			u: LUrl{
				url:      "https://management4cannibals.blogspot.com/?page=5&f=7",
				protocol: "https",
				host:     "management4cannibals.blogspot.com",
				path:     "",
				name:     "",
				args:     "?page=5",
				err:      nil,
			},
		},
		// #6
		test{
			eu: "http://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEgLcLAk7mDanogb2rGgJ-6QgDeRUHJ3hjFBLFynCpD_KrbdYo2Wk6a7hsTSHAIu7mQ2sctE3LKbVx_bC8p-dBKr-ynMqGhgykeqOYBrnjKugAbC32PF6grqN2JaPHKmfQ/s220/150px-Scrat.jpg",
			u: LUrl{
				url:      "//blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEgLcLAk7mDanogb2rGgJ-6QgDeRUHJ3hjFBLFynCpD_KrbdYo2Wk6a7hsTSHAIu7mQ2sctE3LKbVx_bC8p-dBKr-ynMqGhgykeqOYBrnjKugAbC32PF6grqN2JaPHKmfQ/s220/150px-Scrat.jpg",
				protocol: "http",
				host:     "blogger.googleusercontent.com",
				path:     "/img/b/R29vZ2xl/AVvXsEgLcLAk7mDanogb2rGgJ-6QgDeRUHJ3hjFBLFynCpD_KrbdYo2Wk6a7hsTSHAIu7mQ2sctE3LKbVx_bC8p-dBKr-ynMqGhgykeqOYBrnjKugAbC32PF6grqN2JaPHKmfQ/s220",
				name:     "150px-Scrat.jpg",
				args:     "",
				err:      nil,
			},
		},
		// #7
		test{
			eu: "http://eldar.com/user/login",
			u: LUrl{
				url:      "http://eldar.com/index.php/user/login?destination=/index.php/node/92%23comment-form",
				protocol: "http",
				host:     "eldar.com",
				path:     "/user",
				name:     "login",
				args:     "",
				err:      nil,
			},
		},
		// #8
		test{
			eu: "/user/bob/cv.htm",
			u: LUrl{
				url:      "/./user/bob/cv.htm",
				protocol: "",
				host:     "",
				path:     "/user/bob",
				name:     "cv.htm",
				args:     "",
				err:      nil,
			},
		},
		// #9
		test{
			eu: "/user/bob/cv.htm",
			u: LUrl{
				url:      "/../user/bob/cv.htm",
				protocol: "",
				host:     "",
				path:     "/user/bob",
				name:     "cv.htm",
				args:     "",
				err:      nil,
			},
		},
		// #10
		test{
			eu: "/user/bob/cv.htm",
			u: LUrl{
				url:      "/user/./bob/cv.htm",
				protocol: "",
				host:     "",
				path:     "/user/bob",
				name:     "cv.htm",
				args:     "",
				err:      nil,
			},
		},
		// #11
		test{
			eu: "/bob/cv.htm",
			u: LUrl{
				url:      "/user/../bob/cv.htm",
				protocol: "",
				host:     "",
				path:     "/bob",
				name:     "cv.htm",
				args:     "",
				err:      nil,
			},
		},
		// #12
		test{
			eu: "bob/cv.htm",
			u: LUrl{
				url:      "./bob/cv.htm",
				protocol: "",
				host:     "",
				path:     "bob",
				name:     "cv.htm",
				args:     "",
				err:      nil,
			},
		},
		// #13
		test{
			eu: "bob/cv.htm",
			u: LUrl{
				url:      "../bob/cv.htm",
				protocol: "",
				host:     "",
				path:     "bob",
				name:     "cv.htm",
				args:     "",
				err:      nil,
			},
		},
		// #14
		test{
			eu: "bob/cv.htm",
			u: LUrl{
				url:      "user/../bob/cv.htm",
				protocol: "",
				host:     "",
				path:     "bob",
				name:     "cv.htm",
				args:     "",
				err:      nil,
			},
		},
		// #15
		test{
			eu: "user/bob/cv.htm",
			u: LUrl{
				url:      "user/bob/cv.htm",
				protocol: "",
				host:     "",
				path:     "user/bob",
				name:     "cv.htm",
				args:     "",
				err:      nil,
			},
		},


	}
	for i, test := range tests {
		r := ParseUrl(test.u.url)
		u := r.Url()
		if u != test.eu {
			t.Errorf("Url() error in test %d Expected URL to read is %q, actual %q", i, test.eu, u)
		}
		if r.String() != test.u.String() {
			t.Errorf("Parse error in test %d:\ne: %v\nr: %v", i, test.u.String(), r.String())
		}
	}
}
