package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func UploadFile(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Upload file
		file, handler, err := r.FormFile("image")

		// JIKA VALUE PROPERTI TDK SESUAI
		if err != nil {
			fmt.Println(err)
			json.NewEncoder(w).Encode("Error file doesn't match!")
			return
		}
		defer file.Close()
		const MAX_UPLOAD_SIZE = 10 << 20 // 10MB

		// 10MB MAXSIZE
		r.ParseMultipartForm(MAX_UPLOAD_SIZE)
		if r.ContentLength > MAX_UPLOAD_SIZE {
			w.WriteHeader(http.StatusBadRequest)
			response := Result{Code: http.StatusBadRequest, Message: "Max size in 10mb"}
			json.NewEncoder(w).Encode(response)
			return
		}

		// JIKA FOLDER UPLOADS TIDAK ADA
		tempFile, err := ioutil.TempFile("uploads", "image-*"+handler.Filename)
		if err != nil {
			fmt.Println(err)
			fmt.Println("PATH UPLOAD NOT FOUND!")
			json.NewEncoder(w).Encode(err)
			return
		}
		defer tempFile.Close()

		// read all of the contents of our uploaded file into a
		// byte array
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}

		// write this byte array to our temporary file
		tempFile.Write(fileBytes)

		data := tempFile.Name()
		// filename := data[8:] // split uploads/

		// add filename to ctx
		ctx := context.WithValue(r.Context(), "dataFile", data)
		next.ServeHTTP(w, r.WithContext(ctx))

		// updatectx := context.WithValue(r.Context(), "dataUpdate", filename)
		// next.ServeHTTP(w, r.WithContext(updatectx))
	})
}
