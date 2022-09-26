package handlers

import (
	"net/http"

	checkRequest "ascii/internal/checkRequest"
	"ascii/internal/errors"
)

func Home(w http.ResponseWriter, r *http.Request) {
	errCheckMethod := checkRequest.CheckMethodHome(w, r)
	if errCheckMethod != 200 {
		check := checkRequest.CheckStatus(errCheckMethod)
		w.WriteHeader(errCheckMethod)
		errors.CheckErrors(w, check)
		return
	}

	errCheckPath := checkRequest.CheckPath(w, r, "/")
	if errCheckPath != 200 {
		check := checkRequest.CheckStatus(errCheckPath)
		w.WriteHeader(errCheckPath)
		errors.CheckErrors(w, check)
		return
	}

	errParsFiles := checkRequest.ParsFiles(w, "./ui/templates/index.html")
	if errParsFiles != 200 {
		check := checkRequest.CheckStatus(errParsFiles)
		w.WriteHeader(errParsFiles)
		errors.CheckErrors(w, check)
		return
	}
}

func Ascii(w http.ResponseWriter, r *http.Request) {
	errCheckMethod := checkRequest.CheckMethodAscii(w, r)
	if errCheckMethod != 200 {
		check := checkRequest.CheckStatus(errCheckMethod)
		w.WriteHeader(errCheckMethod)
		errors.CheckErrors(w, check)
		return
	}

	errCheckPath := checkRequest.CheckPath(w, r, "/ascii-art")
	if errCheckPath != 200 {
		check := checkRequest.CheckStatus(errCheckPath)
		w.WriteHeader(errCheckPath)
		errors.CheckErrors(w, check)
		return
	}

	errAscii := checkRequest.AsciiCheck(w, r)
	if errAscii != 200 {
		check := checkRequest.CheckStatus(errAscii)
		w.WriteHeader(errAscii)
		errors.CheckErrors(w, check)
		return
	}
}

func About(w http.ResponseWriter, r *http.Request) {
	errCheckMethod := checkRequest.CheckMethodHome(w, r)
	if errCheckMethod != 200 {
		check := checkRequest.CheckStatus(errCheckMethod)
		w.WriteHeader(errCheckMethod)
		errors.CheckErrors(w, check)
		return
	}

	errCheckPath := checkRequest.CheckPath(w, r, "/about.html")
	if errCheckPath != 200 {
		check := checkRequest.CheckStatus(errCheckPath)
		w.WriteHeader(errCheckPath)
		errors.CheckErrors(w, check)
		return
	}

	errParsFiles := checkRequest.ParsFiles(w, "./ui/templates/about.html")
	if errParsFiles != 200 {
		check := checkRequest.CheckStatus(errParsFiles)
		w.WriteHeader(errParsFiles)
		errors.CheckErrors(w, check)
		return
	}
}
