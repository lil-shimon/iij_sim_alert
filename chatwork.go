package iij_sim_alert

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const apiToken = "db7a413c746d1fb3defb60e10a7e6382"
const apiUrl = "https://www.chatwork.com/#!rid205437150"

func postChatwork() error {
	values := url.Values{}
	values.Set("body", "SIM通信用アラート\ntest")

	req, err := http.NewRequest(
		"POST",
		apiUrl,
		strings.NewReader(values.Encode()),
		)

	if err != nil {
		return err
	}

	req.Header.Set("X-ChatWorkToken", apiToken)
	req.Header.Set("Content-Type", "text/html; charset=utf-8")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("success")
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println("failed")
	}

	return err
}
