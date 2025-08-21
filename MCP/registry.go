package main

import (
	"github.com/input-api/mcp-server/config"
	"github.com/input-api/mcp-server/models"
	tools_content "github.com/input-api/mcp-server/tools/content"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_content.CreateHead_content_typeTool(cfg),
	}
}
