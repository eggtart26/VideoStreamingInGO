package main

import (
	"io"
	"os"
	"io/ioutil"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func streamHandler(w http.ResponseWriter, r *http.ReadRequest, p httprouter.Params) {
	vid := p.ByName("vid-id")
	vl := VIDEO_DIR + vid

	video, err := os.Open(vl)
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "Interval Error")
		return
	}

	w.Header().Set("Content-Type", "video/mp4")
	http.ServeContent(w, r, "", time.Now(), video)

	defer video.Close()
}

func uploadHandler(w http.ResponseWriter, r *http.ReadRequest, p httprouter.Params) {

}