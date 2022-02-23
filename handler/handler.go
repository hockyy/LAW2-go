package handler

import (
	gen "LAW2-go/gen/openapi"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"sync"
)

type ApplicationServer struct {
	Lock sync.Mutex
}

func isSame(a, b string) bool {
	lenStr := len(a)
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < lenStr; i++ {
		if a[i] == b[lenStr-i-1] {
			continue
		}
		return false
	}
	return true
}

func (p *ApplicationServer) LoginLoginPost(ctx echo.Context) error {
	var loginPost gen.BodyLoginLoginPost
	loginPost.Username = ctx.FormValue("username")
	loginPost.Password = ctx.FormValue("password")
	fmt.Println(loginPost)
	msg := isSame(loginPost.Username, loginPost.Password)
	strMsg := "not authorized in golang"
	if msg {
		strMsg = "authorized in golang"
	}
	code := 200
	return ctx.JSON(http.StatusOK, gen.MessageResponse{
		Code: &code,
		Msg:  &strMsg,
	})
}

func (p *ApplicationServer) ReadModuloModuloPost(ctx echo.Context) error {
	var mathQuery gen.MathQuery
	_ = ctx.Bind(&mathQuery)
	fmt.Println(mathQuery)
	if mathQuery.B == 0 {
		ret := -404
		return ctx.JSON(http.StatusOK, gen.MathResponse{
			Query:  &mathQuery,
			Result: &ret,
		})
	}
	ret := mathQuery.A % mathQuery.B
	return ctx.JSON(http.StatusOK, gen.MathResponse{
		Query:  &mathQuery,
		Result: &ret,
	})
}

func binaryPower(a, b int) *int {
	ret := 1
	MOD := 1000000007
	for b > 0 {
		if b&1 == 1 {
			ret = (ret * a) % MOD
		}
		a = (a * a) % MOD
		b >>= 1
	}
	return &ret
}

func (p *ApplicationServer) ReadPowerPowerABGet(ctx echo.Context, a int, b int) error {
	ret := 200
	return ctx.JSON(http.StatusOK, gen.ValueResponse{
		Code: &ret,
		Msg:  binaryPower(a, b),
	})
}

func NewApplicationServer() *ApplicationServer {
	return &ApplicationServer{}
}

func (p *ApplicationServer) ReadRootGet(ctx echo.Context) error {
	ret := 200
	resp := "Hello world through golang!"
	return ctx.JSON(http.StatusOK, gen.MessageResponse{
		Code: &ret,
		Msg:  &resp,
	})
}
