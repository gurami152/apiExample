package main

import (
	"github.com/gin-gonic/gin"
	"math"
	"strconv"
	"strings"
)

type resultObject struct {
	Index  int
	Number int
	Answer bool
}

type numbers struct {
	Numbers string `json:"data"`
}

func checkNumber(c *gin.Context) {
	number := c.Param("number")

	var numberInt int
	numberInt, _ = strconv.Atoi(number)
	var result = checkPrimeNumber(numberInt)

	c.JSON(200, gin.H{
		"data": result,
	})
}

// postAlbums adds an album from JSON received in the request body.
func checkNumbers(c *gin.Context) {

	var body struct {
		Numbers string
	}

	err := c.Bind(&body)
	if err != nil {
		return
	}

	var number []int = toIntArray(body.Numbers)

	var result bool

	var results []resultObject

	var valueToAppend resultObject

	for index, element := range number {
		result = checkPrimeNumber(element)

		valueToAppend = resultObject{
			Index:  index,
			Number: element,
			Answer: result,
		}

		results = append(results, valueToAppend)
		// index is the index where we are
		// element is the element from someSlice for where we are
	}

	c.JSON(200, gin.H{
		"data": results,
	})
}

func checkPrimeNumber(num int) bool {
	if num < 2 {
		return false
	}
	sqRoot := int(math.Sqrt(float64(num)))
	for i := 2; i <= sqRoot; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func toIntArray(str string) []int {
	chunks := strings.Split(str, ",")

	var res []int
	for _, c := range chunks {
		i, _ := strconv.Atoi(c) // error handling ommitted for concision
		res = append(res, i)
	}

	return res
}

func main() {
	router := gin.Default()
	router.POST("/numbers", checkNumbers)
	router.GET("/number/:number", checkNumber)
	router.Run()
}
