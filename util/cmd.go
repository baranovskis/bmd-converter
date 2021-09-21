package util

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/baranovskis/bmd-converter/bmd"
	"github.com/urfave/cli"
)

func CommandAction(c *cli.Context, ba bmd.BaseBmdAdapter) error {
	fileName := c.Args().Get(0)
	b := bmd.NewBmdInfo(ba)
	return doWork(b, fileName)
}

func doWork(b *bmd.BaseBmdInfo, fileName string) error {
	inputExt := filepath.Ext(fileName)

	outputName := fileName
	if inputExt == ".bmd" {
		outputName = strings.Replace(outputName, ".bmd", ".csv", -1)
		fmt.Println("Decrypting:", fileName, "->", outputName)
	} else if inputExt == ".csv" {
		outputName = strings.Replace(outputName, ".csv", ".bmd", -1)
		fmt.Println("Encrypting:", fileName, "->", outputName)
	} else {
		return errors.New("unknown file extension")
	}

	inputFile, err := os.Open(fileName)
	if err != nil {
		return err
	}

	defer inputFile.Close()

	outputFile, err := os.OpenFile(outputName, os.O_TRUNC|os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	fileInfo, err := inputFile.Stat()
	if err != nil {
		return err
	}

	fileSize := fileInfo.Size()
	data := make([]byte, fileSize)

	bytesRead, err := inputFile.Read(data)
	if err != nil {
		return err
	}

	fmt.Println("bytes read =", bytesRead)

	if inputExt == ".bmd" {
		return b.Decrypt(data, outputFile)
	}

	return b.Encrypt(data, outputFile)
}
