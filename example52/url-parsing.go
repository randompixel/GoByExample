package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {

	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	// Parse URL and check for errors
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	// postgres
	fmt.Println(u.Scheme)

	// user:pass
	fmt.Println(u.User)

	// user
	fmt.Println(u.User.Username())

	// pass
	// second field for "is password set?"
	p, _ := u.User.Password()
	fmt.Println(p)

	// host.com:5432
	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	// host.com
	fmt.Println(host)
	// 5432
	fmt.Println(port)

	// /path
	fmt.Println(u.Path)
	// f
	fmt.Println(u.Fragment)

	// k=v
	fmt.Println(u.RawQuery)

	// Parse query params into map
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0]) // v
}
