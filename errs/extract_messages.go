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

			/*
				Sometimes "detailedError" is not a pointer anymore.
				I think it has to do with something internally within gin.
				Therefore, check for a detailed error structure and pointer.
			*/

			var detailedError DetailedError
			var detailedErrorPointer *DetailedError
			if errors.As(err, &detailedError) {
				err = detailedError.innerError
			} else if errors.As(err, &detailedErrorPointer) {
				err = detailedErrorPointer.innerError
			} else {
				err = nil
			}
		}
	}

	return result
}
