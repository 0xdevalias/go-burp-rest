# go-burp-rest

(**tl;dr:** Golang API client for [vmware/burp-rest-api](https://github.com/vmware/burp-rest-api))

Artisanally hand crafted, short pour, small batch API client for consuming [vmware/burp-rest-api](https://github.com/vmware/burp-rest-api) from Golang. Designed for the discerning DevSecOpCyberBuzzword professional who can't bring themselves to interact with the monstrosity produced by [swagger-codegen](https://github.com/swagger-api/swagger-codegen).

**Note:** There are a few hacky bits, particularly in the helpers.. Use at your own risk (or contribute a fix!)

## Usage

Dependencies are managed with [dep](https://github.com/golang/dep):

```
dep ensure
```

Run a sample program:

```
func main() {
	c := DefaultClient("http://127.0.0.1:8090") // Set to where your API is listening

	v, err := c.BurpVersion()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Burp Version: %+v\n", v)
}
```

There is some more hacky code in `main.go` for 'test in main' purposes.

## Improvements

* Implement get/set config endpoints
* TEST ALL THE THINGS!
* https://github.com/vmware/burp-rest-api/issues/35
* Add a CLI using [spf13/cobra](https://github.com/spf13/cobra)

## You may also enjoy

* [0xdevalias/docker-burp-rest-api](https://github.com/0xdevalias/docker-burp-rest-api) : Easily build/run vmware/burp-rest-api in Docker.
