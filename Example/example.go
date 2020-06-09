package Example

import (
	connect "github.com/vatsimnetwork/connect-go"
	"github.com/vatsimnetwork/connect-go/handle"
	"fmt"
	"log"
	"net/http"
)

func callback(w http.ResponseWriter, r *http.Request) {

	user, err := Handle.Callback(w, r)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(user)

	fmt.Printf("%v (%v)", user.NameFull, user.CID)

}

func main() {
	connect.URL()
	http.HandleFunc("/connect/callback", callback)
	http.ListenAndServe(":3000", nil)
}