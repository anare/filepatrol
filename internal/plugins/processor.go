package plugins

type Processor interface {
	Process(file FilePayload, poster Poster, processedDir string) error
}
