package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/remote"
	"github.com/google/go-containerregistry/pkg/v1/tarball"
)

func main() {
	ref, err := name.ParseReference("container-infra-local.hzisoj70.china.nsn-net.net/pause:3.3", name.StrictValidation)
	if err != nil {
		fmt.Println(err.Error())
	}

	/*
		img, err := remote.Image(ref, remote.WithAuthFromKeychain(authn.DefaultKeychain))
		if err != nil {
			fmt.Println("Error: ", err.Error())
		}

		cfgHash, err := img.ConfigName()
		if err != nil {
			fmt.Println("Error: ", err.Error())
		}

		fmt.Println(cfgHash)
	*/

	pausePath := filepath.Join(os.Getenv("HOME"), "pause:latest.tar")
	//pauseTag, err := name.NewTag("pause:demo")
	//if err != nil {
	//	fmt.Println("Error: ", err.Error())
	//}

	localImg, err := tarball.ImageFromPath(pausePath, nil)
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}

	options := make([]remote.Option, 0)
	options = append(options, remote.WithAuthFromKeychain(authn.DefaultKeychain))

	err = remote.Write(ref, localImg, options...)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	img, err := remote.Image(ref, remote.WithAuthFromKeychain(authn.DefaultKeychain))
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}

	cfgHash, err := img.ConfigName()
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}

	fmt.Println(cfgHash)
}
