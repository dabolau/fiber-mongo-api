package controllers

import (
	"fiber-mongo-api/databases"
	"fiber-mongo-api/models"
	"fiber-mongo-api/responses"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// 用户集合
var userCollection = databases.GetCollection(databases.MongoClient, "users")

// 验证器
var validate = validator.New()

// 用户信息
func UserHandler(c *fiber.Ctx) error {

	return c.Status(fiber.StatusOK).JSON(responses.ResponseUser{
		Message:    "查询所有成功",
		StatusCode: fiber.StatusOK,
	})
}

// 用户详情
func UserDetailHandler(c *fiber.Ctx) error {
	// 数据模型
	var user models.User
	// 获取编号
	objectId, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ResponseUser{
			Message:    "服务器错误",
			StatusCode: fiber.StatusInternalServerError,
		})
	}
	// 查询数据
	err = userCollection.FindOne(c.Context(), bson.M{"_id": objectId}).Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ResponseUser{
			Message:    "数据库服务器错误",
			StatusCode: fiber.StatusInternalServerError,
		})
	}
	// 查询成功
	return c.Status(fiber.StatusOK).JSON(responses.ResponseSingleUser{
		Data:       user,
		Message:    "查询成功",
		StatusCode: fiber.StatusOK,
	})

}

// 用户添加
// https://docs.gofiber.io/api/app#route-handlers
func UserAddHandler(c *fiber.Ctx) error {
	// 数据模型
	var requestUser models.User
	// 验证请求参数
	err := c.BodyParser(&requestUser)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ResponseUser{
			Message:    "请求参数错误",
			StatusCode: fiber.StatusBadRequest,
		})
	}
	// 使用验证器验证必填参数
	err = validate.Struct(&requestUser)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ResponseUser{
			Message:    "验证参数错误",
			StatusCode: fiber.StatusBadRequest,
		})
	}
	// 新增编号
	requestUser.ID = primitive.NewObjectID()
	// 添加数据
	result, err := userCollection.InsertOne(c.Context(), &requestUser)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ResponseUser{
			Message:    "数据库服务器错误",
			StatusCode: fiber.StatusInternalServerError,
		})
	}
	log.Println(result.InsertedID)
	// 添加成功
	return c.Status(fiber.StatusOK).JSON(responses.ResponseSingleUser{
		Data:       requestUser,
		Message:    "添加成功",
		StatusCode: fiber.StatusOK,
	})
}

// 用户编辑
func UserChangeHandler(c *fiber.Ctx) error {

	return c.Status(fiber.StatusOK).JSON(responses.ResponseUser{
		Message:    "编辑成功",
		StatusCode: fiber.StatusOK,
	})
}

// 用户删除
func UserDeleteHandler(c *fiber.Ctx) error {
	// 获取编号
	objectId, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ResponseUser{
			Message:    "服务器错误",
			StatusCode: fiber.StatusInternalServerError,
		})
	}
	// 删除数据
	result, err := userCollection.DeleteOne(c.Context(), bson.M{"_id": objectId})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ResponseUser{
			Message:    "数据库服务器错误",
			StatusCode: fiber.StatusInternalServerError,
		})
	}
	// 编号未找到
	if result.DeletedCount < 1 {
		return c.Status(fiber.StatusNotFound).JSON(responses.ResponseUser{
			Message:    "编号未找到",
			StatusCode: fiber.StatusNotFound,
		})
	}
	// 删除成功
	return c.Status(fiber.StatusOK).JSON(responses.ResponseUser{
		Message:    "删除成功",
		StatusCode: fiber.StatusOK,
	})
}
