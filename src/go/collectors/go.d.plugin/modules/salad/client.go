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
	ipAddress net.IP
}

type Node struct {
	MachineId string `json:"MachineId"`
	Status    string `json:"Status"`
}

func NewClient() (*Client, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return &Client{
		ipAddress: localAddr.IP,
	}, nil
}

func (c *Client) GetNodeCount() (int, error) {
	transport := http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := http.Client{
		Transport: &transport,
	}
	url := fmt.Sprintf("https://%s:8443/dump/health", c.ipAddress)
	resp, err := client.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	payload, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	var nodes []Node
	if err = json.Unmarshal(payload, &nodes); err != nil {
		return 0, err
	}
	return len(nodes), nil
}
