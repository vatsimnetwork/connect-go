# Go implementation for VATSIM Connect

[![Codacy Badge](https://api.codacy.com/project/badge/Grade/f4d08015b33b4c2eadbbaa7aeb808d6d)](https://www.codacy.com/gh/vatsimnetwork/connect-go?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=vatsimnetwork/connect-go&amp;utm_campaign=Badge_Grade)

Code from [MatanBudimir](https://github.com/MatanBudimir) and [philsjh](https://github.com/philsjh).

<h3>Installation</h3>

<code>
    go get "github.com/vatsimnetwork/connect-go"
</code>

Or

<code>
import "github.com/vatsimnetwork/connect-go"
</code>

<h3>Example Code</h3>
```go

package main

import (
	connect "connect-go"
	con "connect-go/handle"
	"fmt"
	"log"
	"net/http"
)

func callback(w http.ResponseWriter, r *http.Request) {

	user, err := con.Callback(w, r)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(user.Data)

}

func main() {
	connect.URL()
	http.HandleFunc("/connect/callback", callback)
	http.ListenAndServe(":69", nil)
}
```
