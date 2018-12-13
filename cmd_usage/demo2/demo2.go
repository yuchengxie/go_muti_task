package demo2

import (
	"fmt"
	"os/exec"
)

func main() {
	var(
		cmd *exec.Cmd
		output []byte
		err error
	)
	cmd=exec.Command("/bin/bash","-c","sleep 5;ls -l;sleep 1;echo hello")

	if output,err=cmd.CombinedOutput();err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Println(string(output))

}


















