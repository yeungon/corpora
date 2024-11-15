package home

import (
	"fmt"
	"io"
	"net/http"

	"github.com/yeungon/corpora/modules/home/models"
)

func (app *Controller) SearchManticore(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
	models.Manticore()

}
