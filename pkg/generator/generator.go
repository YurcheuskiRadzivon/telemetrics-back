package generator

import "github.com/google/uuid"

type UUVGenerator interface {
	NewSessionID() string
}

type Generator struct{}

func (g *Generator) NewSessionID() string {
	return uuid.New().String()
}

func (g *Generator) NewManageSessionID() string {
	return uuid.New().String()
}
