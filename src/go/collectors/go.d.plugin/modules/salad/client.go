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
	MachineId string          `json:"MachineId"`
	Status    string          `json:"Status"`
	DCSummary map[string]bool `json:"DCSummary"`
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
			return 0, 0, 0, fmt.Errorf("unknown node status: %s", node.Status)
		}
	}
	return active, quarnatined, zombied, nil
}

func (c *Client) CollectData() (map[string]int64, map[string]int64, error) {
	nodeMx := map[string]int64{}
	destMx := map[string]int64{}

	url := fmt.Sprintf("https://%s:8443/dump/health", c.ipAddress)
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nodeMx, destMx, err
	}
	defer resp.Body.Close()
	payload, err := io.ReadAll(resp.Body)
	if err != nil {
		return nodeMx, destMx, err
	}
	var nodes []Node
	if err = json.Unmarshal(payload, &nodes); err != nil {
		return nodeMx, destMx, err
	}
	nodeMx["active"] = 0
	nodeMx["quarantined"] = 0
	nodeMx["zombied"] = 0

	for _, dest := range destinations {
		destMx[dest] = 0
	}

	for _, node := range nodes {
		_, ok := nodeMx[node.Status]
		if !ok {
			return nodeMx, destMx, fmt.Errorf("unknown node status: %s", node.Status)
		}
		nodeMx[node.Status]++
		for _, dest := range destinations {
			v, ok := node.DCSummary[dest]
			if v && ok {
				destMx[dest]++
			}
		}
	}

	return nodeMx, destMx, nil
}
