package main

import (
    "fmt"
    "strings"
    "strconv"
    "bufio"
    "os"
)

type NetworkDetails struct {
    networksubnet map[string][]string
}

func (n* NetworkDetails) setnetworksubnet() {
    n.networksubnet = make(map[string][]string)
    
    n.networksubnet["Class A"] = []string{"255.0.0.0", "Internal"}
    n.networksubnet["Class B"] = []string{"255.255.0.0", "Internal"}
    n.networksubnet["Class C"] = []string{"255.255.255.0", "External"}
}:

func checknetworkclass(ipaddress string) (string, int) {
    byteblocks := strings.Split(ipaddress, ".")
    numberofbytes := len(byteblocks)

    fmt.Println(len(byteblocks))
    hostaddress := -1

    var currentclass string

    for index, currentvalue := range(byteblocks) {
        currentnumber, err := strconv.Atoi(currentvalue)

        if err != nil {
            fmt.Println(err)
            currentclass = "Error in the address"
        }

        if numberofbytes < 4 {
            fmt.Println("Length check failed")
            break
        }

        if(index == 0) {
            if (currentnumber == 127) {
                currentclass = "Local Host Range"
            } else if(currentnumber >= 1 && currentnumber <= 126) {
                currentclass = "Class A"
            } else if(currentnumber >= 128 && currentnumber <= 191) {
                currentclass = "Class B"
            } else if(currentnumber >= 192 && currentnumber <= 223) {
                currentclass = "Class C"
            } else {
                currentclass = "Base number doesnt work"
            }
        } else {
            if (currentnumber < 0 || currentnumber > 255) {
                currentclass = "Not Valid Network: out of range"
                break
            } 
        }
    }

    return currentclass, hostaddress
}


func main() {
    scanner := bufio.NewScanner(os.Stdin)


    info := NetworkDetails{}
    info.setnetworksubnet()

    for {
        baseCIDR := 24
        fmt.Print("Enter Network: ")
        scanner.Scan()
        ipaddress := scanner.Text()
        
        class, hostaddress := checknetworkclass(ipaddress)
        fmt.Println(hostaddress)
        fmt.Println(class)

        defaultsubnet := info.networksubnet[class][0]
        networktype := info.networksubnet[class][1]

        fmt.Println(defaultsubnet)
        fmt.Println(networktype)

        if class == "Class C" {
            fmt.Println("How many networks do you want?")

            scanner.Scan()
            numberofnetworks, err := strconv.Atoi(scanner.Text())

            if err != nil {
                fmt.Println("")
            }

            var bitsneeded int

            for bitsneeded = 0; bitsneeded <= 8; bitsneeded++ {
                if (1 << bitsneeded) >= numberofnetworks {
                    break
                }
            }

            fmt.Printf("bits needed: %d", bitsneeded)
            fmt.Printf("base cidr: %d", baseCIDR)
        }

    }
}
