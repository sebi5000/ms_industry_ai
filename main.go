package main

import (
	"context"
	"flag"
	"github.com/mark3labs/mcp-go/server"
	"ms_salespower_mcp/usecases"
	"os"
)

func main() {

	var transport, baseURL, port string

	//MCP Server flags
	flag.StringVar(&transport, "transport", "stdio", "Transport to use (STDIO or SSE)")
	flag.StringVar(&baseURL, "baseURL", "http//:localhost", "Base URL")
	flag.StringVar(&port, "port", "3001", "Port")

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

	//Start Server in SSE Mode
	if transport == "sse" {
		sseServer := server.NewSSEServer(mcpServer, server.WithBaseURL(baseURL))
		err := sseServer.Start(":" + port)

		if err != nil {
			panic(err)
		}

	} else { //Start Server in STDIO Mode
		s := server.NewStdioServer(mcpServer)
		err := s.Listen(context.Background(), os.Stdin, os.Stdout)

		if err != nil {
			panic(err)
		}
	}
}
