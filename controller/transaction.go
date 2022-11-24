package controller

import (
	"github.com/NubeIO/rubix-assist/amodel"
	"github.com/gin-gonic/gin"
)

func getTransactionBody(ctx *gin.Context) (dto *amodel.Transaction, err error) {
	err = ctx.ShouldBindJSON(&dto)
	return dto, err
}

func (inst *Controller) TransactionsSchema(ctx *gin.Context) {
	// responseHandler(schema.GetTransactionSchema(), nil, ctx)
}

func (inst *Controller) GetTransaction(c *gin.Context) {
	team, err := inst.DB.GetTransaction(c.Params.ByName("uuid"))
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(team, err, c)
}

func (inst *Controller) GetTransactions(c *gin.Context) {
	teams, err := inst.DB.GetTransactions()
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(teams, err, c)
}

func (inst *Controller) CreateTransaction(c *gin.Context) {
	m := new(amodel.Transaction)
	err := c.ShouldBindJSON(&m)
	team, err := inst.DB.CreateTransaction(m)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(team, err, c)
}

func (inst *Controller) UpdateTransaction(c *gin.Context) {
	body, _ := getTransactionBody(c)
	team, err := inst.DB.UpdateTransaction(c.Params.ByName("uuid"), body)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(team, err, c)
}

func (inst *Controller) DeleteTransaction(c *gin.Context) {
	q, err := inst.DB.DeleteTransaction(c.Params.ByName("uuid"))
	if err != nil {
		responseHandler(nil, err, c)
	} else {
		responseHandler(q, err, c)
	}
}

func (inst *Controller) DropTransactions(c *gin.Context) {
	team, err := inst.DB.DropTransactions()
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(team, err, c)
}
