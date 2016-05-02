package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,
		`<html>
<head>
    <title>Index Page</title>
</head>
<body>
    <h1>Index Page</h1>
    <p>Hello %s</p>
</body>
</html>
`, "%username%")
}
