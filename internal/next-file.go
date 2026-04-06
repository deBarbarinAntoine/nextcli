package internal

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strings"
	"text/template"
	
	"nextcli/internal/enum"
)

func NewNextFile(name string, fileType enum.FileType) *NextFile {
	nextFile := &NextFile{
		Name: name,
		Type: fileType,
	}
	
	nextFile.initialize()
	
	fmt.Printf(":: nextcli: [INFO] generating file %s\n", nextFile.Path)
	
	return nextFile
}

func (n *NextFile) initialize() {
	
	if n.Type == enum.FileTypes.Page {
		n.Name = strings.Trim(n.Name, "/")
		n.Path = filepath.Join(AppFolder, filepath.Clean(fmt.Sprintf("%s/%s%s", n.Name, n.Type, PageExtension)))
		n.Name = PageFuncName
		return
	}
	
	var extension string
	switch n.Type {
	case enum.FileTypes.Atom, enum.FileTypes.Molecule, enum.FileTypes.Organism, enum.FileTypes.Template:
		extension = ComponentExtension
	case enum.FileTypes.Model, enum.FileTypes.Provider, enum.FileTypes.Service, enum.FileTypes.Repository, enum.FileTypes.Util:
		n.IsLib = true
		extension = LibExtension
	default:
		extension = ComponentExtension
	}
	
	n.Path = filepath.Join(n.Type.String(), fmt.Sprintf("%s%s", n.Name, extension))
}

func (n *NextFile) getContent() ([]byte, error) {
	if n.IsLib {
		return nil, nil
	}
	
	tmpl := template.New(componentTemplateFile)
	
	_, err := tmpl.ParseFS(Templates, componentTemplateFile)
	if err != nil {
		return nil, fmt.Errorf(":: nextcli: [ERROR] failed to parse component template: %w", err)
	}
	
	var buffer bytes.Buffer
	
	err = tmpl.ExecuteTemplate(&buffer, componentTemplateName, n)
	if err != nil {
		return nil, fmt.Errorf(":: nextcli: [ERROR] failed to execute component template: %w", err)
	}
	
	return buffer.Bytes(), nil
}

func (n *NextFile) Generate() error {
	err := CreateFolders(n.Path)
	if err != nil {
		return err
	}
	
	content, err := n.getContent()
	if err != nil {
		return err
	}
	
	return CreateFile(n.Path, content)
}
