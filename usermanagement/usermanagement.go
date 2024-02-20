package usermanagement

import (
    "bufio"
    "fmt"
    "golang.org/x/crypto/ssh"
    "os"
    "strings"
)

// Credentials holds the SSH login credentials.
type Credentials struct {
    Username string
    Password string
}

// ReadCredentials reads the SSH login credentials from a config file.
func ReadCredentials(configPath string) (*Credentials, error) {
    file, err := os.Open(configPath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    creds := &Credentials{}
    for scanner.Scan() {
        line := scanner.Text()
        if strings.HasPrefix(line, "Username:") {
            creds.Username = strings.TrimSpace(line[len("Username:"):])
        } else if strings.HasPrefix(line, "Password:") {
            creds.Password = strings.TrimSpace(line[len("Password:"):])
        }
    }

    if err := scanner.Err(); err != nil {
        return nil, err
    }

    return creds, nil
}

// CreateUserWithSudo connects to the given host as the specified user and creates a new user with sudo permissions.
func CreateUserWithSudo(host, newUser, newUserPassword string, creds *Credentials) error {
    config := &ssh.ClientConfig{
        User: creds.Username,
        Auth: []ssh.AuthMethod{
            ssh.Password(creds.Password),
        },
        HostKeyCallback: ssh.InsecureIgnoreHostKey(), // NOTE: In production, replace with a more secure option.
    }

    // Connect to the host
    connection, err := ssh.Dial("tcp", fmt.Sprintf("%s:22", host), config)
    if err != nil {
        return fmt.Errorf("failed to dial: %s", err)
    }
    defer connection.Close()

    // Execute commands to create a new user and add to sudo group
    session, err := connection.NewSession()
    if err != nil {
        return fmt.Errorf("failed to create session: %s", err)
    }
    defer session.Close()

    createUserCmd := fmt.Sprintf("sudo useradd -m -s /bin/bash %s && echo '%s:%s' | sudo chpasswd", newUser, newUser, newUserPassword)
    if err := session.Run(createUserCmd); err != nil {
        return fmt.Errorf("failed to run command: %s", err)
    }

    grantSudoCmd := fmt.Sprintf("echo '%s ALL=(ALL) NOPASSWD:ALL' | sudo tee /etc/sudoers.d/%s", newUser, newUser)
    session, err = connection.NewSession()
    if err != nil {
        return fmt.Errorf("failed to create session: %s", err)
    }
    defer session.Close()

    if err := session.Run(grantSudoCmd); err != nil {
        return fmt.Errorf("failed to run command: %s", err)
    }

    return nil
}

