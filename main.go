package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/manifoldco/promptui"
)

func main() {
	packageList := []string{"=Done="}
	installList := make([]string, 0)

	// Get shell script file
	err := filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".sh" && path != "install.sh" && path != "start.sh" {
			packageList = append(packageList, path[:len(path)-3])
		}
		return nil
	})

	// Prompt for what to install
	done := false
	for done == false {
		prompt := promptui.Select{
			Label: "Select packages to install (top-down order)",
			Items: packageList,
		}

		_, result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		if result == "=Done=" {
			done = true
		} else {
			installList = append(installList, result)
			for i, v := range packageList {
				if v == result {
					packageList = append(packageList[:i], packageList[i+1:]...)
				}
			}
		}
	}

	// Generate shell script
	script := ""
	for _, v := range installList {
		script += "sudo sh " + v + ".sh\n"
	}

	// Write to file
	if _, err := os.Stat("./install.sh"); err == nil {
		err = os.Remove("./install.sh")
		if err != nil {
			panic(err)
		}
	}
	err = ioutil.WriteFile("./install.sh", []byte(script), 0644)
	if err != nil {
		panic(err)
	}

	// Finish message
	fmt.Println("\nShell script 'install.sh' created.\nStart installation...\n")
}
