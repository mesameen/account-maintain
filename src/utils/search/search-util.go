package searchUtil

import (
	"fmt"
	errorHelpers "go-gin-test-job/src/common/error-helpers"
	"go-gin-test-job/src/logger"
	"strings"

	"github.com/gin-gonic/gin"
)

// GetSearchByParams fetches search columns and values from search query param
func GetSearchByParams(c *gin.Context, data, separator string, availableSearchFieldList []string) (map[string]string, error) {
	searchResult := make(map[string]string)
	availableSearchFields := make(map[string]bool)
	// Convert availableSearchFieldList to a map for faster lookup
	for _, field := range availableSearchFieldList {
		availableSearchFields[field] = true
	}
	// Split and process the search parameters
	searchList := strings.Split(data, separator)
	for _, searchLine := range searchList {
		searchLine = strings.TrimSpace(searchLine)
		if searchLine == "" {
			continue
		}
		parts := strings.Fields(searchLine) // Splits by whitespace
		if len(parts) != 2 {
			logger.Logger.Error().Msg(fmt.Sprintf("invalid search parameter: %s", searchLine))
			// Return a structured bad request error
			return nil, errorHelpers.RespondBadRequestError(c, fmt.Sprintf("invalid search parameter: %s", searchLine))
		}
		colName := strings.TrimSpace(parts[0])
		colVal := strings.TrimSpace(parts[1])
		// Validate order field
		if !availableSearchFields[colName] {
			logger.Logger.Error().Msg(fmt.Sprintf("unsupported column (%s) to search.", colName))
			// Return a structured bad request error
			return nil, errorHelpers.RespondBadRequestError(c, fmt.Sprintf("unsupported column (%s) to search.", colName))
		}
		if colVal == "" {
			logger.Logger.Error().Msg(fmt.Sprintf("cannot search by column: %s value: %s", colName, colVal))
			// Return a structured bad request error
			return nil, errorHelpers.RespondBadRequestError(c, fmt.Sprintf("column (%s) value isn't valid.", colName))
		}
		// Avoid duplicate order fields
		if _, exists := searchResult[colName]; !exists {
			searchResult[colName] = fmt.Sprintf("%%%s%%", colVal)
		}
	}

	return searchResult, nil
}
