package main

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"

	// "os"
	"time"

	pdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func Convert(file multipart.File, handler *multipart.FileHeader, outFile io.Writer) {
	fmt.Println("start converting")

	defer file.Close()

	pdfg, err := pdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}
	page := pdf.NewPageReader(file)
	pdfg.AddPage(page)
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	// Write buffer contents to file on disk
	_, err = outFile.Write(pdfg.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("finished converting")

}

func HandleConvert(w http.ResponseWriter, r *http.Request) {
	fmt.Println("received request")
	fmt.Println(r.Form, "\n", r.Body, "\n", r.Header)
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	file, handler, err := r.FormFile("index.html")
	if err != nil {
		fmt.Println(handler)
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	Convert(file, handler, w)
	fmt.Println("finished request")

}

func HealtCheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("OK"))
	w.WriteHeader(http.StatusOK)
}

func main() {
	fmt.Println("Starting server")
	mux := http.NewServeMux()
	mux.HandleFunc("/convert", HandleConvert)
	mux.HandleFunc("/healthcheck", HealtCheck)


	s := &http.Server{
		Addr:           "0.0.0.0:8082",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Println("Server started")
	log.Fatal(s.ListenAndServe())
}
