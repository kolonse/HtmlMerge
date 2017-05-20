package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func load() *goquery.Document {
	file, err := os.Open(*html)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		panic(err)
	}
	return doc
}

func getFilePath(root string, src string) string {
	if src[0] == '/' {
		path, _ := filepath.Abs(filepath.Join(root, src))
		return path
	}
	path, _ := filepath.Abs(filepath.Join(filepath.Dir(*html), src))
	return path
}

func readFile(root string, src string) string {
	f, err := os.Open(getFilePath(root, src))
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf, _ := ioutil.ReadAll(f)
	return string(buf)
}

func mergeFile(root, flag string, q Queue) string {
	fmt.Println("once process:")
	str := ""
	for i := 0; i < len(q); i++ {
		hl, _ := (q[i].(*goquery.Selection)).Attr(flag)
		fmt.Println(flag, " ->", hl)
		str += readFile(root, hl)
	}
	fmt.Println()
	return str
}

func getNewFile(content, suffix string) string {
	m := md5.New()
	m.Write([]byte(content))
	return hex.EncodeToString(m.Sum(nil)) + suffix
}

func writeFile(root, suffix, content string) string {
	name := getNewFile(content, suffix)
	path := filepath.Join(root, name)
	ioutil.WriteFile(path, []byte(content), 666)
	return "/" + strings.Replace(path, "\\", "/", len(path))
}

func replaceNode(node *goquery.Selection, q Queue, flag, newSrc string) {
	attr := ""
	if flag == "script" {
		attr = "src"
	} else if flag == "link" {
		attr = "href"
	}
	(q[0].(*goquery.Selection)).SetAttr(attr, newSrc)
	for i := 1; i < len(q); i++ {
		(q[i].(*goquery.Selection)).Remove()
	}
}

func mergeScript(node *goquery.Selection, q *Queue) {
	jsQueue := NewQueue()
	jsFlag := -1
	reg := regexp.MustCompile("^(https?://|//)")
	node.Each(func(i int, n *goquery.Selection) {
		if n.Is("script") { // 处理 css link
			src, exist := n.Attr("src")
			if !exist || src[len(src)-len(JS_FLAG):] != JS_FLAG {
				return
			}
			if reg.Match([]byte(src)) {
				return
			}
			if i != jsFlag+1 && !jsQueue.Empty() && jsQueue.Size() != 1 {
				newsrc := writeFile(*jsdir, ".js", mergeFile(*jsdir, "src", jsQueue))
				replaceNode(node, jsQueue, "script", newsrc)
				jsQueue.Clear()
			}

			jsFlag = i
			jsQueue.In(n)
			return
		}
		q.In(n.Children())
	})
	if !jsQueue.Empty() && jsQueue.Size() != 1 {
		newsrc := writeFile(*jsdir, ".js", mergeFile(*jsdir, "src", jsQueue))
		replaceNode(node, jsQueue, "script", newsrc)
	}
}

func mergeCss(node *goquery.Selection, q *Queue) {
	cssSec := node.Find("link")
	reg := regexp.MustCompile("^(https?://|//)")

	cssQueue := NewQueue()
	cssFlag := -1
	cssSec.Each(func(i int, n *goquery.Selection) {
		src, exist := n.Attr("href")
		if !exist || src[len(src)-len(CSS_FLAG):] != CSS_FLAG {
			return
		}
		if reg.Match([]byte(src)) {
			return
		}

		if i != cssFlag+1 && !cssQueue.Empty() && cssQueue.Size() != 1 {
			newsrc := writeFile(*cssdir, ".css", mergeFile(*jsdir, "href", cssQueue))
			replaceNode(node, cssQueue, "link", newsrc)
			cssQueue.Clear()
		}

		cssFlag = i
		cssQueue.In(n)
		return
	})

	if !cssQueue.Empty() && cssQueue.Size() != 1 {
		newsrc := writeFile(*cssdir, ".css", mergeFile(*jsdir, "href", cssQueue))
		replaceNode(node, cssQueue, "link", newsrc)
	}
}

func run() {
	var doc = load()
	q := NewQueue()
	mergeCss(doc.Children(), &q)
	q.In(doc.Children())
	for !q.Empty() {
		node := q.Out().(*goquery.Selection)
		mergeScript(node, &q)
	}
	f := *html
	if !*replace {
		base := filepath.Base(f)
		ext := filepath.Ext(f)
		dir := filepath.Dir(f)
		f = filepath.Join(dir, base[0:len(base)-len(ext)]+".k.merge"+ext)
		//		fmt.Println(base, ext, f)
	}
	h, _ := doc.Html()
	ioutil.WriteFile(f, []byte(h), 666)
}
