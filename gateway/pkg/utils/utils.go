package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"

	"github.com/valyala/fasthttp"
)

func GetSHA256Checksum(filePath string, checksum string) (bool, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return false, err
	}
	defer file.Close()

	h := sha256.New()
	if _, err := io.Copy(h, file); err != nil {
		return false, err
	}
	computedChecksum := hex.EncodeToString(h.Sum(nil))
	return computedChecksum == checksum, nil
}

func DownloadFile(filepath string, fileUrl string, headers map[string]string) (err error) {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(fileUrl)

	// Set headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Correctly use fasthttp.Do to send the request and populate the response
	err = fasthttp.Do(req, resp)
	if err != nil {
		return err
	}

	// Check server response
	if resp.StatusCode() != fasthttp.StatusOK {
		return fmt.Errorf("bad status: %d", resp.StatusCode())
	}

	// Write the body to file
	_, err = out.Write(resp.Body())
	if err != nil {
		return err
	}

	return nil
}
