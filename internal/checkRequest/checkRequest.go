package checkRequest

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"ascii/internal/ascii"
)

func CheckMethodHome(w http.ResponseWriter, r *http.Request) int {
	if r.Method != http.MethodGet {
		return http.StatusMethodNotAllowed
	}
	return http.StatusOK
}

func CheckMethodAscii(w http.ResponseWriter, r *http.Request) int {
	if r.Method != http.MethodPost {
		if r.Method == http.MethodGet {
			return http.StatusNotFound
		}
		return http.StatusMethodNotAllowed
	}

	return http.StatusOK
}

func CheckPath(w http.ResponseWriter, r *http.Request, path string) int {
	if r.URL.Path != path {
		return http.StatusNotFound
	}

	return http.StatusOK
}

func ParsFiles(w http.ResponseWriter, path string) int {
	ts, err := template.ParseFiles(path)
	if err != nil {
		return http.StatusInternalServerError
	}

	d := struct {
		Data string
		Col  string
	}{}

	err = ts.Execute(w, d)

	if err != nil {
		return http.StatusInternalServerError
	}

	return http.StatusOK
}

func AsciiCheck(w http.ResponseWriter, r *http.Request) int {
	Input := r.FormValue("input")
	banner := r.FormValue("banner")

	Res, errAscii := ascii.AsciArt(Input, banner)
	btn := r.FormValue("download")

	if errAscii != http.StatusOK {
		if errAscii == http.StatusBadRequest {
			return http.StatusBadRequest
		} else if errAscii == http.StatusInternalServerError {
			return http.StatusInternalServerError
		}
	}

	if btn == "download" {
		w.Header().Set("Content-Disposition", "attachment; filename=ASCII-ART")
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Content-Length", strconv.Itoa(len(Res)))
		w.Write([]byte(Res))
		return http.StatusOK
	}

	ts, err := template.ParseFiles("./ui/templates/index.html")
	if err != nil {
		return http.StatusInternalServerError
	}

	d := struct {
		Data string
		Col  string
	}{
		Data: Res,
		Col:  Input,
	}

	err = ts.Execute(w, d)

	if err != nil {
		return http.StatusInternalServerError
	}

	return http.StatusOK
}

func CheckStatus(status int) string {
	switch status {
	case http.StatusBadRequest:
		return fmt.Sprintf("%d %s", http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	case http.StatusNotFound:
		return fmt.Sprintf("%d %s", http.StatusNotFound, http.StatusText(http.StatusNotFound))
	case http.StatusMethodNotAllowed:
		return fmt.Sprintf("%d %s", http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
	case http.StatusInternalServerError:
		return fmt.Sprintf("%d %s", http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}
	return ""
}
