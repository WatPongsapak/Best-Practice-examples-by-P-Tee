package api

import (
	"net/http"

	"Advance-Golang-Programming/advanced/topic_2_framework/workshop_structure_answer/internal/config"
	"Advance-Golang-Programming/advanced/topic_2_framework/workshop_structure_answer/internal/db"
	"Advance-Golang-Programming/advanced/topic_2_framework/workshop_structure_answer/internal/grade"
	"Advance-Golang-Programming/advanced/topic_2_framework/workshop_structure_answer/internal/merchant"
	"Advance-Golang-Programming/advanced/topic_2_framework/workshop_structure_answer/internal/ping"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Name        string
	Path        string
	Method      string
	Endpoint    gin.HandlerFunc
	Middlewares []gin.HandlerFunc
}

func Init(conf *config.Config, dbConn *db.Database) http.Handler {
	gradeSrv := grade.NewService()
	gradeHandler := grade.NewHandler(gradeSrv)

	merchantRepo := merchant.NewMongoDB(dbConn.MongoDBConn.Session)
	merchantSrv := merchant.NewService(conf, merchantRepo)
	merchantHandler := merchant.NewHandler(conf, merchantSrv)

	apiv1 := []Route{
		{
			Name:        "ping",
			Method:      http.MethodGet,
			Path:        "/ping",
			Endpoint:    ping.Endpoint,
			Middlewares: []gin.HandlerFunc{},
		},
		{
			Name:        "calculate grade",
			Method:      http.MethodGet,
			Path:        "calculate_grade",
			Endpoint:    gradeHandler.Calculate,
			Middlewares: []gin.HandlerFunc{},
		},
		{
			Name:        "register merchant",
			Method:      http.MethodPost,
			Path:        "/merchant/register",
			Endpoint:    merchantHandler.Register,
			Middlewares: []gin.HandlerFunc{},
		},
		{
			Name:        "inquiry merchant",
			Method:      http.MethodPost,
			Path:        "/merchant/information",
			Endpoint:    merchantHandler.Information,
			Middlewares: []gin.HandlerFunc{},
		},
	}

	ro := gin.New()

	v1 := ro.Group("/v1")
	for _, route := range apiv1 {
		v1.Handle(route.Method, route.Path, append(route.Middlewares, route.Endpoint)...)
	}
	return ro
}
