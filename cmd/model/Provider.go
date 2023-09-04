package model

import (
	"RbsBurndownChart/cmd/config"
	"RbsBurndownChart/cmd/types"
	"go/types"
)

type Provider struct {
	tasks       taskReader
	projects    *DBProject
	devs        *DBDeveloper
	currentUser *Developer
}

func NewProvider(config *config.Config) *Provider {
	return nil
}

func (p *Provider) MakeBurndownChart(devLogin string) (*types.Burndownchart, error) {
	return nil, nil
}
func (p *Provider) readUser(devLogin string) (*types.Developer, error) {
	return nil, nil
}
func (p *Provider) readProject(projectId int64) (*types.Project, error) {
	return nil, nil
}
func (p *Provider) readTasks(projectId int64) ([]types.Task, error) {
	return nil, nil
}
func (p *Provider) costSum(tasks []types.Task) (int64, error) {
	return 0, nil
}
func (p *Provider) readProjectTeam(projectId int64) ([]types.Developer, error) {
	return nil, nil
}
func (p *Provider) readWorkingHours(devs []types.Developer) (map[string]int64, error) {
	return nil, nil
}

type taskReader interface {
	Read(p *Project) ([]types.Task, error)
}
