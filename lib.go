package lib

import "fmt"

var APP = Info{
	Name:    "lib",
	Version: "0.0.1_dev",
	Author:  "lyonmu@foxmail.com",
}

type Info struct {
	Name    string
	Version string
	Author  string
}

func (i Info) String() string {
	return fmt.Sprintln("Name:", i.Name, "Version:", i.Version, "Author:", i.Author)
}
