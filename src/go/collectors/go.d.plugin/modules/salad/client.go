package salad

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"slices"
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

func (c *Client) CollectHealth(mx map[string]int64) error {
	url := fmt.Sprintf("https://%s:8443/dump/health", c.ipAddress)
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	payload, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var nodes []Node
	if err = json.Unmarshal(payload, &nodes); err != nil {
		return err
	}
	for _, status := range knownStatuses {
		key := fmt.Sprintf("status.%s", status)
		mx[key] = 0
	}

	for _, dest := range knownDestinations {
		key := fmt.Sprintf("destination.%s", dest)
		mx[key] = 0
	}

	for _, node := range nodes {
		if slices.Index(knownStatuses, node.Status) == -1 {
			return fmt.Errorf("unknown node status: %s", node.Status)
		}
		key := fmt.Sprintf("status.%s", node.Status)
		mx[key]++
		for _, dest := range knownDestinations {
			v, ok := node.DCSummary[dest]
			if v && ok {
				key := fmt.Sprintf("destination.%s", dest)
				mx[key]++
			}
		}
	}

	return nil
}
