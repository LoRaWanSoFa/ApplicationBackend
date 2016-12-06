package distributor

import (
	"fmt"
	"net/http"
	"net/url"
)

type RestUplinkConnector interface {
	NewNode(nodeId string, longitude float64, latitude float64) (*http.Response, error)
	UpdateNode(nodeId string, longitude float64, latitude float64) (*http.Response, error)
	DeleteNode(nodeId string) (*http.Response, error)
	NewSensor(nodeId string, sensorId string) (*http.Response, error)
	DeleteSensor(nodeId string, sensorId string) (*http.Response, error)
	NewData(nodeId string, message url.Values) (*http.Response, error)
}

type restUplinkConnector struct {
	basicURL string
}

func NewRestUplinkConnector(ip string, apiKey string) RestUplinkConnector {
	connector := new(restUplinkConnector)
	connector.basicURL = createBasicURL(ip, apiKey)
	return connector
}

func createBasicURL(ip string, apiKey string) string {
	return fmt.Sprintf("http://%s/api/%s", ip, apiKey)
}

func (r *restUplinkConnector) NewNode(nodeId string, longitude float64, latitude float64) (*http.Response, error) {
	//Post Request:
	//http://ip/api/{apikey}/newnode/{id}/{longitude}/{latitude}
	//Needs to be called if a new node is added to the backend.
	url := fmt.Sprintf("%s/newnode/%s/%f/%f", r.basicURL, nodeId, longitude, latitude)
	resp, err := http.PostForm(url, nil)
	return resp, err
}

func (r *restUplinkConnector) UpdateNode(nodeId string, longitude float64, latitude float64) (*http.Response, error) {
	//Put Request:
	//http://ip/api/{apikey}/updatenode/{id}/{longitude}/{latitude}
	//Needs to be called if the position of a node has changed.
	url := fmt.Sprintf("%s/updatenode/%s/%f/%f", r.basicURL, nodeId, longitude, latitude)
	resp, err := http.Get(url)
	return resp, err
}

func (r *restUplinkConnector) DeleteNode(nodeId string) (*http.Response, error) {
	//Delete Request:
	//http://ip/api/{apikey}/deletenode/{id}
	//Needs to be called if a node gets removed from the backend.
	client := &http.Client{}
	url := fmt.Sprintf("%s/deletenode/%s", r.basicURL, nodeId)
	req, _ := http.NewRequest("DELETE", url, nil)
	resp, err := client.Do(req)
	return resp, err
}

func (r *restUplinkConnector) NewSensor(nodeId string, sensorId string) (*http.Response, error) {
	//Post Request:
	//http://ip/api/{apikey}/newsensor/{id}/{nodeid}
	//Needs to be called if a sensor is added to a node.
	//Does he need any other sensor information?
	// if yes -> add json again
	url := fmt.Sprintf("%s/newsensor/%s/%s", r.basicURL, nodeId, sensorId)
	resp, err := http.PostForm(url, nil)
	return resp, err
}

func (r *restUplinkConnector) DeleteSensor(nodeId string, sensorId string) (*http.Response, error) {
	//Delete Request:
	//http://ip/api/{apikey}/deletesensor/{id}
	//Needs to be called if a sensor is removed from a node.
	client := &http.Client{}
	url := fmt.Sprintf("%s/deletesensor/%s/%s", r.basicURL, nodeId, sensorId)
	req, _ := http.NewRequest("DELETE", url, nil)
	resp, err := client.Do(req)
	return resp, err
	// isn't nodeId needed as well as sensorId
}

func (r *restUplinkConnector) NewData(nodeId string, message url.Values) (*http.Response, error) {
	//Post Request:
	//http://ip/api/{apikey}/newdata
	//Needs to be called if data from a node comes in.
	// add json to it
	url := fmt.Sprintf("%s/newData/%s", r.basicURL, nodeId)
	resp, err := http.PostForm(url, message)
	// isn't nodeId needed here?
	return resp, err
}
