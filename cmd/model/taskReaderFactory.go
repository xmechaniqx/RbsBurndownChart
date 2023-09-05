package model

import (
	"RbsBurndownChart/cmd/config"
	"RbsBurndownChart/cmd/types"
)

type taskReaderFactory struct {
	config *config.Config
}

func New(config *config.Config) *taskReaderFactory {
	return nil
}

func (tr *taskReaderFactory) Make(project *types.Project) taskReader {
	return nil
}
