package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		html := `
			<!DOCTYPE html>
			<html>
			<head>
				<title>Cross-Origin Example</title>
				<script>
					function makeCrossOriginRequest() {
						var xhr = new XMLHttpRequest();
						xhr.open("GET", "http://server2.com:8082/api/data", true);
						xhr.onreadystatechange = function() {
							if (xhr.readyState === 4 && xhr.status === 200) {
								console.log(xhr.responseText);
							}
						};
						xhr.send();
					}
				</script>
			</head>
			<body>
				<h1>Cross-Origin Example</h1>
				<button onclick="makeCrossOriginRequest()">Make Cross-Origin Request</button>
			</body>
			</html>
		`

		fmt.Fprintf(w, html)
	})

	err := http.ListenAndServe("server1.com:8081", nil)
	if err != nil {
		panic(err)
	}

}
