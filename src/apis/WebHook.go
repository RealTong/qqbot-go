package apis

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Test(w http.ResponseWriter, r *http.Request) {
	// 打印body内容
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(body))

}

func errorApi(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

}

func warnApi(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

}

func infoApi(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

}

func buildApi(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

}
