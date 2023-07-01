package GoFind

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

func SayHi() string {
	return "Hi! yingshaoxo!"
}

func worker(address chan string, results chan string, timeout_in_milliseconds int) {
	for uri := range address {
		connection, err := net.DialTimeout("tcp", uri, time.Duration(1000*1000*timeout_in_milliseconds))
		//fmt.Println(err)
		if err != nil {
			results <- ""
			continue
		}
		connection.Close()
		//fmt.Println(uri)
		results <- uri
	}
}

func scan_ports(hosts []string, startPort int, endPort int, timeout_in_milliseconds int) []string {
	//1-65535
	urls := make([]string, 0)

	//address := make(chan string, 10000)
	address := make(chan string, 65535)
	results := make(chan string)

	for i := 0; i < cap(address); i++ {
		go worker(address, results, timeout_in_milliseconds) // now we have 10000 workers
	}

	/*
		for _, host := range hosts {
			go func(host_ string) {
				for i := startPort; i <= endPort; i++ {
					address <- fmt.Sprintf("%s:%d", host_, i)
				}
			}(host)
		}
	*/
	go func() {
		for _, host := range hosts {
			for i := startPort; i <= endPort; i++ {
				address <- fmt.Sprintf("%s:%d", host, i)
			}
		}
	}()

	for _, _ = range hosts {
		for i := startPort; i <= endPort; i++ {
			uri := <-results
			if uri != "" {
				urls = append(urls, uri)
			}
		}
	}

	close(address)
	close(results)

	return urls
}

func ScanPorts(host string, startPort int, endPort int, timeout_in_milliseconds int) string {
	var hosts = []string{host}
	urls := scan_ports(hosts, startPort, endPort, timeout_in_milliseconds)
	json_result, err := json.Marshal(urls)
	if err != nil {
		return ""
	} else {
		return string(json_result)
	}
}

func get_all_hosts_of_a_network(network string) []string {
	hosts := make([]string, 0)

	// convert string to IPNet struct
	_, ipv4Net, err := net.ParseCIDR(network) // 192.168.0.0/16
	if err != nil {
		return hosts
	}

	// convert IPNet struct mask and address to uint32
	// network is BigEndian
	mask := binary.BigEndian.Uint32(ipv4Net.Mask)
	start := binary.BigEndian.Uint32(ipv4Net.IP)

	// find the final address
	finish := (start & mask) | (mask ^ 0xffffffff)

	// loop through addresses as uint32
	for i := start; i <= finish; i++ {
		// convert back to net.IP
		ip := make(net.IP, 4)
		binary.BigEndian.PutUint32(ip, i)
		hosts = append(hosts, ip.String())
	}

	return hosts
}

func ScanAllHosts(network string, startPort int, endPort int, timeout_in_milliseconds int) string {
	hosts := get_all_hosts_of_a_network(network)

	urls := scan_ports(hosts, startPort, endPort, timeout_in_milliseconds)

	json_result, err := json.Marshal(urls)
	if err != nil {
		return ""
	} else {
		return string(json_result)
	}
}

func FakePing(address string) bool {
	connection, err := net.Dial("tcp", address)
	if err != nil {
		return false
	}
	connection.Close()
	return true
}

func HttpGet(url string, timeout_in_milliseconds int) string {
	client := http.Client{
		Timeout: time.Duration(timeout_in_milliseconds) * time.Millisecond,
	}
	res, err := client.Get(url)
	if err != nil {
		return err.Error()
	} else {
		response_body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err.Error()
		}
		return string(response_body)
	}
}

func HttpPost(url string, json_string string, timeout_in_milliseconds int) string {
	jsonBody := []byte(json_string)
	bodyReader := bytes.NewReader(jsonBody)

	reqest, err := http.NewRequest(http.MethodPost, url, bodyReader)
	if err != nil {
		return err.Error()
	}
	reqest.Header.Add("Content-Type", "application/json")

	client := &http.Client{
		Timeout: time.Duration(timeout_in_milliseconds) * time.Millisecond,
	}
	response, err := client.Do(reqest)
	if err != nil {
		return err.Error()
	}
	defer response.Body.Close()

	response_body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err.Error()
	}
	return string(response_body)
}

func post_worker(address chan string, jsonString string, results chan string, timeout_in_milliseconds int) {
	for uri := range address {
		results <- HttpPost(uri, jsonString, timeout_in_milliseconds)
	}
}

func post_to_those_hosts_ports(hosts []string, startPort int, endPort int, subURL string, jsonString string, timeout_in_milliseconds int) []string {
	//1-65535
	result_list := make([]string, 0)

	address := make(chan string, 65535)
	results := make(chan string)

	for i := 0; i < cap(address); i++ {
		go post_worker(address, jsonString, results, timeout_in_milliseconds) // now we have 10000 workers
	}

	go func() {
		for _, host := range hosts {
			for i := startPort; i <= endPort; i++ {
				if strings.HasPrefix(subURL, "/") {
					address <- fmt.Sprintf("http://%s:%d%s", host, i, subURL)
				} else {
					address <- fmt.Sprintf("http://%s:%d/%s", host, i, subURL)
				}
			}
		}
	}()

	for _, _ = range hosts {
		for i := startPort; i <= endPort; i++ {
			a_result := <-results
			result_list = append(result_list, a_result)
		}
	}

	close(address)
	close(results)

	return result_list
}

func Post_to_the_host(host string, startPort int, endPort int, subURL string, jsonString string, timeout_in_milliseconds int) string {
	var hosts = []string{host}

	result_list := post_to_those_hosts_ports(hosts, startPort, endPort, subURL, jsonString, timeout_in_milliseconds)

	json_result, err := json.Marshal(result_list)
	if err != nil {
		return ""
	} else {
		return string(json_result)
	}
}

func Post_to_the_network(network string, startPort int, endPort int, subURL string, jsonString string, timeout_in_milliseconds int) string {
	hosts := get_all_hosts_of_a_network(network)

	result_list := post_to_those_hosts_ports(hosts, startPort, endPort, subURL, jsonString, timeout_in_milliseconds)

	json_result, err := json.Marshal(result_list)
	if err != nil {
		return ""
	} else {
		return string(json_result)
	}
}
