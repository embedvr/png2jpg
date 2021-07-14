package main

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/h2non/bimg"
)

func main() {
	cwd, err := os.Getwd()
	reader := bufio.NewReader(os.Stdin)
	if err != nil {
		log.Fatalln("Something unexpected happened at startup. There is no current working directory.")
	}
	log.Println("| png2jpg -> " + cwd)
	log.Println("| THIS WILL CONVERT *ALL* PNG FILES IN *ALL* DIRECTORIES UNDER THE DIRECTORY ABOVE. ARE YOU SURE YOU WANT TO CONTINUE?")
	log.Println("| y / N -> ")
	c, _, err := reader.ReadRune()
	if err != nil {
		log.Fatalln(err)
	}
	switch strings.ToLower(string(c)) {
	case "n":
		os.Exit(0)
		break
	case "y":
		log.Println("| OK.")
	default:
		os.Exit(0)
		break
	}
	err = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		//log.Println("| " + path)
		if strings.HasSuffix(info.Name(), ".png") {
			i, err := bimg.Read(path)
			if err != nil {
				log.Fatalln("| " + err.Error())
			}
			i, err = bimg.NewImage(i).Convert(bimg.JPEG)
			if err != nil {
				log.Fatalln("| " + err.Error())
			}

			if bimg.NewImage(i).Type() == "jpeg" {
				bimg.Write(path, i)
				log.Println("| Converted", path, " to jpg.")
			}
		}
		return nil
	})
	if err != nil {
		log.Fatalln(err)
	}
}
