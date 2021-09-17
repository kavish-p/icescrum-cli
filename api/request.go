package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/spf13/viper"
)

func DOiceScrumPOST(requestPath string, token string, payload string) {

	base_url := viper.Get("base_url").(string)
	url := base_url + requestPath
	method := "POST"

	payload_reader := strings.NewReader(payload)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload_reader)

	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("x-icescrum-token", token)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
