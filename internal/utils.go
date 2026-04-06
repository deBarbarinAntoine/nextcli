package internal

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func CreateFolders(path string) error {
	
	path = filepath.Clean(path)
	path = filepath.Dir(path)
	
	if path == "." {
		return nil
	}
	
	err := os.MkdirAll(path, 0755)
	if err != nil && !os.IsExist(err) {
		return fmt.Errorf(":: nextcli: [ERROR] failed to create folders: %w", err)
	}
	
	return nil
}

func CreateFile(path string, content []byte) error {
	
	path = filepath.Clean(path)
	
	if path == "." {
		return nil
	}
	
	if content == nil {
		file, err := os.Create(path)
		if err != nil {
			return fmt.Errorf(":: nextcli: [ERROR] failed to create file: %w", err)
		}
		
		err = file.Close()
		if err != nil {
			return fmt.Errorf(":: nextcli: [ERROR] failed to close file: %w", err)
		}
		
		return nil
	}
	
	err := os.WriteFile(path, content, 0644)
	if err != nil {
		return fmt.Errorf(":: nextcli: [ERROR] failed to write file: %w", err)
	}
	
	return nil
}

func GetUserInput(prompt string) (string, error) {
	fmt.Println(prompt)
	fmt.Print("> ")
	reader := bufio.NewReader(os.Stdin)
	urlStr, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf(":: nextcli [ERROR] failed to read user input: %w", err)
	}
	return strings.TrimSpace(urlStr), nil
}
