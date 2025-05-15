package hr

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"math/rand/v2"
)

//TOOLS for Absences

func NewAddAbsenceToHRTool() (mcp.Tool, server.ToolHandlerFunc) {
	return mcp.NewTool("add_absence_to_hr",
			mcp.WithDescription("Adds a new absence to the HR System"),
			mcp.WithString("absence_type",
				mcp.Required(),
				mcp.Description("The type of absence e.g. holiday")),
			mcp.WithString("begin",
				mcp.Required(),
				mcp.Description("The begin date of the absence")),
			mcp.WithString("end",
				mcp.Required(),
				mcp.Description("The end date of the absence")),
		),
		addAbsenceToHRToolHandler
}

func addAbsenceToHRToolHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	absenceType := request.Params.Arguments["absence_type"]
	begin := request.Params.Arguments["begin"]
	end := request.Params.Arguments["end"]
	return mcp.NewToolResultText(fmt.Sprintf("Die Abwesenheit %s vom: %s bis zum: %s wurde erfolgreich hinzugefügt", absenceType, begin, end)), nil
}

func NewGetContingentTool() (mcp.Tool, server.ToolHandlerFunc) {
	return mcp.NewTool("get_contingent",
			mcp.WithDescription("Gets the user current contingent of a leave type e.g. holiday"),
			mcp.WithString("absence_type",
				mcp.Required(),
				mcp.Description("The type of absence e.g. holiday")),
		),
		addAbsenceToHRToolHandler
}

func getContingentHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	absenceType := request.Params.Arguments["absence_type"]
	contingent := rand.IntN(30)
	return mcp.NewToolResultText(fmt.Sprintf("Your current contingent for absence type %s is %d", absenceType, contingent)), nil
}
