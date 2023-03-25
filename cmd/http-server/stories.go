package main

// func (app *App) handleUploadImage(rw http.ResponseWriter, r *http.Request) {

// 	r.ParseMultipartForm(10 << 20)

// 	// get handler for filename, size and headers
// 	file, handler, err := r.FormFile("image")
// 	if err != nil {
// 		sendErrorResponse(rw, http.StatusBadRequest, nil, "invalid name/ file missing")
// 		return
// 	}
// 	defer file.Close()

// 	// allow only png
// 	if !(strings.Contains(handler.Header.Get("Content-Type"), "png")) && !(strings.Contains(handler.Header.Get("Content-Type"), "jpg")) {
// 		sendErrorResponse(rw, http.StatusBadRequest, nil, "invalid image format, only .png allowed")
// 		return
// 	}

// 	x := r.Context().Value("user").(UserDetails)

// 	// putting image in image store
// 	fileName := x.MailId + ".png"
// 	s3ImageUrl, err := app.objectStore.Put(file, fileName)
// 	if err != nil {
// 		sendErrorResponse(rw, http.StatusInternalServerError, nil, err.Error())
// 		return
// 	}
// 	sendResponse(rw, 200, s3ImageUrl, "ok")
// 	return
// }
