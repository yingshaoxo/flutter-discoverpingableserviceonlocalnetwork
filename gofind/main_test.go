package GoFind

import (
	"log"
	"testing"
)

func TestGetPortsFromAHost(t *testing.T) {
	//var found string = ScanPorts("localhost", 0, 65535)
	var found string = ScanPorts("localhost", 5000, 5002, 500)
	log.Println(found)
	if found == "" {
		t.Fatalf("there should has ports")
	}
}

func TestGetAllHostsFromNetwork(t *testing.T) {
	//hosts := get_all_hosts_of_a_network("192.168.0.0/16")
	hosts := get_all_hosts_of_a_network("192.168.50.0/24")
	log.Println(hosts)
	if len(hosts) < 1 {
		t.Fatalf("we should get a lot of hosts here")
	}
}

func TestGetAllReachablePortsFromNetwork(t *testing.T) {
	urls := ScanAllHosts("192.168.49.0/24", 5000, 5030, 500)
	log.Println(urls)
	if len(urls) < 1 {
		t.Fatalf("we should get a lot of hosts here")
	}
}

func TestGetAllReachablePortsFromNetwork2(t *testing.T) {
	urls := ScanAllHosts("192.168.115.203/16", 5000, 5030, 500)
	log.Println(urls)
	if len(urls) < 1 {
		t.Fatalf("we should get a lot of hosts here")
	}
}

func TestFakePing(t *testing.T) {
	if FakePing("localhost:80") == false {
		t.Fatalf("localhost:80 should have something")
	}
}

func TestHttpGet(t *testing.T) {
	log.Println(HttpGet("http://192.168.49.32:1919/play", 1000))
}

func TestHttpPost(t *testing.T) {
	log.Println(HttpPost("http://192.168.49.32:1919/play_post", "{}", 1000))
}

func Test_post_to_the_host(t *testing.T) {
	log.Println(Post_to_the_host("192.168.49.32", 0, 20000, "/play_post", "{}", 500))
}

func Test_post_to_the_network(t *testing.T) {
	Post_to_the_network("192.168.49.0/24", 0, 20000, "/play_post", "{}", 10000)
}
