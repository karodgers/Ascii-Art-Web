package handlers

import (
	"fmt"
	"net/http"
)

func DownloadAsciiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/405", http.StatusSeeOther)
		return
	}

	asciiArt := r.FormValue("ascii_art")
	if asciiArt == "" {
		http.Redirect(w, r, "/400", http.StatusSeeOther)
		return
	}
	contentLength := len(asciiArt)

	// headers used for triggering download
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Disposition", `attachment; filename="ascii_art.txt"`)
	w.Header().Set("Content-Length", fmt.Sprintf("%d", contentLength))

	// writing the asscii art to the response body
	fmt.Fprint(w, asciiArt)
}
