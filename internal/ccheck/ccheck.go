package ccheck

import (
	"fmt"
	"log"

	"github.com/spf13/afero"
)

type Application struct {
	afs        afero.Afero
	entrypoint string
}

func NewApplication(afs afero.Afero, entrypoint string) *Application {
	return &Application{
		afs:        afs,
		entrypoint: entrypoint,
	}
}

func (application Application) Execute() (int, error) {
	service := &Service{
		afs: application.afs,
	}

	files, err := service.GetFiles()
	if err != nil {
		log.Panic(err)
	}

	for _, file := range files {
		fmt.Println(file)
	}

	return 0, nil
}
