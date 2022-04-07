package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
)

const PasteBinCutoff = 50

const PasteBinHostname = "https://0x0.st"

func UploadToPastebin(str string) string {
	/*
		Input = text of command output,
		would have been printed out. To be
		uploaded to the pastebin

		Output = URL to item in pastebin
	*/

	var pars []string
	for _, par := range strings.Split(str, "\n") {
		if strings.TrimSpace(par) != "" {
			pars = append(pars, strings.TrimSpace(par))
		}
	}
	str = strings.Join(pars, "\n")

	buf := new(bytes.Buffer)

	w := multipart.NewWriter(buf)
	f, err := w.CreateFormFile("file", "output.txt")
	if err != nil {
		return fmt.Sprintf("Error during multipart form %+v", err)
	}
	if _, err = f.Write([]byte(str)); err != nil {
		return fmt.Sprintf("Error during multipart form %+v", err)
	}
	w.Close()

	req, err := http.NewRequest("POST", PasteBinHostname, (buf))
	if err != nil {
		return fmt.Sprintf("Error during New Request: %+v", err)
	}
	req.Header.Set("Content-Type", w.FormDataContentType())

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return fmt.Sprintf("Error during Upload request: %+v", err)
	}
	defer resp.Body.Close()

	out, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Sprintf("Error during pastebin upload: %+v\n", err)
	}

	return string(out)

}
