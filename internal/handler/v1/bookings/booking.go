package bookings

import (
	"io"
	"net/http"
)

func Test(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}
