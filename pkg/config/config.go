package config

import (
	"encoding/json"
	"os"
)

const (
	PINS_CONF_FILE   = "/src/github.com/alistair-english/DRC2019/config/pins.json"
	CV_CONF_FILE     = "/src/github.com/alistair-english/DRC2019/config/cv.json"
	SERIAL_CONF_FILE = "/src/github.com/alistair-english/DRC2019/config/serial.json"
)

// PinConfig is the datatype for control pin information
type PinConfig struct {
	SteeringPin int `json:"steeringPin"`
	DrivePin    int `json:"drivePin"`
	MaxSpeed    int `json:"maxSpeed"`
}

// HSV is the datatype for a hsv value
type HSV struct {
	H float64 `json:"H"`
	S float64 `json:"S"`
	V float64 `json:"V"`
}

//CVConfig is the datatype for the CV information
type CVConfig struct {
	LeftLower  HSV `json:"leftLower"`
	LeftUpper  HSV `json:"leftUpper"`
	RightLower HSV `json:"rightLower"`
	RightUpper HSV `json:"rightUpper"`
}

// SerialConfig is the datatype for the Serial information
type SerialConfig struct {
	Port      string `json:"port"`
	Baud      int    `json:"baud"`
	TimeoutMs int    `json:"timeoutMs"`
}

// GetPinConfig gets the pin configuration information from a json file
func GetPinConfig() PinConfig {
	var pins PinConfig
	pinsConfigFile, err := os.Open(os.Getenv("GOPATH") + PINS_CONF_FILE)
	defer pinsConfigFile.Close()
	if err != nil {
		panic(err)
	}
	jsonParser := json.NewDecoder(pinsConfigFile)
	jsonParser.Decode(&pins)
	return pins
}

// GetCVConfig gets the CV configuration information from a json file
func GetCVConfig() CVConfig {
	var cv CVConfig
	cvConfigFile, err := os.Open(os.Getenv("GOPATH") + CV_CONF_FILE)
	defer cvConfigFile.Close()
	if err != nil {
		panic(err)
	}
	jsonParser := json.NewDecoder(cvConfigFile)
	jsonParser.Decode(&cv)
	return cv
}

// GetSerialConfig gets the camera configuration information from a json file
func GetSerialConfig() SerialConfig {
	var scf SerialConfig
	serialConfigFile, err := os.Open(os.Getenv("GOPATH") + SERIAL_CONF_FILE)
	defer serialConfigFile.Close()
	if err != nil {
		panic(err)
	}
	jsonParser := json.NewDecoder(serialConfigFile)
	jsonParser.Decode(&scf)
	return scf
}
