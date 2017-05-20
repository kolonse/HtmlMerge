# HTML js/css merge tool #

**How to get code:**

    go get github.com/kolonse/HtmlMerge

use `HtmlMerge -h` for detail

> -html

>     -html=<string> html file path.

> -jsdir
>     -jsdir=<string> js file root-path.must be web-root-path
> -cssdir
> 
>     -cssdir=<string> css file root-path.must be web-root-path
> -replace
>     -replace=<bool> 'true' will be replace source file, 'fase' will add 'k.merge' flag
# example: #

    HtmlMerge -html=test/test.html -jsdir=test -cssdir=test

out reuslt:

    test.k.merge.html
