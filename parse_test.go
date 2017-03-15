package sshconfig

import (
    "testing"
    . "github.com/smartystreets/goconvey/convey"
)

func TestParseFromConfig(t *testing.T) {

    Convey("ただしいconfigファイルを読み込ませるとスライスに格納される", t, func() {

        assumptions := []SSHConfig{
            {
                Host: "st2",
                HostName: "st2.dev",
                User: "development",
                Port: 22,
            },
            {
                Host: "test1",
                HostName: "192.168.0.112",
                User: "development",
                Port: 19122,
            },
            {
                Host: "test3",
                HostName: "test.serve",
                User: "user",
                IdentityFile: "~/Downloads/test.pem",
            },
        }

        result, err := Parse("./.sshconfig_test")

        So(err, ShouldEqual, nil)

        for index, data := range result {
            So(data.Host, ShouldEqual, assumptions[index].Host)
            So(data.HostName, ShouldEqual, assumptions[index].HostName)
            So(data.Port, ShouldEqual, assumptions[index].Port)
            So(data.User, ShouldEqual, assumptions[index].User)
            So(data.IdentityFile, ShouldEqual, assumptions[index].IdentityFile)
        }
    })
}
