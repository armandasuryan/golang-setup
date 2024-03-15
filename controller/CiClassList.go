package controller

import (
	"strconv"
	"sygap-cmdb/utils"

	"github.com/gofiber/fiber/v2"
)

func (c *controller) GetCiClassList(ctx *fiber.Ctx) error {
	c.logger.Println("Execute function GetCiClassList")

	page, _ := strconv.Atoi(ctx.Query("page", "1"))
	limit, _ := strconv.Atoi(ctx.Query("limit", "10"))

	CiClassList, err := c.model.GetCiClassList()
	if err != nil {
		c.logger.Println("Error Execute Function GetCiClass In Controller:", err)
		return ctx.Status(404).JSON(utils.Response{
			StatuseCode: 404,
			Message:     "Error",
			Error:       err.Error(),
		})
	}

	data, metapagination := utils.GetPaginated(ctx, page, limit, CiClassList)

	return ctx.Status(200).JSON(utils.Response{
		StatuseCode: 200,
		Message:     "success",
		Data:        data,
		Meta:        utils.Meta{Pagination: metapagination},
	})
}
