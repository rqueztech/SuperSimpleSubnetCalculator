package main

import (
    "fmt"
    "SuperSimpleSubnetCalculator/internal/network"
    "bufio"
    "os"
)

func main() {
    fmt.Println("Enter the IP address to check the class: ")
    reader := bufio.NewReader(os.Stdin)
    ipaddress, _ := reader.ReadString('\n')
    networkclass := network.CheckNetworkClass(ipaddress)
    fmt.Println(networkclass)
}
