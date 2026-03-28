package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/gofiber/fiber/v2"
)

// assume this our own server

func main() {

	app := fiber.New()

	app.Get("/api", api)
	app.Get("/api2", api2)

	app.Listen(":8001")

}

func init() {
	hystrix.ConfigureCommand("api", hystrix.CommandConfig{
		Timeout: 					500,    // ถ้า API ตอบช้ากว่า 500 มิลลิวินาที จะถือว่า request นั้นล้มเหลวทันที
		RequestVolumeThreshold: 	1,      // ใช้ request เพียง 1 ครั้งก็สามารถเริ่มประเมินได้
		ErrorPercentThreshold: 		100,    // ถ้า request ทั้งหมดล้มเหลว (100%) จะเปิด Circuit Breaker
		SleepWindow: 				15000,  // หลังจากเปิด Circuit แล้ว จะรอ 15 วินาที ก่อนลองส่ง request ใหม่อีกครั้ง (half-open)
	})

	hystrix.ConfigureCommand("api2", hystrix.CommandConfig{
		Timeout:                500,
		RequestVolumeThreshold: 4,
		ErrorPercentThreshold:  50,
		SleepWindow:            15000,
	})

	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()
	go func() {
		fmt.Println("Hystrix stream running on :8002")
		err := http.ListenAndServe(":8002", hystrixStreamHandler)
		if err != nil {
			panic(err)
		}
	}()

}

func api(c *fiber.Ctx) error {

	output := make(chan string, 1)
	hystrix.Go("api", func() error {

		res, err := http.Get("http://localhost:8000/api")
		if err != nil {
			return err
		}
		defer res.Body.Close()

		data, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		msg := string(data)
		fmt.Println(msg)

		output <- msg

		return nil
	}, func(err error) error {
		fmt.Println(err)
		return nil
	})

	out := <-output
	
	return c.SendString(out)
}

func api2(c *fiber.Ctx) error {

	output := make(chan string, 1)
	hystrix.Go("api2", func() error {

		res, err := http.Get("http://localhost:8000/api")
		if err != nil {
			return err
		}
		defer res.Body.Close()

		data, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		msg := string(data)
		fmt.Println(msg)

		output <- msg

		return nil
	}, func(err error) error {
		fmt.Println(err)

		return nil
	})

	out := <-output

	return c.SendString(out)
}