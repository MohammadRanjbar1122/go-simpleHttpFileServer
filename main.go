/*
MIT License

Copyright (C) 2024 MohammadRanjbar

This code is created by MohammadRanjbar and shared on GitHub: https://github.com/MohammadRanjbar1122

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

1. The original author's name ("MohammadRanjbar") and the GitHub link (https://github.com/MohammadRanjbar1122) must remain visible in all copies or substantial portions of the code.
2. Credit to the original author must be provided in any derivative works.

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

The software is provided "AS IS", without warranty of any kind, express or implied, including but not limited to the warranties of merchantability, fitness for a particular purpose, and noninfringement. In no event shall the authors or copyright holders be liable for any claim, damages, or other liability, whether in an action of contract, tort, or otherwise, arising from, out of, or in connection with the software or the use or other dealings in the software.

By using this code, you agree to respect the conditions mentioned above.

Thank you for your understanding and support!
*/

package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	version = "v1.0.0"
)

var path string
var domain string
var oss string = runtime.GOOS

func useFake(args ...any) {}

func logRequestHandler(domain string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		log.Printf("Request from IP: %s", ip)
		w.Header().Add("Server", "HTTP FILE SS 1.0.0")
		if domain == strings.Split(r.Host, ":")[0] {
			// TODO : You can use the status code 200 here ❤
			//w.WriteHeader(200)
			http.FileServer(http.Dir(path)).ServeHTTP(w, r)
		} else if domain == "_" {
			// TODO : You can use the status code 200 here ❤
			//w.WriteHeader(200)
			http.FileServer(http.Dir(path)).ServeHTTP(w, r)
		} else {
			w.WriteHeader(403)
			w.Write([]byte(""))
		}
	})
}

func main() {
	useFake(version)
	port := os.Getenv("PORT")
	path = os.Getenv("PATH_DIR")
	domain = os.Getenv("DOMAIN_L")
	// TODO : For example, you can write ssl settings for the server.

	if port == "" {
		port = "3000"
	}

	if path == "" {
		switch oss {
		case "linux":
			path = "/"
		default:
			path = filepath.Join(strings.ToLower(filepath.VolumeName(os.Getenv("windir"))), "/")
		}
	}

	if domain == "" {
		domain = "_"
	}

	log.Println("File Server Test Start On 0.0.0.0:" + port + " ...")
	log.Println("Waiting on path : ", path)

	go func() {
		http.Handle("/", logRequestHandler(domain))
		err := http.ListenAndServe("0.0.0.0:"+port, nil)
		if err != nil {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	select {}
}
