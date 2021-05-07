package main

import "sinx/snet"

func main() {
	s:=snet.NewServer("[sinx V0.1]")

	s.Serve()
}
