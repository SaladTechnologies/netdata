package salad

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
)

type Client struct {
	ipAddress  net.IP
	httpClient http.Client
}

type Node struct {
	MachineId string `json:"MachineId"`
	Status    string `json:"Status"`
}

func NewClient() (*Client, error) {
	cl := Client{}
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	cl.ipAddress = localAddr.IP
	transport := http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	cl.httpClient = http.Client{Transport: &transport}
	return &cl, nil
}

func (c *Client) GetNodeCount() (int, int, int, error) {

	url := fmt.Sprintf("https://%s:8443/dump/health", c.ipAddress)
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return 0, 0, 0, err
	}
	defer resp.Body.Close()
	payload, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, 0, 0, err
	}
	var nodes []Node
	if err = json.Unmarshal(payload, &nodes); err != nil {
		return 0, 0, 0, err
	}
	active := 0
	quarnatined := 0
	zombied := 0
	for _, node := range nodes {
		switch node.Status {
		case "active":
			active++
		case "quarantined":
			quarnatined++
		case "zombied":
			zombied++
		default:
			return 0, 0, 0, fmt.Errorf("Unknown node status: %s", node.Status)
		}
	}
	return active, quarnatined, zombied, nil
}
