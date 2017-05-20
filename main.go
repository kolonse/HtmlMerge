// HtmlMerge project main.go
package main

import (
	"flag"
)

var JS_FLAG = ".js"
var CSS_FLAG = ".css"

var INVALID_DIR = "%$@|"

//var maxSize = flag.Int("maxsize", 500*1000, "-maxsize=<int> 合并后最大文件限制")
//var js = flag.Bool("js", true, "-js=<bool> 是否进行js压缩")
//var css = flag.Bool("css", true, "-css=<bool> 是否进行css压缩")

var html = flag.String("html", INVALID_DIR, "-html=<string> html path")
var jsdir = flag.String("jsdir", INVALID_DIR, "-jsdir=<string> js file root-path.must be web-root-path")
var cssdir = flag.String("cssdir", INVALID_DIR, "-cssdir=<string> css file root-path.must be web-root-path")

//var html = flag.String("html", "promotionmanager.html", "-htmldir=<string> html文件路径")
//var jsdir = flag.String("jsdir", "test", "-jsdir=<string> js文件根路径,必须是web发布的根路径")
//var cssdir = flag.String("cssdir", "test", "-cssdir=<string> css文件根路径,必须是web发布的根路径")
var replace = flag.Bool("replace", false, "-replace=<bool> true will be replace source file, fase will add 'k.merge' flag")

func main() {
	flag.Parse()
	paramCheck()
	run()
}
