package controllers

import (
	"fiber-mongo-api/databases"
	"fiber-mongo-api/models"
	"fiber-mongo-api/responses"
	"log"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 用户集合
var userCollection = databases.GetCollection(databases.MongoClient, "users")

// 验证器
var validate = validator.New()

// 用户信息
func UserAllHandler(c *fiber.Ctx) error {
	// 数据模型
	var users []models.User
	// 分页信息
	var findOptions options.FindOptions
	var page int64 = 1
	var pageSize int64 = 10
	if !(c.Params("page") == "") {
		pageInt, _ := strconv.Atoi(c.Params("page"))
		if pageInt == 0 {
			pageInt = 1
		}
		page = int64(pageInt)
	}
	if !(c.Params("pageSize") == "") {
		pageSizeInt, _ := strconv.Atoi(c.Params("pageSize"))
		if pageSizeInt == 0 {
			pageSizeInt = 10
		}
		pageSize = int64(pageSizeInt)
	}
	if pageSize > 0 {
		findOptions.SetSkip((page - 1) * pageSize)
		findOptions.SetLimit(pageSize)
	}
	// 查询数据
	cursor, err := userCollection.Find(c.Context(), bson.M{}, &findOptions)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(responses.ResponseUser{
			Message:    "查询失败",
			StatusCode: fiber.StatusNotFound,
		})
	}
	defer cursor.Close(c.Context())
	// 获取所有数据
	err = cursor.All(c.Context(), &users)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(responses.ResponseUser{
			Message:    "查询失败",
			StatusCode: fiber.StatusNotFound,
		})
	}
	// 查询成功
	return c.Status(fiber.StatusOK).JSON(responses.ResponseMultipleUser{
		Datas:      users,
		Message:    "查询成功",
		StatusCode: fiber.StatusOK,
	})
}

// 用户详情
// https://docs.gofiber.io/api/app#route-handlers
func UserDetailHandler(c *fiber.Ctx) error {
	// 数据模型
	var user models.User
	// 获取编号
	objectId, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ResponseUser{
			Message:    "编号错误",
			StatusCode: fiber.StatusInternalServerError,
		})
	}
	// 查询数据
	err = userCollection.FindOne(c.Context(), bson.M{"_id": objectId}).Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(responses.ResponseUser{
			Message:    "查询失败",
			StatusCode: fiber.StatusNotFound,
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
			Message:    "添加失败",
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
// https://docs.gofiber.io/api/app#route-handlers
func UserChangeHandler(c *fiber.Ctx) error {
	// 数据模型
	var requestUser models.User
	var user models.User
	// 获取编号
	objectId, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ResponseUser{
			Message:    "编号错误",
			StatusCode: fiber.StatusInternalServerError,
		})
	}
	// 验证请求参数
	err = c.BodyParser(&requestUser)
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
	// 查询数据
	err = userCollection.FindOne(c.Context(), bson.M{"_id": objectId}).Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(responses.ResponseUser{
			Message:    "查询失败",
			StatusCode: fiber.StatusNotFound,
		})
	}
	// 更新数据
	result, err := userCollection.UpdateOne(c.Context(), bson.M{"_id": objectId}, bson.M{"$set": requestUser})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ResponseUser{
			Message:    "更新失败",
			StatusCode: fiber.StatusInternalServerError,
		})
	}
	log.Println(result)
	// 查询数据
	err = userCollection.FindOne(c.Context(), bson.M{"_id": objectId}).Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(responses.ResponseUser{
			Message:    "查询失败",
			StatusCode: fiber.StatusNotFound,
		})
	}
	// 更新成功
	return c.Status(fiber.StatusOK).JSON(responses.ResponseSingleUser{
		Data:       user,
		Message:    "更新成功",
		StatusCode: fiber.StatusOK,
	})
}

// 用户删除
// https://docs.gofiber.io/api/app#route-handlers
func UserDeleteHandler(c *fiber.Ctx) error {
	// 数据模型
	var user models.User
	// 获取编号
	objectId, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ResponseUser{
			Message:    "编号错误",
			StatusCode: fiber.StatusInternalServerError,
		})
	}
	// 查询数据
	err = userCollection.FindOne(c.Context(), bson.M{"_id": objectId}).Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(responses.ResponseUser{
			Message:    "查询失败",
			StatusCode: fiber.StatusNotFound,
		})
	}
	// 删除数据
	result, err := userCollection.DeleteOne(c.Context(), bson.M{"_id": objectId})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ResponseUser{
			Message:    "删除失败",
			StatusCode: fiber.StatusInternalServerError,
		})
	}
	log.Println(result.DeletedCount)
	// 删除成功
	return c.Status(fiber.StatusOK).JSON(responses.ResponseSingleUser{
		Data:       user,
		Message:    "删除成功",
		StatusCode: fiber.StatusOK,
	})
}
