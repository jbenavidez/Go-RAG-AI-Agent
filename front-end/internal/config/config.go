package config

import (
	pb "frontend/proto/generated"
	"html/template"
)

type AppConfig struct {
	GRPCClient    pb.AIAgentServiceClient
	UseCache      bool
	TemplateCache map[string]*template.Template
}
