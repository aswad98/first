package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	//e.POST("/", hello)
	e.POST("sub", SUB)
	e.POST("add", ADD)
	e.POST("mul", MUL)
	e.POST("mod", MOD)
	e.POST("pow", POW)
	e.POST("div", DIV)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
/*func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}*/

type Input struct {
	Num1 int64 `json:"number1"`
	Num2 int64 `json:"number2"`
}

type Response struct {
	Result int64 `json:"result"`
}

func ADD(c echo.Context) error {
	n := new(Input)

	if err := c.Bind(n); err != nil {
		return err
	}
	add := n.Num1 + n.Num2
	result := Response{
		add,
	}
	return c.JSON(http.StatusOK, result)
}

func MUL(c echo.Context) error {
	n := new(Input)

	if err := c.Bind(n); err != nil {
		return err
	}
	if n.Num1 == 0 || n.Num2 == 0 {
		c.JSON(http.StatusForbidden, "ANY INPUT IS ZERO THE RESULT IS ZERO")
	}
	add := n.Num1 * n.Num2
	result := Response{
		add,
	}
	return c.JSON(http.StatusOK, result)
}

func SUB(c echo.Context) error {
	n := new(Input)

	if err := c.Bind(n); err != nil {
		return err
	}
	add := n.Num1 - n.Num2
	result := Response{
		add,
	}
	return c.JSON(http.StatusOK, result)
}
func DIV(c echo.Context) error {
	n := new(Input)

	if err := c.Bind(n); err != nil {
		return err
	}
	if n.Num2 == 0 {
		c.JSON(http.StatusForbidden, "INVALID INPUT OF NUMBER 2:")
	}
	div := n.Num1 / n.Num2
	result := Response{
		div,
	}
	return c.JSON(http.StatusOK, result)
}

func MOD(c echo.Context) error {
	n := new(Input)
	if err := c.Bind(n); err != nil {
		return err
	}
	MOD := n.Num1 % n.Num2
	result := Response{
		MOD,
	}
	return c.JSON(http.StatusOK, result)
}

func POW(c echo.Context) error {
	n := new(Input)
	if err := c.Bind(n); err != nil {
		return err
	}
	if n.Num1 == 0 {

		c.JSON(http.StatusForbidden, "IF BASE IS ZERO THEN ANY VALUE OF POWER RESULT IS ALWAYS 0")

	}
	if n.Num1 == 1 {

		c.JSON(http.StatusForbidden, "IF BASE IS 1 THEN ANY VALUE OF POWER RESULT IS ALWAYS 1")
	}

	if n.Num2 == 0 {
		c.JSON(http.StatusForbidden, "IF POWER IS ZERO THEN ANY VALUE OF BASE RESULT IS ALWAYS 1")
	}

	var pow int64
	pow = 1
	for n.Num2 != 0 {
		pow *= n.Num1
		n.Num2 -= 1
	}
	result := Response{
		pow,
	}

	return c.JSON(http.StatusOK, result)
}
