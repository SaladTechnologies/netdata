package salad

import (
	"crypto/tls"
	"encoding/json"
	"errors"
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

type Counters struct {
	Error           int64 `json:"Error"`
	ErrorClosed     int64 `json:"ErrorClosed"`
	Streaming       int64 `json:"Streaming"`
	StreamingClosed int64 `json:"StreamingClosed"`
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
	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}
	payload, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var nodes []Node
	if err = json.Unmarshal(payload, &nodes); err != nil {
		return err
	}
	for _, status := range knownStatuses {
		mx[status] = 0
	}

	for _, dest := range knownDestinations {
		mx[dest] = 0
	}

	for _, node := range nodes {
		if slices.Index(knownStatuses, node.Status) == -1 {
			return fmt.Errorf("unknown node status: %s", node.Status)
		}
		mx[node.Status]++
		for _, dest := range knownDestinations {
			v, ok := node.DCSummary[dest]
			if v && ok {
				mx[dest]++
			}
		}
	}

	return nil
}

func (c *Client) CollectCounters(mx map[string]int64) error {
	url := fmt.Sprintf("https://%s:8443/counters?raw", c.ipAddress)
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}

	payload, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var counters Counters
	if err = json.Unmarshal(payload, &counters); err != nil {
		return err
	}
	mx["streams.active"] = counters.Streaming - counters.StreamingClosed
	return nil
}
