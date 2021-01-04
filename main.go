package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func main() {
	fmt.Println("vim-go")
	reader := bufio.NewReader(os.Stdin)
	re, err := regexp.Compile("<title>(.*)</title>")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		return

	}

	for {
		in, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
		}

		in = strings.Replace(in, "\n", "", -1)
		resp, err := http.Get(fmt.Sprintf("https://www.goodreads.com/search?q=%s&id=", in))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			return
		}
		x := re.FindSubmatch(body)
		out := ""
		if x != nil {
			out = string(x[1])
		}
		fmt.Printf("%s,%s\n", in, out)

	}
}
