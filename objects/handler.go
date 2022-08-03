package objects

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(formatRequest(r))
	fmt.Println()
	fmt.Println()

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func formatRequest(r *http.Request) string {

	var request []string

	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)

	request = append(request, fmt.Sprintf("Host: %v", r.Host))

	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}

	// 处理POST请求体
	if r.Method == "POST" {
		r.ParseForm()
		request = append(request, "\n")
		request = append(request, r.Form.Encode())
	}

	bodyBuffer, _ := ioutil.ReadAll(r.Body)
	if len(bodyBuffer) == 0 {
		fmt.Println("空内容")
	}
	request = append(request, string(bodyBuffer))

	return strings.Join(request, "\n")
}
