package trackusecase

import (
	"log"
	"os"
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
