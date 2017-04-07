package main

import (
  "fmt"
  "os"
  "os/exec"
  "log"
)

func main() {
	// Setting up the default commit message.
	// @todo Maintain a list of previous messages to cycle through. Maybe.
	arg := "Default"
	// Do we have any command line arguments?
	if len(os.Args) > 1 {
		// Set the commit message from the first argument passed.
		arg = os.Args[1]
	}

	// The name of our git binary, I don't have the full path to git in, assuming git is available
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

// Trying out being able to exec things. Just die if it fails.
func tryExec(cmd string, cmdArgs []string){
	out, err := exec.Command(cmd, cmdArgs...).Output()
	if err != nil {
		log.Fatal(err)
	}
	sha := string(out)
	fmt.Println(sha)
}

