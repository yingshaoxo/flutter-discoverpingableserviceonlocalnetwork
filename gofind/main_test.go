package GoFind

import (
	"log"
	"testing"
)

func TestGetPortsFromAHost(t *testing.T) {
	//var found string = ScanPorts("localhost", 0, 65535)
	var found string = ScanPorts("localhost", 5000, 5002)
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
	urls := ScanAllHosts("192.168.50.0/24", 5000, 5010)
	log.Println(urls)
	if len(urls) < 1 {
		t.Fatalf("we should get a lot of hosts here")
	}
}

func TestGetAllReachablePortsFromNetwork2(t *testing.T) {
	SetTimeOut(100)
	urls := ScanAllHosts("192.168.50.0/16", 5000, 5010)
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
