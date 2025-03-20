package model

import "encoding/xml"

const (
	IN_PROGRESS     = "IN_PROGRESS"
	ERROR           = "ERROR"
	READY           = "READY"
	PARTITION_READY = "PARTITION_READY"
)

type HashCrackRequest struct {
	Hash      string `json:"hash"`
	MaxLength int    `json:"maxLength"`
}

type HashCrackResponse struct {
	RequestID string
}

type HashStatusRequest struct {
	RequestID string `json:"requestId"`
}

type HashStatusResponse struct {
	Status   string   `json:"status"`
	Data     []string `json:"data"`
	Progress int      `json:"progress"`
}

type HashCrackManagerRequest struct {
	XMLName    xml.Name `xml:"CrackHashManagerRequest"`
	RequestId  string   `xml:"RequestId"`
	PartNumber int      `xml:"PartNumber"`
	PartCount  int      `xml:"PartCount"`
	Hash       string   `xml:"Hash"`
	MaxLength  int      `xml:"MaxLength"`
	Alphabet   Alphabet `xml:"Alphabet"`
}

type Alphabet struct {
	Symbols []string `xml:"symbols"`
}

type WorkerResult struct {
	RequestID string `json:"requestId"`
	Word      string `json:"word"`
}
