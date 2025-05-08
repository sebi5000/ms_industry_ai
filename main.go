package main

import (
	"github.com/mark3labs/mcp-go/server"
	"ms_salespower_mcp/usecases"
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

	//STANDARD IN OUT FOR LOCAL MCP
	//s := server.NewStdioServer(mcpServer)
	//err := s.Listen(context.Background(), os.Stdin, os.Stdout)

	sseServer := server.NewSSEServer(mcpServer, server.WithBaseURL("http://localhost"))

	if err := sseServer.Start(":3001"); err != nil {
		panic(err)
	}
}
