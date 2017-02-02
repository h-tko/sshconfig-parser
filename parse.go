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
        }

    }

    if conf.Host != "" {
        result = append(result, conf)
    }

    return result, nil
}
