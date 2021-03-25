package main

import (
	"encoding/base64"
	"fmt"
	qrcode "github.com/skip2/go-qrcode"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, req *http.Request) {
	//start := time.Now()
	//defer func() {
	//	total := time.Now().Sub(start).Microseconds()
	//	fmt.Println(total)
	//}()
	var png []byte
	png, err := qrcode.Encode("dclub://pushaction/?data=eyJtZXRob2QiOjYsIm9yZGVyX2hhc2giOiJkZDVkZDdlNTIwM2Y0NDM4M2RjN2M2YzBhNzljMTAwODlkZjkyZjc0In0=", qrcode.Medium, 256)
	if err != nil {
		fmt.Print(err)
	}

	res := base64.StdEncoding.EncodeToString(png)

	res = "<html><img src=\"data:image/png;base64, " + res + "\" alt=\"Red dot\" /></html>"

	w.Write([]byte(res))
}

func main() {

	fmt.Println("hello, port: ", os.Getenv("PORT"))
	http.HandleFunc("/", hello)

	http.ListenAndServe(":"+os.Getenv("PORT"), nil)

}
