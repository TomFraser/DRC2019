package main

import (
	"github.com/alistair-english/DRC2019/pkg/arch"
	"github.com/alistair-english/DRC2019/pkg/services/cameraservice"
	"github.com/alistair-english/DRC2019/pkg/services/cvservice"
)

func main() {
	router := arch.NewRouter()

	calService := &cvservice.CalibratorService{}
	camService, _ := cameraservice.NewFileReaderCamera("../recorder/recording_06-30-2019_21:42:55.avi")

	router.Register(calService)
	router.Register(camService)

	calService.Start()
	camService.Start()

	// router is blocking
	router.Start()
}
