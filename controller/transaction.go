package controller

import (
	"github.com/NubeIO/rubix-assist/pkg/model"
	"github.com/gin-gonic/gin"
)

func getTransactionBody(ctx *gin.Context) (dto *model.Transaction, err error) {
	err = ctx.ShouldBindJSON(&dto)
	return dto, err
}

func (inst *Controller) TransactionsSchema(ctx *gin.Context) {
	//reposeHandler(schema.GetTransactionSchema(), nil, ctx)
}

func (inst *Controller) GetTransaction(c *gin.Context) {
	team, err := inst.DB.GetTransaction(c.Params.ByName("uuid"))
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(team, err, c)
}

func (inst *Controller) GetTransactions(c *gin.Context) {
	teams, err := inst.DB.GetTransactions()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(teams, err, c)
}

func (inst *Controller) CreateTransaction(c *gin.Context) {
	m := new(model.Transaction)
	err = c.ShouldBindJSON(&m)
	team, err := inst.DB.CreateTransaction(m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(team, err, c)
}

func (inst *Controller) UpdateTransaction(c *gin.Context) {
	body, _ := getTransactionBody(c)
	team, err := inst.DB.UpdateTransaction(c.Params.ByName("uuid"), body)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(team, err, c)
}

func (inst *Controller) DeleteTransaction(c *gin.Context) {
	q, err := inst.DB.DeleteTransaction(c.Params.ByName("uuid"))
	if err != nil {
		reposeHandler(nil, err, c)
	} else {
		reposeHandler(q, err, c)
	}
}

func (inst *Controller) DropTransactions(c *gin.Context) {
	team, err := inst.DB.DropTransactions()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(team, err, c)
}
