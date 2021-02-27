package main

import (
	"fmt"
	"testing"
	"os"
	"os/exec"
)

func TestVerifyPrimitiveAccess(t *testing.T){
	cmnd := exec.Command("primitive", "-i=./SampleImage/testimg.png", "-o=./SampleImage/testcomplete.png", "-n=1", "-s=28")
	cmnd.Run()

	if _, err := os.Stat("./SampleImage/testcomplete.png"); err == nil {
		os.Remove("./SampleImage/testcomplete.png")
		fmt.Println("Test complete, Primitive has been installed successfuly")
	} else {
		t.Error(" ERROR !!! Primitive is not found.\nInstall Primitive with 'go get -u github.com/fogleman/primitive'\n Additionaly try '$export PATH=$PATH:$(go env GOPATH)/bin'\n  retest to verify")
	}
}