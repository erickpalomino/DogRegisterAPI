package controllers

import (
	"dog-app/models"
	"dog-app/utils/token"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterDiagnostic(c *gin.Context) {
	var diag models.Diagnostic
	if err := c.ShouldBindJSON(&diag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user_id, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u, err := models.GetUserByID(user_id)
	diagnostic, err := diag.SaveDiagnostic()
	diagnostic.Doctor = u.Username
	models.GetDB().Save(&diagnostic)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "diagnostic registration success", "data": diagnostic})
}

func UploadXrayBloodResult(c *gin.Context) {
	idstring := c.Params.ByName("id")
	id, paramerr := strconv.ParseUint(idstring, 10, 64)
	if paramerr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": paramerr.Error()})
		return
	}
	diagnostic, qerr := models.GetDiagnosticById(id)
	if qerr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": qerr.Error()})
		return
	}
	br, brh, brerr := c.Request.FormFile("br")
	if brerr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": brerr.Error()})
		return
	}
	xray, xrayh, xrayerr := c.Request.FormFile("xray")
	if xrayerr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": xrayerr.Error()})
		return
	}
	brout, oserr := os.Create("public/" + brh.Filename)
	if oserr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": oserr.Error()})
		return
	}
	defer brout.Close()
	_, cpyerr := io.Copy(brout, br)
	if cpyerr != nil {
		log.Fatal(cpyerr)
		return
	}

	xrayout, osxerr := os.Create("public/" + xrayh.Filename)
	if osxerr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": osxerr.Error()})
		return
	}
	defer xrayout.Close()
	_, cpyerr = io.Copy(xrayout, xray)
	if cpyerr != nil {
		log.Fatal(cpyerr)
		return
	}
	brfilepath := "./public/" + brh.Filename
	brurl, errbrurl := CloudinaryUpload(brfilepath, "br"+strconv.FormatUint(id, 10))
	if errbrurl != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": errbrurl.Error()})
	}
	xrayfilepath := "./public/" + xrayh.Filename
	xrayurl, errxurl := CloudinaryUpload(xrayfilepath, "xray"+strconv.FormatUint(id, 10))
	if errxurl != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": errxurl.Error()})

	}
	diagnostic.BloodResult = brurl
	diagnostic.XrayPic = xrayurl
	models.GetDB().Save(&diagnostic)
	return
}
