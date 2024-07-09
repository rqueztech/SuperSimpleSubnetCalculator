package network

import (
    "testing"
)

func TestCalculateClass(t *testing.T) {
    // Class A addresses
    classA := []string{
        "10.0.0.0",
        "8.8.8.8",
        "11.22.33.44",
        "15.16.17.18",
        "20.21.22.23",
        // Add more Class A addresses as needed
    }

    // Class B addresses
    classB := []string{
        "172.16.0.0",
        "191.168.0.1",
        "130.192.10.10",
        "172.17.18.19",
        "172.18.19.20",
        // Add more Class B addresses as needed
    }

    // Class C addresses
    classC := []string{
        "192.168.1.1",
        "203.0.113.1",
        "198.51.100.1",
        "192.168.2.2",
        "192.168.3.3",
        // Add more Class C addresses as needed
    }

    // Out of range
    outofrange := []string{
        "256.0.0.0",
        "0.256.0.0",
        "0.0.256.0",
        "0.0.0.256",
    }

    loopbackevenwrong := []string{
        "127.0.0.1",
        "127.0.0.3",
    }

    wronglength := []string{
        "192.168.1",
    }

    for _, addr := range wronglength {
        if CheckNetworkClass(addr) != "Length check failed" {
            t.Errorf("Address %s should have failed length check", addr)
        }
    }

    for _, addr := range loopbackevenwrong {
        if CheckNetworkClass(addr) != "Loopback: Only 127.0.0.1" {
            t.Errorf("Gives Loopback 127.0.0.1 even for %s", addr)
        }
    }

    for _, addr := range outofrange {
        if CheckNetworkClass(addr) != "Not Valid Network: out of range" {
            t.Errorf("Address %s should be out of range", addr)
        }
    }

    for _, addr := range classA {
        if CheckNetworkClass(addr) != "Class A" {
            t.Errorf("Address %s should be Class A", addr)
        }
    }

    for _, addr := range classB {
        if CheckNetworkClass(addr) != "Class B" {
            t.Errorf("Address %s should be Class B", addr)
        }
    }

    for _, addr := range classC {
        if CheckNetworkClass(addr) != "Class C" {
            t.Errorf("Address %s should be Class C", addr)
        }
    }
}
