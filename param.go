package main

import (
	"errors"
	"os"
)

func paramCheck() {
	if *html == INVALID_DIR {
		panic(errors.New("必须设置 -html 为有效html文件路径"))
	}
	if *jsdir == INVALID_DIR {
		panic(errors.New("必须设置 -jsdir 为有效路径"))
	}
	if *cssdir == INVALID_DIR {
		panic(errors.New("必须设置 -cssdir 为有效路径"))
	}
	if *htmldir == INVALID_DIR {
		panic(errors.New("必须设置 -htmldir 为有效路径"))
	}
	_, err := os.Stat(*html)
	if err != nil {
		panic(err)
	}

	_, err = os.Stat(*jsdir)
	if err != nil {
		panic(err)
	}
	_, err = os.Stat(*cssdir)
	if err != nil {
		panic(err)
	}
}
