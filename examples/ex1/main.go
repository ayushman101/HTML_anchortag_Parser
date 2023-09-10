package main

import (
	"fmt"
	"link"
	"strings"
)

func main(){

str := `<!DOCTYPE html>
<html>
<head>
    <title>Anchor Tag Example</title>
</head>
<body>
    <h1>Welcome to my website</h1>
    <p>This is a simple example of an anchor tag.</p>
    
    <!-- Anchor tag with href attribute -->
    <a href="https://www.example.com">Visit Example.com</a>
    
    <p>Click the link above to visit Example.com.</p>
</body>
</html>	
`

temp := strings.NewReader(str)

fmt.Println(link.Parse(temp))

}
