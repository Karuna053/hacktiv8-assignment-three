package controllers

import (
	"autoreload-data/model"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func AutoReloadHandler(c *gin.Context) {
	// return the response
	c.String(200, "Auto Reload")
	// return nil

	// c.JSON(200, gin.H{
	// 	"message": "Auto Reload",
	// })
}

func AutoReload() {
	for {
		// call the function with goroutine
		min := 1
		max := 100

		// generate random number
		numberWater := min + rand.Intn(max-min)
		numberWind := min + rand.Intn(max-min)

		// update data file json
		updateData(numberWater, numberWind)

		// log water
		logWater(numberWater)

		// log wind
		logWind(numberWind)

		// makeAPIRequest
		makeAPIRequest(numberWater, numberWind)

		// sleep for 15 seconds
		time.Sleep(15 * time.Second)
	}
}

func updateData(numberWater int, numberWind int) {
	data := model.Data{}

	data.Water = numberWater
	data.Wind = numberWind

	dataWaterWind := model.Status{
		Status: data,
	}

	// update the data file jsonprint
	jsonPrint, _ := json.MarshalIndent(dataWaterWind, "", "    ")

	err := os.WriteFile("file.json", jsonPrint, 0644)
	if err != nil {
		panic(err)
	}

	log.Println(string(jsonPrint))
}

func logWater(numberWater int) {
	if numberWater < 5 {
		result := "aman"
		fmt.Println("status water : ", result)
	} else if numberWater > 5 && numberWater <= 8 {
		result := "siaga"
		fmt.Println("status water : ", result)
	} else if numberWater > 8 {
		result := "bahaya"
		fmt.Println("status water : ", result)
	}
}

func logWind(numberWind int) {
	if numberWind < 5 {
		result := "aman"
		fmt.Println("status Wind : ", result)
	} else if numberWind > 5 && numberWind <= 8 {
		result := "siaga"
		fmt.Println("status Wind : ", result)
	} else if numberWind > 8 {
		result := "bahaya"
		fmt.Println("status Wind : ", result)
	}
}

func makeAPIRequest(numberWater int, numberWind int) {
	data := model.Data{
		Water: numberWater,
		Wind:  numberWind,
	}

	status := model.Status{
		Status: data,
	}

	_, err := json.Marshal(status)
	if err != nil {
		log.Println(err)
	}

	// make the api request

	resp, err := http.Post("http://localhost:8080/auto-reload", "application/json", nil)
	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("API Request Success")
}
