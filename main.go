package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/h2non/bimg"
)

func main() {
	cwd, err := os.Getwd()
	name := flag.String("named", "nil", "select only files with certain name")
	rename := flag.String("rename", "nil", "rename all selected files with specified name")
	flag.Parse()
	reader := bufio.NewReader(os.Stdin)
	if err != nil {
		log.Fatalln("Something unexpected happened at startup. There is no current working directory.")
	}
	log.Println("| png2jpg -> " + cwd)
	if *name != "nil" && *rename != "nil" {
		log.Println("| png2jpg ->", *name, "->", *rename)
	} else if *name != "nil" && *rename == "nil" {
		log.Println("| png2jpg ->", *name, "->", *name+".jpg")
	} else if *name == "nil" && *rename != "nil" {
		log.Println("| png2jpg -> *", "->", *rename)
	}
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
		if info.Name() != *name && *name != "nil" {
			return nil
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
				if *rename != "nil" {
					bimg.Write(strings.TrimSuffix(path, info.Name())+*rename, i)
				} else {
					bimg.Write(strings.TrimSuffix(path, ".png")+".jpg", i)
				}
				_ = os.Remove(path)
				log.Println("| Converted", path, " to jpg.")
			}
		}
		return nil
	})
	if err != nil {
		log.Fatalln(err)
	}
}
