# HTML js/css merge tool #

**How to get code:**

    go get github.com/kolonse/HtmlMerge

use `HtmlMerge -h` for detail

> -html
> 
>     -html=<string> html file path.
> -htmldir
>
>     -htmldir=<string> html file root-path. must be web-root-path
> -jsdir
> 
>     -jsdir=<string> js file root-path.must be web-root-path
> -cssdir
> 
>     -cssdir=<string> css file root-path.must be web-root-path
> -replace
> 
>     -replace=<bool> 'true' will be replace source file, 'fase' will add 'k.merge' flag
> -outfile
> 
>     -outfile=<bool> true will out source file to stdout;false not
# example: #

    HtmlMerge -html=test/test.html -jsdir=test -cssdir=test -htmldir=test

out reuslt:

    test.k.merge.html
