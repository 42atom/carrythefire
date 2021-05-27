package app

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func Start(src, dst string) error {
	//Pause or Warning check
	checkRes, err := pauseOrWarning(dst)
	if !checkRes {
		return err
	}
	if err != nil {
		log.Println(err)
	}

	//Check src is exists
	if _, err := os.Stat(src); os.IsNotExist(err) {
		return fmt.Errorf("source '%s' isn't exits, please check again", src)
	}
	//Check dst is exists,
	if _, err := os.Stat(dst); os.IsNotExist(err) {
		return fmt.Errorf("target '%s' isn't exits, please check again", dst)
	}
	srcPlotFiles, err := getPlotFileIn(src)
	if err != nil {
		return err
	}
	dstPlotFiles, err := getPlotFileIn(dst)
	if err != nil {
		return err
	}
	needCopyPlotFiles := map[string]int64{}
	for fname, fsize := range srcPlotFiles {
		dsize, ok := dstPlotFiles[fname]
		if ok {
			//Compare files size if exists
			if compareFileSize(fsize, dsize) {
				needCopyPlotFiles[fname] = fsize
			}
		} else {
			needCopyPlotFiles[fname] = fsize
		}
	}
	if len(needCopyPlotFiles) <= 0 {
		log.Println("No file need to be copy to dst")
		return nil
	}
	//Start copy files
	for fname, fsize := range needCopyPlotFiles {
		source := fmt.Sprintf("%s/%s", src, fname)
		target := fmt.Sprintf("%s/%s", dst, fname)
		log.Printf("Need file copy: %s", fname)

		tsize, err := cpFile(source, target)
		if err != nil {
			return err
		}
		if tsize != fsize {
			return fmt.Errorf("src %s(%d) is not equal dst %s(%d)", source, fsize, target, tsize)
		}
		log.Printf("Successfully copy %s(%d) to %s(%d)\n", source, fsize, target, tsize)
		err = delFile(source)
		if err != nil {
			return err
		}
		log.Printf("Delete source file: %s\n", source)
	}
	return nil
}

func compareFileSize(ssize, dsize int64) bool {
	return ssize != dsize
}

func getPlotFileIn(dir string) (map[string]int64, error) {
	res := map[string]int64{}
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		ext := filepath.Ext(path)
		if ext == ".plot" {
			res[info.Name()] = info.Size()
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func cpFile(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func delFile(f string) error {
	return os.Remove(f)
}
