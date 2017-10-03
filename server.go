package main

import (
	"net/http"
	"html/template"
	"regexp"
	"fmt"
)

func exampleHandler( w http.ResponseWriter, r *http.Request){

	switch r.Method {
	case http.MethodGet:
		t, err := template.ParseFiles("html/submit.gtpl")
		if err != nil {
			http.Error(w, "Error processing template", http.StatusInternalServerError)
		}
		t.Execute(w, nil)
		break
	case http.MethodPost:
		idTax := r.PostFormValue("taxId")
		pattern := "[0-3][0-9][01][0-9]{3}[0-9]{5}"
		res, err := regexp.MatchString(pattern, idTax)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		if res {
			fmt.Fprintln(w, "Pattern matched")
			return
		}
		fmt.Fprintln(w, "Invalid fødselsnummer")
		break
	default: http.Error(w, "Not implemented", http.StatusNotImplemented)
	}

}

func main(){
	http.HandleFunc("/submit", exampleHandler)
	http.ListenAndServe(":8080", nil)

}
