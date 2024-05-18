package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)
var (
	moduleName string
	domainName string
	domainNamePascalCase string
)


func main() {

	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go helper make:domain <domain-name>")
		os.Exit(1)
	}
	// panic("here")
	command := os.Args[1]
	subcommand := os.Args[2]

	if command == "helper" {
		switch subcommand {
		case "make:domain":
			if len(os.Args) < 4 {
				fmt.Println("Usage: go run main.go helper make:domain <domain-name>")
				os.Exit(1)
			}
			domainName = os.Args[3]
			setModuleName()

			createDomainStructure()
		default:
			fmt.Printf("Unknown subcommand: %s\n", subcommand)
			os.Exit(1)
		}
	} else {
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}
}


func setModuleName()  {
	moduleName = ""
	modFilePath := filepath.Join("go.mod")
	content, err := os.ReadFile(modFilePath)
	if err != nil {
		fmt.Printf("Error reading go.mod file: %v\n", err)
		os.Exit(1)
	}
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "module ") {
			moduleName =  strings.TrimPrefix(line, "module ")
		}
	}
	domainNamePascalCase = ToPascalCase(domainName)
}

func createDomainStructure() {
	basePath := filepath.Join("internal", "domain", domainName)
	subDirs := []string{"dependency","handler", "model", "repository", "seeder","service", "validators"}
	// Create the base directory
	err := os.MkdirAll(basePath, os.ModePerm)
	if err != nil {
		fmt.Printf("Error creating domain directory: %v\n", err)
		os.Exit(1)
	}
	
	// Create the subdirectories
	for _, dir := range subDirs {
		path := filepath.Join(basePath, dir)
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			fmt.Printf("Error creating subdirectory %s: %v\n", dir, err)
			os.Exit(1)
		}
	}
	templates := getTemplates()


	// Create the Go files with boilerplate content
	createGoFiles(basePath,templates)



	fmt.Printf("Domain structure for %s created successfully.\n", domainName)
}

func createGoFiles(basePath string,templates template) {
	// Define the boilerplate content for each file
	fileContents := map[string]string{
		
		fmt.Sprintf("handler/%s_handler.go", domainName) :      templates.handlerTemplate,
		fmt.Sprintf("model/%s.go", domainName):					templates.modelTemplate,
		fmt.Sprintf("repository/%s_repository.go", domainName):  templates.repositoryTemplate,
		fmt.Sprintf("service/%s.service.go", domainName):  templates.serviceTemplate,
		"validators/validators.go":  templates.validatorsTemplate,
		fmt.Sprintf("dependency/%s_dependency.go", domainName):  templates.dependencyTemplate,
	}

	// Create the Go files
	for file, content := range fileContents {
		path := filepath.Join(basePath, file)
		err := os.WriteFile(path, []byte(content), 0644)
		if err != nil {
			fmt.Printf("Error creating file %s: %v\n", file, err)
			os.Exit(1)
		}
	}
}


func ToPascalCase(str string) string {
    if len(str) == 0 {
        return ""
    }
    return strings.ToUpper(string(str[0])) + strings.ToLower(str[1:])
}


type template struct {
	handlerTemplate string
	modelTemplate string
	repositoryTemplate string
	serviceTemplate string
	validatorsTemplate string
	dependencyTemplate string
}

func getTemplates() template {


	handlerTemplate := (
`package handler

import (
"`+moduleName+`/internal/domain/`+domainName + `/service"

)

// `+domainName+`Handler handles HTTP requests related to `+domainName+`.
type `+ domainNamePascalCase+`Handler struct {
service  *service.`+domainNamePascalCase+`Service
}

// NewsHandler creates a new `+ domainNamePascalCase+`Handler instance.
func New`+domainNamePascalCase+`Handler(service *service.`+domainNamePascalCase+`Service) *`+domainNamePascalCase+`Handler {
return &`+domainNamePascalCase+`Handler{service: service}
}
// Add your handler methods here
`)



	modelTemplate := (`package model

// `+domainNamePascalCase+` represents the domain model for `+ domainName+`.
type `+domainNamePascalCase+` struct {
// Define your model fields here
}
`)

	repositoryTemplate := (`package repository
import (
	"gorm.io/gorm"

)

// `+domainNamePascalCase+`Repository handles database operations for the %s domain.
type `+domainNamePascalCase+`Repository struct {
	db *gorm.DB
}

// New`+domainNamePascalCase+`Repository creates a new Repository instance.
func New`+domainNamePascalCase+`Repository(db *gorm.DB) *`+domainNamePascalCase+`Repository {
return &`+domainNamePascalCase+`Repository{db: db}
}

// Add your repository methods here
`)

	serviceTemplate := `package service

import (
	"`+moduleName+`/internal/domain/`+domainName+`/repository"
)

// Service provides business logic for the `+domainName+` domain.
type `+domainNamePascalCase+`Service struct {
	repo *repository.`+domainNamePascalCase+`Repository
}

// NewService creates a new Service instance.
func New`+domainNamePascalCase+`Service(repo *repository.`+domainNamePascalCase+`Repository) *`+domainNamePascalCase+`Service {
	return &`+domainNamePascalCase+`Service{repo: repo}
}

// Add your service methods here
`

validatorsTemplate := `package validators

// Add your validator functions here
`	


dependencyTemplate := (`package dependency
import (
	"`+moduleName+`/internal/domain/`+domainName+`/handler"
	"`+moduleName+`/internal/domain/`+domainName+`/repository"
	"`+moduleName+`/internal/domain/`+domainName+`/service"

	"gorm.io/gorm"
)

func Setup`+domainNamePascalCase+`Dependencies(db *gorm.DB) (*handler.`+domainNamePascalCase+`Handler, error) {
	`+domainName+`Repo := repository.New`+domainNamePascalCase+`Repository(db)

	`+domainName+`Service := service.New`+domainNamePascalCase+`Service(`+domainName+`Repo)

	`+domainName+`Handler := handler.New`+domainNamePascalCase+`Handler(`+domainName+`Service)

	return `+domainName+`Handler, nil
}
`)
		
	return template{
		handlerTemplate: handlerTemplate,
		modelTemplate: modelTemplate,
		repositoryTemplate: repositoryTemplate,
		serviceTemplate: serviceTemplate,
		validatorsTemplate: validatorsTemplate,
		dependencyTemplate: dependencyTemplate,
	}
}


