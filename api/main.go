package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const filesDir = "./flies"

func main() {
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/api/health", healthHandler)
	http.HandleFunc("/api/vpns", vpnsHandler)
	http.HandleFunc("/api/vpns/random", randomVpnHandler)
	http.Handle("/files/", http.StripPrefix("/files/", http.HandlerFunc(filesHandler)))

	port := getPort()
	log.Printf("Server running on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("status: online"))
}

func vpnsHandler(w http.ResponseWriter, r *http.Request) {
	files, err := getOvpnFiles()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error reading files"))
		return
	}
	baseURL := getBaseURL(r)
	for _, f := range files {
		fmt.Fprintf(w, "%s/files/%s\n", baseURL, f)
	}
}

func randomVpnHandler(w http.ResponseWriter, r *http.Request) {
	files, err := getOvpnFiles()
	if err != nil || len(files) == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("no files found"))
		return
	}
	baseURL := getBaseURL(r)
	idx := rand.Intn(len(files))
	fmt.Fprintf(w, "%s/files/%s\n", baseURL, files[idx])
}

func filesHandler(w http.ResponseWriter, r *http.Request) {
	name := filepath.Base(r.URL.Path)
	filePath := filepath.Join(filesDir, name)
	f, err := os.Open(filePath)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("file not found"))
		return
	}
	defer f.Close()
	w.Header().Set("Content-Disposition", "attachment; filename="+name)
	w.Header().Set("Content-Type", "application/octet-stream")
	io.Copy(w, f)
}

func getOvpnFiles() ([]string, error) {
	dir, err := os.Open(filesDir)
	if err != nil {
		return nil, err
	}
	defer dir.Close()
	files, err := dir.Readdirnames(-1)
	if err != nil {
		return nil, err
	}
	var ovpnFiles []string
	for _, f := range files {
		if strings.HasSuffix(f, ".ovpn") {
			ovpnFiles = append(ovpnFiles, f)
		}
	}
	return ovpnFiles, nil
}

func getBaseURL(r *http.Request) string {
	proto := "http"
	if r.TLS != nil {
		proto = "https"
	}
	if forwardedProto := r.Header.Get("X-Forwarded-Proto"); forwardedProto != "" {
		proto = forwardedProto
	}
	host := r.Host
	return fmt.Sprintf("%s://%s", proto, host)
}
