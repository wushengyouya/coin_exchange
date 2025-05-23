package tools

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"
)

func Post(url string, params any) ([]byte, error) {
	marshal, _ := json.Marshal(params)
	s := string(marshal)
	reqBody := strings.NewReader(s)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, url, reqBody)
	if err != nil {
		return nil, err
	}
	httpRsp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()
	rspBody, err := io.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}
	return rspBody, nil
}
