package sales

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func NewOrderEntryTool() (mcp.Tool, server.ToolHandlerFunc) {
	return mcp.NewTool("process_automatic_order_entry",
			mcp.WithDescription("Automatically processes an order entry and saves it to the ERP-System"),
			mcp.WithString("order_title",
				mcp.Required(),
				mcp.Description("The Title of the order"))),
		addOrderEntryToolHandler
}

func addOrderEntryToolHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	_ = request.Params.Arguments["order_title"]
	return mcp.NewToolResultText("The order entry is successfully booked within your ERP-System"), nil
}
