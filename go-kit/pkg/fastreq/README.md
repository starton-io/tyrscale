# fastreq

fastreq is a fast http client for golang.


example:

```go
package main

import (
	"fmt"
	"github.com/starton-io/tyrscale/go-kit/pkg/fastreq"
)

func main() {
	httpClient := fastreq.NewBuilder().
		SetResponseTimeout(time.Duration(cfg.ReadTimeout) * time.Second).
		Build()
	resp, err := httpClient.Get("https://www.baidu.com", map[string]string{
		"Content-Type": "application/json",
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
}
```
