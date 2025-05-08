package usecases

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// PROMPT Templates for Visit Reports

// NewFormatVisitReportPrompt provides the prompt and handler function to format a visit report
func NewFormatVisitReportPrompt() (mcp.Prompt, server.PromptHandlerFunc) {
	return mcp.NewPrompt("format_visit_report",
			mcp.WithPromptDescription("Formats a sales visit report into a more readable and structured format")),
		formatVisitReportPromptHandler
}

func formatVisitReportPromptHandler(ctx context.Context, request mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {

	promptText := `You are a sales assistant and formatting the visit reports of your colleagues. If someone enters a 
visit report, you will format it like the following:

Visit of customer <CUSTOMER NAME>
1. Participants: <PARTICIPANT-1>, <PARTICIPANT-2> ...
2. Content: *Here you add the content given by the user. But you summarize it to the meaningful insights with maximum of 500 words*
3. Tasks: <TASK-1>, <TASK-2>, <TASK-3> ...

Notice everything between <> are placeholder which you have to fill with data from the visit report. Ignore if there
is no clear data in the visit report about it. Furthermore everything between ** should be handled as a instruction.
`
	return mcp.NewGetPromptResult("The formatted visit report",
		[]mcp.PromptMessage{
			mcp.NewPromptMessage(mcp.RoleAssistant,
				mcp.NewTextContent(promptText)),
		}), nil
}

//TOOLS for Visit Reports

func NewAddVisitReportToSalesforceTool() (mcp.Tool, server.ToolHandlerFunc) {
	return mcp.NewTool("add_visit_report_to_salesforce",
			mcp.WithDescription("Adds a visit report to Salesforce"),
			mcp.WithString("visit_report",
				mcp.Required(),
				mcp.Description("The visit report which should to be added"))),
		addVisitReportToSalesforceToolHandler
}

func addVisitReportToSalesforceToolHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	_ = request.Params.Arguments["visit_report"]
	return mcp.NewToolResultText("The visit report is successfully added to Salesforce"), nil
}
