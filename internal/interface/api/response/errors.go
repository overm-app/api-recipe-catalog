package response

import (
	"errors"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	appErrors "github.com/overm-app/api-recipe-catalog/internal/domain/errors"
)

func HandleError(c *gin.Context, sugar *zap.SugaredLogger, err error, requestID string) {
    var appErr *appErrors.AppError

    if errors.As(err, &appErr) {
        c.Set("error_code", string(appErr.Code))
        if appErr.Err != nil {
            c.Set("error_cause", appErr.Err.Error())
        }
        c.AbortWithStatusJSON(appErr.HTTPStatus, gin.H{
            "code":    string(appErr.Code),
            "message": appErr.Message,
        })
        return
    }

    c.Set("error_cause", err.Error())
    c.Set("error_code", string(appErrors.ErrInternal))
    c.AbortWithStatusJSON(500, gin.H{
        "code":    string(appErrors.ErrInternal),
        "message": "An unexpected error occurred",
    })
}
