package main

// we dont need external web server
// go already (net/http)

// we need who can generate RESTFUL (data - json)
// we have installed gin

// go get -u github.com/gin-gonic/gin

import (
	"fmt"
	"net/http"

	//"net/http"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type Order struct {
	OrderId int  `json:"orderId"`
	CustomerName string `json:"custName"`
	OrderReview string `json:"review"`
}


func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World from GIN to Swiggy",
	})
}

// slice of orders
var orders = [] Order {
	 Order {101, "Hemanth", "Good Food "},
	Order {102, "Bhavi", "Liked the food"},
	Order {103, "Aarthi", "the seating was good"},
}

func  GetOrders(c *gin.Context)  {
	c.JSON(200, &orders)
}

func PostOrder (c *gin.Context) {
	body := c.Request.Body

	content, err := ioutil.ReadAll(body)
	if err!= nil {
		fmt.Println("Sorry No Content :", err.Error())
	}

	fmt.Println(string(content))

	c.JSON(http.StatusCreated, gin.H {
		"message" :string(content),
	})

// we have to append it to slice ???
}

func GetSpecificOrder (c *gin.Context) {
	orderId := c.Param("orderId")
	order := c.Param("order")

	c.JSON(http.StatusOK, gin.H{
		"orderId" :orderId,
		"order":order,
	})


}

// http://localhost:5656/api/query/orders?minNoOfOrders=100&showOrder=asc
func GetSpecificOrdersByQuery(c *gin.Context) {
	minNumberOfOrders :=  c.DefaultQuery("minNoOfOrders", "100")
	showOrder := c.Query("showOrder")

	c.JSON(200, gin.H{
		"minNoOfOrders": minNumberOfOrders,
		"showOrder":showOrder,
	})

}

func main() {
	// router:=gin.new(), this will not carry middle tier work
	router := gin.Default()

	api:= router.Group("/api")

	// http://localhost:5656/api/
	api.GET("/",  HomePage)
	api.GET("/orders", GetOrders) // to return all orders
	api.POST("/order", PostOrder)

	api.GET("/orders/:orderId/:order", GetSpecificOrder) // to return 1 order
	// if this req is go with query it will be send to the caller
	api.GET("/query/orders/", GetSpecificOrdersByQuery)


	router.Run("localhost:5656") //ipaddress:8080/ping
}

