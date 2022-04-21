package filesystem

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func CreateDefaultSubdirectories() error {
	subdirectories := []string{"projects", "packages", "builds", "tmp"}
	if err := CreateSubdirectories(subdirectories); err != nil {
		return err
	}
	return nil
}

func CreateSubdirectories(paths []string) error {
	for _, path := range paths {
		if err := CreateSubdirectory(path); err != nil {
			log.Print(err)
			return err
		}
	}
	return nil
}

func CreateSubdirectory(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0755)
		return nil
	} else if err == nil {
		return nil
	} else {
		log.Print(err)
		return err
	}
}

func CreateProjectSourceDirectory(name string) error {
	path := fmt.Sprintf("projects/%v", name)
	err := CreateSubdirectory(filepath.FromSlash(path))
	if err != nil {
		return err
	}
	return nil
}

func CreatePackageSourceFromBytes(name string, version int64, data []byte) error {
	path := fmt.Sprintf("packages/%v/%d/src", name, version)
	filepath.FromSlash(name)
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for n, byte := range data {
		if n%1024^2 == 0 {
			file.Sync()
		}
		err = writer.WriteByte(byte)
		if err != nil {
			return err
		}
		writer.Flush()
	}
	writer.Flush()
	file.Sync()

	return nil
}

func CreateProjectSourceFromBytes(name string, data []byte) error {
	path := fmt.Sprintf("projects/%v/src", name)
	filepath.FromSlash(path)
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	writer := bufio.NewWriter(file)
	for n, byte := range data {
		if n%1024^2 == 0 {
			file.Sync()
		}
		err = writer.WriteByte(byte)
		if err != nil {
			return err
		}
		writer.Flush()
	}
	writer.Flush()
	file.Sync()

	return err
}
