package mapping

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"payment-simulator/internal/models"
)

func MapIsoPacs008ToPo(c *gin.Context) (*models.PaymentOrder, error) {
	var po models.PaymentOrder
	if err := c.ShouldBindXML(&po); err != nil {
		fmt.Println("Error in Binding to XML for PaymentOrder: ", po)
		return nil, err
	} else {
		return &po, nil
	}
}
