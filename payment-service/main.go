package main

import (
    "github.com/gofiber/fiber/v2"
)

type PaymentStatus struct {
    PaymentID string `json:"paymentId"`
    Status    string `json:"status"`
    Amount    int    `json:"amount"`
    Currency  string `json:"currency"`
}

type UserPayment struct {
    Username    string `json:"username"`
    PaymentID string `json:"paymentId"`
    Amount    int    `json:"amount"`
    Currency  string `json:"currency"`
    Status    string `json:"status"`
}

func main() {
    app := fiber.New()

    app.Get("/api-public/payment/status/:paymentId", func(c *fiber.Ctx) error {
        paymentId := c.Params("paymentId")

        status := PaymentStatus{
            PaymentID: paymentId,
            Status:    "Completed",
            Amount:    100,
            Currency:  "USD",
        }

        return c.JSON(status)
    })

    app.Get("/api/users/payments", func(c *fiber.Ctx) error {
        userName := c.Get("X-User-Name")

        if userName == "" {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "error": "Must Login",
            })
        }

        payments := []UserPayment{
            {
                Username:    userName,
                PaymentID: "pmt_12345",
                Amount:    100,
                Currency:  "USD",
                Status:    "Completed",
            },
            {
                Username:    userName,
                PaymentID: "pmt_67890",
                Amount:    200,
                Currency:  "EUR",
                Status:    "Pending",
            },
        }

        return c.JSON(payments)
    })

    app.Listen(":3000")
}
// go run main.go
