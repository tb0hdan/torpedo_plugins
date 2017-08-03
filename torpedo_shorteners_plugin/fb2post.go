package torpedo_shorteners_plugin

import (
    "fmt"
    "log"
    "net/url"
    "strings"

    "golang.org/x/net/html"
)

func shorten_fb(s string) {
    result, _ := url.QueryUnescape(s)
doc, err := html.Parse(strings.NewReader(result))
if err != nil {
    // ...
}
var f func(*html.Node)
f = func(n *html.Node) {
    if n.Type == html.ElementNode && n.Data == "iframe" {
	// Do something with n...
	fmt.Printf("%+v\n", n.Attr[0].Val)
    u, err := url.Parse(n.Attr[0].Val)
    if err != nil {
	log.Fatal(err)
    }
    q := u.Query()
    fmt.Printf("%+v\n", q["href"])

    }
    for c := n.FirstChild; c != nil; c = c.NextSibling {
	f(c)
    }
}
f(doc)
}
