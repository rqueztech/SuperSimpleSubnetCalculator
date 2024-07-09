package network

import(
    "fmt"
    "strconv"
    "strings"
)

type NetworkDetails struct {
    networkgeneral map[string][]string
}

func CallThisFunction() {
    fmt.Println("This is the main project here")
}

func (n* NetworkDetails) SetNetworkGeneral() {
    n.networkgeneral = make(map[string][]string)

    n.networkgeneral["Class A"] = []string{"255.0.0.0", "Internal"}
    n.networkgeneral["Class B"] = []string{"255.255.0.0", "Internal"}
    n.networkgeneral["Class C"] = []string{"255.255.255.0", "Internal"}
}

func CheckNetworkClass(ipaddress string) string {
    byteblocks := strings.Split(ipaddress, ".")
    numberofbytes := len(byteblocks)

    var currentclass string

    for index, currentvalue := range(byteblocks) {
        currentvalue = strings.TrimSpace(currentvalue)
        currentnumber, err := strconv.Atoi(currentvalue)

        if err != nil {
            fmt.Println(err)
        }

        if numberofbytes < 4 {
            return "Length check failed"
        }

        if(index == 0) {
            if (currentnumber == 127) {
                currentclass = "Loopback: Only 127.0.0.1"
                break
            } else if(currentnumber >= 1 && currentnumber <= 126) {
                currentclass = "Class A"
            } else if(currentnumber >= 128 && currentnumber <= 191) {
                currentclass = "Class B"
            } else if(currentnumber >= 192 && currentnumber <= 223) {
                currentclass = "Class C"
            } else {
                currentclass = "Not Valid Network: out of range"
            }
        } else {
            if (currentnumber < 0 || currentnumber > 255) {
                currentclass = "Not Valid Network: out of range"
                break
            } 
        }
    }

    return currentclass
}
