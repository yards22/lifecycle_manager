package main

// func (app *App) handleUploadImage(rw http.ResponseWriter, r *http.Request) {

// 	r.ParseMultipartForm(10 << 20)

// 	// get handler for filename, size and headers
// 	file, _, err := r.FormFile("image")
// 	if err != nil {
// 		fmt.Println(err)
// 		sendErrorResponse(rw, http.StatusBadRequest, nil, "invalid name/ file missing")
// 		return
// 	}
// 	defer file.Close()

// 	// allow only png
// 	// if !(strings.Contains(handler.Header.Get("Content-Type"), "jpeg")) {
// 	// 	sendErrorResponse(rw, http.StatusBadRequest, nil, "invalid image format, only .png allowed")
// 	// 	return
// 	// }

// 	// putting image in image store
// 	//TODO: change the filename
// 	fileName := "rithvik_" + ".png"
// 	s3ImageUrl, err := app.objectStore.Put(file, fileName)

// 	if err != nil {
// 		sendErrorResponse(rw, http.StatusInternalServerError, nil, err.Error())
// 		return
// 	}
// 	sendResponse(rw, 200, s3ImageUrl, "ok")

// }
