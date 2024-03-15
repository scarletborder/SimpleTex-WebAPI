package req

import (
	"bytes"
	"io"
	"net/http"
	url "net/url"

	"github.com/scarletborder/SimpleTex-WebAPI/crypto"

	"github.com/scarletborder/SimpleTex-WebAPI/constant"
)

const FORMULA, DOCUMENT, AUTO = "formula", "document", "auto"

// type RecConfig struct {
// 	// Currently, we detect there are three recognition modes could be used
// 	// "formula", "document", "auto"
// 	RecMode string `default:"auto"`
// 	// proxyUrl is like "http://127.0.0.1:7890"
// 	ProxyUrl string `default:""`
// }

func UploadOCR(body *bytes.Buffer, file_type string) (string, error) {
	var client *http.Client

	if constant.Config().Proxies != "" {
		proxy, err := url.Parse(constant.Config().Proxies)
		if err != nil {
			return "", err
		}
		transport := &http.Transport{
			Proxy: http.ProxyURL(proxy),
		}
		client = &http.Client{
			Transport: transport,
		}
	} else {
		client = &http.Client{}
	}
	req_url, err := crypto.Get_request_url()
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("POST", req_url, body)
	if err != nil {
		return "", err
	}

	// 设置Header
	req.Header.Set("Content-Type", file_type)
	req.Header.Set("authority", "server.simpletex.cn")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("origin", "https://simpletex.cn")
	req.Header.Set("referer", "https://simpletex.cn/")
	req.Header.Set("sec-ch-ua", "Not")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", "Windows")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-site")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	text, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	ret, err := crypto.Parse_result(string(text))
	if err != nil {
		return "", err
	}
	return ret, nil
}
