package main

import (
    "fmt"
    "discovery"
    "usermanagement"
)

func main() {
    // Discover machines on the subnet
    machines, err := discovery.DiscoverMachines("192.168.1.0/24")
    if err != nil {
        fmt.Println("Error discovering machines:", err)
        return
    }

    // Detect OS and handle user creation
    for _, machine := range machines {
        /*osType, err := osdetection.DetectOS(machine)
        if err != nil {
            fmt.Println("Error detecting OS for machine:", machine, err)
            continue
        }*/

        // Example login details - use secure storage and handling for real credentials
        loginDetails := usermanagement.LoginDetails{
            Username: "admin",
            Password: "password",
        }

	osType = "Linux"
        // Attempt to create a user with sudo permissions
        if err := usermanagement.CreateUserWithSudo(machine, osType, loginDetails); err != nil {
            fmt.Println("Error creating user on machine:", machine, err)
        }
    }

    /*
    // Generate a network map
    if err := networkmap.GenerateMap(machines, "network.dot"); err != nil {
        fmt.Println("Error generating network map:", err)
    }*/
}

