package main

import (
  "fmt"
  "os"
  "os/exec"
  "log"
)

func main() {
	arg := "Default"
	if(len(os.Args) > 1){
		arg = os.Args[1]
	}
	// fmt.Println(arg)
	cmdGit := "git"
	//git add .
	cmdArgs := []string{"add", "."}
	tryExec(cmdGit, cmdArgs)

	//git commit -m "my commit message || Default"
	cmdArgs = []string{"commit", "-m", arg}
	tryExec(cmdGit, cmdArgs)

	//git push
	cmdArgs = []string{"push"}
	tryExec(cmdGit, cmdArgs)
}

func tryExec(cmd string, cmdArgs []string){
	out, err := exec.Command(cmd, cmdArgs...).Output()
	if(err != nil){
		log.Fatal(err)
	}
	sha := string(out)
	fmt.Println(sha)
}

