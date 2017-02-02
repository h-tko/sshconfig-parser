package sshconfig

import (
    "bufio"
    "os"
    "strings"
    "strconv"
)

type SSHConfig struct {
    Host     string
    HostName string
    Port     int
    User     string
    IdentityFile string
}

func Parse(fullfilename string) ([]*SSHConfig, error) {
    return config(fullfilename)
}

func config(fullfilename string) ([]*SSHConfig, error) {
    file, err := os.Open(fullfilename)

    if err != nil {
        return nil, err
    }

    defer file.Close()

    var result []*SSHConfig

    reader := bufio.NewReaderSize(file, 4096)
    conf := new(SSHConfig)

    for line := ""; err == nil; line, err = reader.ReadString('\n') {
        line = strings.TrimSpace(line)

        data := strings.Split(line, "#")[0]
        data = strings.TrimSpace(data)

        if len(data) < 1 {
            continue
        }

        dataset := strings.Split(data, " ")
        dataset = extractHasData(dataset)
        key := strings.TrimSpace(dataset[0])
        key = strings.ToLower(key)

        switch key {
        case "host":
            if conf.Host != "" {

                if conf.Port == 0 {
                    conf.Port = 22
                }

                result = append(result, conf)
                conf = new(SSHConfig)
            }

            conf.Host = dataset[1]

        case "hostname":
            conf.HostName = dataset[1]

        case "port":
            port, err := strconv.Atoi(dataset[1])

            if err != nil {
                conf.Port = 22
            } else {
                conf.Port = port
            }

        case "user":
            conf.User = dataset[1]
        case "identityfile":
            conf.IdentityFile = dataset[1]
        }

    }

    if conf.Host != "" {
        result = append(result, conf)
    }

    return result, nil
}

func extractHasData(dataset []string) []string {
    var result []string

    for _, data := range dataset {
        trimed := strings.TrimSpace(data)

        if len(data) > 0 {
            result = append(result, trimed)
        }
    }

    if len(result) < 2 {
        result = append(result, "")
    }

    return result
}
