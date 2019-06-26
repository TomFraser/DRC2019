package cameraservice

import "gocv.io/x/gocv"

// FileReaderCamera is a camera implementation that serves up images from a video
type FileReaderCamera struct {
	Path    string
	Capture *gocv.VideoCapture
}

// NewFileReaderCameraImplementation creates a new camera implementation that will serve up from a video
func NewFileReaderCameraImplementation(path string) *FileReaderCamera {
	vid, _ := gocv.VideoCaptureFile(path)
	return &FileReaderCamera{path, vid}
}

// RunCameraConnection from camera Implementation
func (cam FileReaderCamera) RunCameraConnection(imgRequests <-chan GetImageActionReq) {
	for req := range imgRequests {
		cam.Capture.Read(req.img)
		req.responseChannel <- true
	}
}
