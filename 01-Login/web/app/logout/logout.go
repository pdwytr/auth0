package logout

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler for our logout.
// func Handler(ctx *gin.Context) {
// 	logoutUrl, err := url.Parse("https://" + os.Getenv("AUTH0_DOMAIN") + "/v2/logout")
// 	if err != nil {
// 		ctx.String(http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	scheme := "http"
// 	if ctx.Request.TLS != nil {
// 		scheme = "https"
// 	}

// 	returnTo, err := url.Parse(scheme + "://" + ctx.Request.Host)
// 	if err != nil {
// 		ctx.String(http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	parameters := url.Values{}
// 	parameters.Add("returnTo", returnTo.String())
// 	parameters.Add("client_id", os.Getenv("AUTH0_CLIENT_ID"))
// 	logoutUrl.RawQuery = parameters.Encode()

// 	ctx.Redirect(http.StatusTemporaryRedirect, logoutUrl.String())
// }

// RenderLogoutPage renders a simple HTML logout page.
func Handler(ctx *gin.Context) {
	htmlContent := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Logged Out</title>
	</head>
	<body>
		<h1>You have been successfully logged out</h1>
	</body>
	</html>
	`
	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(htmlContent))
}
