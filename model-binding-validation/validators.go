package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
)

// Binding from json

type Login struct {
	User     string `form:"user" json:"user"  xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"-"`
}

// booking contains binded and validated data
type Booking struct {
	CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

var bookabkeDate validator.Func = func(field_level validator.FieldLevel) bool {
	date, ok := field_level.Field().Interface().(time.Time)
	if ok {
		today := time.Now()
		if today.After(date) {
			return false
		}
	}
	return true
}

func main() {
	router := gin.Default()

	//Example of binding json
	json_binding(router)

	if validtr, ok := binding.Validator.Engine().(*validator.Validate); ok {
		validtr.RegisterValidation("bookabledate", bookabkeDate)
	}

	router.GET("/bookable", getBookable)

	err := router.Run(":8080")
	if err != nil {
		return
	}

}

func getBookable(ctx *gin.Context) {
	var b Booking
	err := ctx.ShouldBindWith(&b, binding.Query)
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Booking dates are valid.!"})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func json_binding(router *gin.Engine) gin.IRoutes {
	return router.POST("/loginJSON", func(context *gin.Context) {
		var json Login
		err := context.ShouldBindJSON(&json)

		//handle error when json bind is returning error
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if json.User != "vrushali" || json.Password != "123!" {
			context.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		context.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})
}
