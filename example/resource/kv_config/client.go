package kv_config

import (
	"encoding/json"
	"github.com/ssst0n3/awesome_libs"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"io/ioutil"
	"net/http"
)

type Client struct {
	Address string `json:"address"`
}

func (c Client) Get(api, key string) (value string, err error) {
	url := awesome_libs.Format("{.address}{.api}/{.key}", awesome_libs.Dict{
		"address": c.Address,
		"api":     api,
		"key":     key,
	})
	resp, err := http.Get(url)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	if resp.StatusCode == http.StatusOK {
		err = json.Unmarshal(body, &value)
		if err != nil {
			awesome_error.CheckErr(err)
			return
		}
	}
	return
}
