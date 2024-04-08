package Business

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/TSF/TFSClient/Commons"
	"github.com/TSF/TFSClient/Commons/Constant"
)

func WritteAlert(request Commons.NameNodeRequest) (http.Response, error) {
	http.ListenAndServe(Constant.PortPrin, nil)
	res, err := http.Post(Constant.Ip_prin+"/namenode/add/metadata", "application/json", request)
	print(err)
	if err != nil {
		fmt.Println(err)
		return http.Response{Status: "404"}, err
	}
	if res == nil {

		res, err := http.Post(Constant.Ip_prin+"/namenode/add/metadata", "application/json", request)
		if err != nil {
			fmt.Println(err)
			return *res.Request.Response, err
		}
	}
	if res == nil {
		res, err := http.Post(Constant.Ip_sec+"/namenode/add/metadata", "application/json", request)
		if err != nil {
			fmt.Println(err)
			return *res.Request.Response, err
		}
	}
	return *res.Request.Response, err
}

func GetReplicas() (Commons.DataNodesList, error) {

	var response Commons.DataNodesList

	var res *http.Response

	res, err := http.Get(Constant.Ip_prin + "/namenode/get/replicas")

	if res == nil {

		res, err := http.Get(Constant.Ip_prin + "/namenode/get/replicas")
		if err != nil {
			fmt.Println(err)
			defer res.Body.Close()
			body, err := ioutil.ReadAll(res.Body)

			json.Unmarshal(body, response)
			return response, err
		}
	}
	if res == nil {
		res, err := http.Get(Constant.Ip_sec + "/namenode/get/replicas")
		if err != nil {
			fmt.Println(err)
			defer res.Body.Close()
			body, err := ioutil.ReadAll(res.Body)

			json.Unmarshal(body, response)
			return response, err
		}
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	json.Unmarshal(body, response)
	return response, err
}

func GetMetadata(fileName string) (Commons.NameNodeResponse, error) {
	var response Commons.NameNodeResponse

	res, err := http.Get(Constant.Ip_prin + "/namenode/get/metadata/" + fileName)

	if res == nil {

		res, err := http.Get(Constant.Ip_prin + "/namenode/get/metadata/" + fileName)
		if err != nil {
			fmt.Println(err)
			defer res.Body.Close()
			body, err := ioutil.ReadAll(res.Body)

			json.Unmarshal(body, response)
			return response, err
		}
	}
	if res == nil {
		res, err := http.Get(Constant.Ip_sec + "/namenode/get/metadata/" + fileName)
		if err != nil {
			fmt.Println(err)
			defer res.Body.Close()
			body, err := ioutil.ReadAll(res.Body)

			json.Unmarshal(body, response)
			return response, err
		}
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	json.Unmarshal(body, response)
	return response, err
}
