package errs

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func ExtractMessages(c *gin.Context) []string {
	var result []string

	for i := len(c.Errors) - 1; i >= 0; i-- {
		err := c.Errors[i].Err

		for err != nil {
			result = append(result, err.Error())

			var detailedError *DetailedError
			if errors.As(err, detailedError) {
				err = detailedError.innerError
			}
		}
	}

	return result
}
