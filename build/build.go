package main

import "hope/pkg/command"

func main() {
	path := "D:/work/code/go/hope"
	//cd ./apps/admin/internal/data && ent generate ./ent/schema
	runCommand := command.RunCommand(path, "cd", "-a")
	print(runCommand)
}
