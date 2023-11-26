package trackusecase

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

type TrackUsecase struct {
	pathToDir string
}

func New(path string) TrackUsecase {
	return TrackUsecase{pathToDir: path}
}

func (tu TrackUsecase) GetTracksName() []string {
	files, err := os.ReadDir(tu.pathToDir)
	if err != nil {
		log.Printf("failed to read directory %s, err: %s", tu.pathToDir, err.Error())
		return []string{}
	}

	tracks := make([]string, 0, len(files))
	for _, file := range files {
		tracks = append(tracks, file.Name())
	}
	return tracks
}

func (tu TrackUsecase) UploadTrack(filename string, fs io.Reader) error {
	filePath := filepath.Join(tu.pathToDir, filename)

	tempFile, err := os.Create(filePath)
	fmt.Println(filePath)
	if err != nil {
		return fmt.Errorf("failed to create temp file: %s", err.Error())
	}
	defer tempFile.Close()

	if _, err := io.Copy(tempFile, fs); err != nil {
		return fmt.Errorf("failed to save file on server")
	}

	return nil
}
