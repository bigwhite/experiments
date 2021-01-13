package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

const resp = `<html>

<body>


<h4>Go项目排名</h4>
<table border="1">
<tr>
  <td>仓库名</td>
  <td>作者</td>
  <td>星星数</td>
</tr>
<tr>
  <td>golang</td>
  <td>Russ Cox</td>
  <td>10000</td>
</tr>
<tr>
  <td>kubernetes</td>
  <td>Brendan Burns </td>
  <td>9000</td>
</tr>
<tr>
  <td>docker</td>
  <td>Solomon Hykes</td>
  <td>8000</td>
</tr>
</table>

</body>
</html>`

func Trending(w http.ResponseWriter, r *http.Request) {
	log.Printf("orig request: %#v", *r)
	//r.ParseMultipartForm(1024)
	//log.Printf("request after parsemutipartform: %#v", *r)
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("read body error: %s", err)
		return
	}
	log.Printf("body = \n%s\n", string(b))
	w.Write([]byte(resp))
}

func main() {
	http.HandleFunc("/trending", Trending)
	var s = http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(Trending),
	}
	s.ListenAndServe()
}
