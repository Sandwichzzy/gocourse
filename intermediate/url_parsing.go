package intermediate

import (
	"fmt"
	"net/url"
)


func main() {

	// [scheme://][userinfo@]host[:port][/path][?query][#fragment]

	rawURL := "https://example.com:8080/path?query=param#fragment"
	//1. parse URL
	parsedURL,err:=url.Parse(rawURL)
	if err!=nil {
		fmt.Println("Error parsing URL",err)
		return
	}

	fmt.Println("Scheme:", parsedURL.Scheme) //https
	fmt.Println("Host:", parsedURL.Host)    //example.com:8080
	fmt.Println("Port:", parsedURL.Port()) //8080
	fmt.Println("Path:", parsedURL.Path) // /path
	fmt.Println("Raw Query:", parsedURL.RawQuery) //query=param
	fmt.Println("Fragment:", parsedURL.Fragment) //fragment
	// fmt.Println("Scheme:", parsedURL.Scheme)

	//2. get queryParams
	rawURL1 :="https://example.com/path?name=john&age=30"
	parsedURL1,err:=url.Parse(rawURL1)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}
	queryParams :=parsedURL1.Query()

	fmt.Println("queryParams",queryParams) // map[age:[30] name:[john]]
	fmt.Println("Name:",queryParams.Get("name")) //Name: john
	fmt.Println("Age:",queryParams.Get("age")) //Age: 30

	//3. Building URL
	// type URL struct {
	//     Scheme      string
	//     Opaque      string    // encoded opaque data
	//     User        *Userinfo // username and password information
	//     Host        string    // host or host:port (see Hostname and Port methods)
	//     Path        string    // path (relative paths may omit leading slash)
	//     RawPath     string    // encoded path hint (see EscapedPath method)
	//     OmitHost    bool      // do not emit empty host (authority)
	//     ForceQuery  bool      // append a query ('?') even if RawQuery is empty
	//     RawQuery    string    // encoded query values, without '?'
	//     Fragment    string    // fragment for references, without '#'
	//     RawFragment string    // encoded fragment hint (see EscapedFragment method)
	// }
	baseURL:=&url.URL{
		Scheme: "https",
		Host: "example.com",
		Path:"/path",
	}

	query:=baseURL.Query()
	query.Set("name","John")
	query.Set("age","30")
	baseURL.RawQuery=query.Encode()

	fmt.Println("Built URL:", baseURL.String()) //Built URL: https://example.com/path?age=30&name=John

	//type Values map[string][]string
	values:=url.Values{}
	//add key-value pairs
	values.Add("name", "Jane")
	values.Add("age", "30")
	values.Add("city", "London")
	values.Add("country", "UK")

	//Encode
	encodedQuery:=values.Encode()

	fmt.Println(encodedQuery) //age=30&city=London&country=UK&name=Jane

	//build a URL
	baseURL1 :="https://example.com/search"
	fullURL:=baseURL1+"?"+encodedQuery

	fmt.Println(fullURL) //https://example.com/search?age=30&city=London&country=UK&name=Jane
}
