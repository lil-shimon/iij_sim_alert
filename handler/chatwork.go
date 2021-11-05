package handler

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/lil-shimon/iij_sim_alert/config"
)

func PostChatwork() error {
	apiToken, apiUrl := config.LoadEnv()
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

	req.Header.Add("X-ChatWorkToken", apiToken)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		fmt.Println("success")
	}

	if res.StatusCode != http.StatusOK {
		fmt.Printf("failed: %v", res)
	}

	return err
}
