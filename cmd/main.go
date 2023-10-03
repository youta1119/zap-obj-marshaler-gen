package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/youta1119/zap-obj-marshaler-gen/internal"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	path := flag.String("f", "", "target file path")
	flag.Parse()

	data, err := internal.ParseAndGenerate(*path)
	if err != nil {
		return fmt.Errorf("internal.ParseAndGenerate: %w", err)
	}
	genFilePath := strings.Replace(*path, ".go", "_zap_obj_marshaler.go", 1)

	_, err = os.Stat(genFilePath)
	if !os.IsNotExist(err) {
		err = os.Remove(genFilePath)
		if err != nil {
			return fmt.Errorf("os.Remove: %w", err)
		}
	}
	file, err := os.Create(genFilePath)
	if err != nil {
		return fmt.Errorf("os.Create: %w", err)
	}
	defer func() {
		_ = file.Close()
	}()
	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("file.Write: %w", err)
	}
	fmt.Printf("generate file. %s\n", genFilePath)
	if err != nil {
		return fmt.Errorf("file.Write: %w", err)
	}
	return nil
}
