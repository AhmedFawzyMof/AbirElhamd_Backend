package main

import (
	"abir-el-hamd/internal/server"
)

func main() {
	server := server.NewServer(":3000", ".")
	server.Start()
}
