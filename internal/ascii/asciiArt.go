package ascii

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var charmap = make(map[rune][]string)

func AsciArt(text, banner string) (string, int) {
	readFile, err := ReadFile(banner)
	if err != http.StatusOK {
		return "", http.StatusInternalServerError
	}

	arg, err := ReadOsStdin(text)
	if err != http.StatusOK {
		return "", http.StatusBadRequest
	}

	dataFile, dataArgs := Split(readFile, arg)

	AddElemInMap(dataFile)

	return PrintAscii(dataArgs), http.StatusOK
}

func ReadOsStdin(arr string) (string, int) {
	arg := ""

	for _, w := range arr {
		if w < 32 || w > 126 {
			if w == 10 || w == 13 {
				arg += string(w)
				continue
			} else {
				return "", http.StatusBadRequest
			}
		}
		arg += string(w)
	}

	arg = strings.ReplaceAll(arg, "\r\n", "\n")
	return arg, http.StatusOK
}

func ReadFile(banner string) (string, int) {
	file, err := os.ReadFile("banners/" + banner)
	if err != nil {
		return "", http.StatusInternalServerError
	}

	errHash := HashSum(banner)
	if errHash != http.StatusOK {
		return "", http.StatusInternalServerError
	}

	return string(file), http.StatusOK
}

func Split(strFile string, strArg string) ([]string, []string) {
	return strings.Split(strFile, "\n"), strings.Split(strArg, "\n")
}

func AddElemInMap(arr []string) {
	n := 0
	temp := []string{}

	for i := 9; i < len(arr); i = i + 9 {
		temp = append(temp, arr[i-8:i]...)
		charmap[rune(32+n)] = temp
		n++
		temp = []string{}
	}
}

func CheckNewLine(args []string) string {
	count := 0

	for i := 0; i < len(args); i++ {
		if args[i] == "" {
			count++
		}
	}

	res := ""

	if count == len(args) {
		for i := 0; i < len(args)-1; i++ {
			res += "\n"
		}
		return res
	}

	return ""
}

func PrintAscii(arg []string) string {
	resultText := ""

	for j := 0; j < len(arg); j++ {
		if arg[j] == "" {
			resultText += "\n"
			continue
		}
		for i := 0; i < 8; i++ {
			for _, q := range arg[j] {
				resultText += charmap[q][i]
			}

			resultText += "\n"

		}
	}

	return resultText
}

func HashSum(path string) int {
	h := md5.New()
	f, err := os.Open("banners/" + path)
	if err != nil {
		log.Print(err)
	}
	defer f.Close()
	_, err = io.Copy(h, f)
	if err != nil {
		log.Print(err)
	}
	hashSum := fmt.Sprintf("%x", h.Sum(nil))

	if path == "standard.txt" {
		if hashSum != "ac85e83127e49ec42487f272d9b9db8b" {
			return http.StatusInternalServerError
		}
	} else if path == "shadow.txt" {
		if hashSum != "a49d5fcb0d5c59b2e77674aa3ab8bbb1" {
			return http.StatusInternalServerError
		}
	} else if path == "thinkertoy.txt" {
		if hashSum != "8efd138877a4b281312f6dd1cbe84add" {
			return http.StatusInternalServerError
		}
	}
	return http.StatusOK
}
