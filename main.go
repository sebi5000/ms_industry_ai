package main

import (
	"context"
	"github.com/mark3labs/mcp-go/server"
	"ms_salespower_mcp/usecases"
	"os"
)

func main() {
	mcpServer := server.NewMCPServer(
		"ms_salespower",
		"0.0.1",
		server.WithResourceCapabilities(true, true),
		server.WithLogging(),
		server.WithRecovery(),
	)

	//Add Prompts which should be active
	visitReportPrompt, visitReportHandler := usecases.NewFormatVisitReportPrompt()
	mcpServer.AddPrompt(visitReportPrompt, visitReportHandler)

	//Add Tools which should be active
	addVisitToSalesforceTool, addVisitToSalesforceHandler := usecases.NewAddVisitReportToSalesforceTool()
	mcpServer.AddTool(addVisitToSalesforceTool, addVisitToSalesforceHandler)

	s := server.NewStdioServer(mcpServer)
	err := s.Listen(context.Background(), os.Stdin, os.Stdout)

	if err != nil {
		panic(err)
	}
}
