package link 

import(
	"io"
	"golang.org/x/net/html"
	"fmt"
	"strings"
)


//Link type represents each element in our 
//DS for all anchor tags
type Link struct{
	Href string
	Text string
}




//Finding all the anchor tags
func dfs(n *html.Node) []*html.Node{
	
	if n.Type==html.ElementNode && n.Data=="a"{
		return []*html.Node{n}
	}
	
	var ret []*html.Node

	for c:= n.FirstChild; c!=nil; c=c.NextSibling{
		
		ret=append(ret, dfs(c)...)
	}

	return ret
	
}


//Find all the texts inside the anchor tag

func buildText(n *html.Node) string {

	if n.Type==html.TextNode {
		return n.Data
	}

	if n.Type!=html.ElementNode {
		return ""
	}

	var ret string

	for c:=n.FirstChild; c!=nil; c=c.NextSibling{

		ret+=buildText(c)
	}

	ret=strings.Join(strings.Fields(ret)," ")

	return ret
}


//Make individual links in our DS
func MakeLink(n*html.Node) Link{

	var retlink Link
	for _, attr:= range n.Attr {
		
		if attr.Key=="href"{
			retlink.Href=attr.Val
			break
		}
	}
	
	retlink.Text=buildText(n)

	return retlink
	
}


//Creating the links DS
func LinkNodes(n *html.Node) ([]Link, error){

	nodes:=dfs(n)
	
	var links []Link

	for _,node := range nodes{
		links=append(links, MakeLink(node))
	}
	
	
	return links,nil
}



func Parse(r io.Reader) ([]Link,error){

	rootNode,err := html.Parse(r)

	if err!=nil{
		return nil,err
	}

	fmt.Println(rootNode)

	ret,err:= LinkNodes(rootNode)
	
	return ret,err
}


