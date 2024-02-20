# netadmin


- A network discovery module for identifying live hosts within a specified subnet.
- An OS detection module (not detailed previously, so I'll assume all targets are Ubuntu machines for simplicity).
- A user management module for creating users with sudo permissions on discovered Ubuntu hosts, using SSH for command execution.
- A simplified approach to configuration management for storing and retrieving SSH login credentials.

**Note:** The full implementation of each module, especially the network discovery and OS detection parts, involves complex logic and considerations not fully covered here. For the purpose of this example, I'll focus on providing a skeleton that outlines how these components might interact.

### Project Structure

```
network-discovery/
│
├── config/
│   └── creds.config        # Configuration file with SSH credentials
│
├── discovery/
│   └── discovery.go        # Network discovery module (simplified example)
│
├── usermanagement/
│   └── usermanagement.go   # User management module for creating users and granting sudo permissions
│
└── main.go                 # Main application entry point that ties everything together
```

### 1. `config/creds.config`

```
Username: admin
Password: adminpassword
```

### 2. `discovery/discovery.go`

This file would contain the simplified discovery logic. Please refer to the previous example provided for network discovery, with the understanding that it needs significant expansion to fully implement all functionalities, such as generating IPs and parsing ARP tables correctly.

### 3. `usermanagement/usermanagement.go`

Refer to the user management module example provided earlier. It demonstrates reading SSH credentials from a config file, connecting to a remote Ubuntu system over SSH, and executing commands to create a user with sudo permissions.

### 4. `main.go`

The main application file orchestrates the discovery of network hosts, identification of their operating systems (assuming Ubuntu for simplicity), and the creation of new users with sudo permissions on each identified host.

```go
package main

import (
    "fmt"
    "network-discovery/discovery"
    "network-discovery/usermanagement"
)

func main() {
    // Example subnet to scan - this should be provided or configured appropriately
    subnet := "10.2.2.0/24"

    // Perform network discovery to find live hosts
    hosts, err := discovery.DiscoverMachines(subnet)
    if err != nil {
        fmt.Printf("Error during discovery: %s\n", err)
        return
    }

    // Read SSH credentials from the config file
    creds, err := usermanagement.ReadCredentials("config/creds.config")
    if err != nil {
        fmt.Printf("Error reading credentials: %s\n", err)
        return
    }

    // Iterate over discovered hosts, assuming they are Ubuntu machines, and create a user with sudo permissions
    for _, host := range hosts {
        fmt.Printf("Processing host: %s\n", host)
        err := usermanagement.CreateUserWithSudo(host, "newuser", "newpassword123", creds)
        if err != nil {
            fmt.Printf("Error managing user on host %s: %s\n", host, err)
        } else {
            fmt.Printf("Successfully set up new user on host %s\n", host)
        }
    }
}
```
