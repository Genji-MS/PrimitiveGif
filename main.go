package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/png"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

var PATH string = "tempImage/"

func main() {
	dir := flag.String("dir", "", " directory where we parse .png files")
	flag.Parse()

	var wg sync.WaitGroup 
	if *dir != "" {
		// Check in directory, if files have proper extension convert them
		files, err := ioutil.ReadDir(*dir)
		check(err)
		printedFiles := 0
		totalFileSize := 0.0
		for _, file := range files {
			if filepath.Ext(file.Name()) == ".png" {
				var fileSequence [8]string
				printedFiles++
				fileName := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name())) //split file extension
				fmt.Println("\033[33mImage series " + fileName + " begin:\033[0m")

				for reps := 0; reps < 8; reps++ {
					sequenceName := fileName + "_" + fmt.Sprint(reps) + ".png"
					fileSequence[reps] = sequenceName
					inputString := "-i=" + fileName + ".png"
					outputString := "-o=" + PATH + sequenceName
					shapesString := "-n=" + "50"
					imageDimension := "-s=" + "256" 
					wg.Add(1)
					fmt.Println("Processing image " + sequenceName)
					go imgCreate(inputString, outputString, shapesString, imageDimension, &wg)
				}
				wg.Wait()

				fmt.Println("\033[32mImage series " + fileName + " complete.\033[0m")
				totalFileSize += gifCreate(fileSequence)
			}
		}
		if printedFiles > 0 {
			FileSize_String := fmt.Sprintf("%.1f", totalFileSize)
			// Terminal commands for color/font https://stackoverflow.com/questions/2924697/how-does-one-output-bold-text-in-bash
			// Black: 30m. Red: 31m. Green: 32m. Yellow: 33m. Blue: 34m. Magenta: 35m. Cyan: 36m. White: 37m. Reset: 0m
			fmt.Printf("\033[32m\033[1mSuccess!\033[0m Generated \033[1m%s\033[0m image \033[34m(\033[36m%skB total\033[34m)\033[0m.\n", strconv.Itoa(printedFiles), FileSize_String)
		} else {
			fmt.Println("\033[31mERROR\033[0m No \033[36m(.png)\033[0m files found in directory")
		}
	}
}

func imgCreate(inputString, outputString, shapesString, imageDimension string, wg *sync.WaitGroup) {
	defer wg.Done()
	cmnd := exec.Command("primitive", inputString, outputString, shapesString, imageDimension)
	cmnd.Run()
	fmt.Println("Image " + outputString[3:] + " Created")
}

func gifCreate(fileSequence [8]string) (fileSize float64) {
	var gifSequence [8]string
	animatedGif := &gif.GIF{}
	gifFileName := strings.Split(fileSequence[0], "_")[0] + ".gif"
	for n, file := range fileSequence {
		gifFile := strings.TrimSuffix((file), filepath.Ext(file)) + ".gif"
		gifSequence[n] = gifFile
		outFile := openOrCreate(PATH + gifFile)
		//convert image to gif and save in folder
		f, err := os.Open(PATH + file)
		check(err)
		imageData, err := png.Decode(f)
		check(err)
		f.Close()
		err = gif.Encode(outFile, imageData, nil)
		check(err)
		outFile.Close()
		//open gif file and add to animated gif
		f, _ = os.Open(PATH + gifFile)
		gifData, _ := gif.Decode(f)
		f.Close()
		animatedGif.Image = append(animatedGif.Image, gifData.(*image.Paletted))
		animatedGif.Delay = append(animatedGif.Delay, 13)
		err = os.Remove(PATH + file)
		check(err)
		err = os.Remove(PATH + gifFile)
		check(err)
	}
	f, _ := os.OpenFile(gifFileName, os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	gif.EncodeAll(f, animatedGif)

	fileInfo, _ := os.Stat(gifFileName)
	fileSize = float64(fileInfo.Size() / 1000)
	return fileSize
}

func openOrCreate(filename string) *os.File {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		file, err := os.Create(filename)
		check(err)
		return file
	} else {
		// trying to write to an existing file gives us errors, delete it instead and create
		err = os.Remove(filename)
		check(err)
		file, err := os.Create(filename)
		check(err)
		return file
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
