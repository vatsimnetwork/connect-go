package examples

import (
	"fmt"
	connect "github.com/vatsimnetwork/connect-go"
	"github.com/vatsimnetwork/connect-go/handle"
	"log"
	"net/http"
)

func callback(w http.ResponseWriter, r *http.Request) {

	user, err := Handle.Callback(w, r)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(user)

	fmt.Printf("%s (%d)", user.NameFull, user.CID)

}

func main() {
	connect.URL()
	http.HandleFunc("/connect/callback", callback)
	log.Fatal(http.ListenAndServe(":3000", nil))
}