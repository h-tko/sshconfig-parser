package sshconfig

import (
    "testing"
)

func TestParse(t *testing.T) {
    result, err := Parse("./.sshconfig_test")

    if err != nil {
        t.Errorf("%v", err)
    }

    for _, data := range result {
        t.Logf("host=%s", data.Host)
        t.Logf("hostname=%s", data.HostName)
        t.Logf("port=%d", data.Port)
        t.Logf("user=%s", data.User)
        t.Log("=================")
    }
}
