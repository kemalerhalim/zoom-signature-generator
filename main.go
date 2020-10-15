package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strconv"
	"time"
)

func main() {
	var meeting = meetingData{
		APIKey:        "YOUR_API_KEY",
		APISecret:     "YOUR_API_SECRET",
		MeetingNumber: "YOUR_MEETING_NUMBER",
		Role:          "0", // 	1 for meeting host, 0 for participants & joining webinars
	}
	fmt.Println(generateSignature(meeting))
}

type meetingData struct {
	APIKey        string
	APISecret     string
	MeetingNumber string
	Role          string
}

func generateSignature(data meetingData) string {
	var timeStamp = (time.Now().UTC().UnixNano() / 1e6) - 30000
	var timeStampStr = strconv.Itoa(int(timeStamp))

	var msg = data.APIKey + data.MeetingNumber + timeStampStr + data.Role
	sEnc := base64.StdEncoding.EncodeToString([]byte(msg))
	h := hmac.New(sha256.New, []byte(data.APISecret))
	h.Write([]byte(sEnc))

	var hash = base64.StdEncoding.EncodeToString(h.Sum(nil))

	var sigNatureStr = data.APIKey + "." + data.MeetingNumber + "." + timeStampStr + "." + data.Role + "." + hash
	sEnc = base64.StdEncoding.EncodeToString([]byte(sigNatureStr))
	return sEnc
}
