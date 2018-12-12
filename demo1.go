package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var (
		cmd *exec.Cmd
	)
	cmd=exec.Command("/bin/bash","-c","echo1;echo2")
	err:=cmd.Run()
	fmt.Println(err)
	//if err!=nil{
	//	fmt.Println(err)
	//}
}












