package main

import "fmt"

func main() {
	var html string
	html = "<html>\n\t<body>\"hello\"</body>\n</html>"
	fmt.Println(html)

	var rawHtml string
	rawHtml = `
<html>
	<body>"hello"</body>
</html>`
	fmt.Println(rawHtml)
}
