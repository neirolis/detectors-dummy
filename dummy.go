package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Debug = true
	e.HideBanner = true
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "	${remote_ip} [${method}] ${host}${uri} ${status} ${bytes_out} ${error}\n",
	}))

	e.GET("/status", status)

	var reverse bool
	x := float32(0.0)
	y := float32(0.8)
	id := 1

	e.POST("/", func(c echo.Context) error {
		items := []Item{
			{
				ID:   fmt.Sprint(id),
				Type: "dummy",
				BBox: [4]float32{x, y, 0.1, 0.1},
				Ang:  []float32{0.1, 0.2},
				Keypoints: map[string][2]float32{
					"0": [2]float32{x + 0.1, 0.1},
					"1": [2]float32{x + 0.1, 0.4},
					"2": [2]float32{x + 0.2, 0.6},
					"3": [2]float32{x + 0.3, 0.9},
				},
				Lines: [][]float32{
					[]float32{
						x + 0.1, 0.1,
						x + 0.1, 0.4,
					},
					[]float32{
						x + 0.2, 0.6,
						x + 0.3, 0.9,
					},
					[]float32{
						0.05, 0.05,
						0.95, 0.05,
						0.95, 0.95,
						0.05, 0.95,
						0.05, 0.05,
					},
				},
				Attributes: map[string]float64{
					"area": 1.23,
				},
			},
			// Item{
			// 	ID:   "2",
			// 	Type: "dummy",
			// 	BBox: [4]float32{x + 0.02, y + 0.02, 0.1, 0.1},
			// },
		}

		if reverse {
			x -= 0.05
			// y -= 0.05
		} else {
			x += 0.05
			// y += 0.05
		}
		if x > 0.9 {
			reverse = true
		}
		if x <= 0 {
			reverse = false
			id++
		}

		return c.JSON(200, items)
	})

	if err := e.Start(":64400"); err != nil {
		log.Fatal(err)
	}
}

type Item struct {
	ID         string                `json:"id"`
	Type       string                `json:"type"`
	BBox       [4]float32            `json:"bbox"`
	Ang        []float32             `json:"ang,omitempty"`
	Keypoints  map[string][2]float32 `json:"keypoints"`
	Lines      [][]float32           `json:"lines"`
	Attributes map[string]float64    `json:"attributes"`
}

func status(c echo.Context) error {
	return c.JSON(200, echo.Map{
		"name": "dummy",
		"type": "detector",
		"path": "/",
		"output": map[string][]string{
			"types": []string{"dummy"},
		},
		"variables": []Variable{
			{Name: "test", Type: "text"},
		},
		"version": "0.0.1",
	})
}

type Variable struct {
	Name string
	Type string
}
