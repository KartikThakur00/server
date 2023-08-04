package controller

import (
	"fmt"
	"path/filepath"
	"react-gin/server/config"
	"react-gin/server/models"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func Start(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "Server Started"})
}

func GetApi(c *gin.Context) {
	employees := []models.Employee{}
	config.DB.Find(&employees)
	c.JSON(200, &employees)
}

func GetFile(c *gin.Context) {
	fileName :=c.Param("file")
	fmt.Println(fileName)
	filePath := filepath.Join("uploads", fileName)
	fmt.Println(filePath)

	// Set the headers for the file transfer and return the file
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	c.File(filePath)
}

func PostApi(c *gin.Context) {
	name := c.Request.FormValue("name")
	gender := c.Request.FormValue("gender")
	from := c.Request.FormValue("from")
	to := c.Request.FormValue("to")
	phone := c.Request.FormValue("phone")
	email := c.Request.FormValue("email")

	finalName := strings.TrimSpace(name)
	finalGender := gender

	// number validation
	resPhn, _ := regexp.MatchString(`^(\+?91[\s]?)?[6789]\d{9}$`, phone)
	if !resPhn {
		c.JSON(400, gin.H{"error": "number not valid"})
		return
	}

	finalPhone, _ := strconv.ParseInt(phone, 10, 0)

	// temp storing file
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	//file validating
	if file.Size > 5000000 || file.Size==0 {
		c.JSON(400, gin.H{"error": "file size can't exceed 5 MB"})
		return
	}

	fname := file.Filename

	resPdf, _ := regexp.MatchString(`\.pdf$`, fname)
	resPng, _ := regexp.MatchString(`\.png$`, fname)
	var fileName string
	if resPdf {
		fileName = (name + "_" + phone + ".pdf")
	} else if resPng {
		fileName = (name + "_" + phone + ".png")
	} else {
		c.JSON(400, gin.H{"error": "file type not acceptable"})
		return
	}

	// Define the path where the file will be saved

	filePath := filepath.Join("uploads", fileName)
	fmt.Println(filePath)
	// Save the file to the defined path
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(500, gin.H{"error": "Failed to save file"})
		return
	}

	//email validation
	resMail, _ := regexp.MatchString(`^[a-zA-Z0-9.!#$%&'*+=?^_{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`, email)
	if !resMail {
		c.JSON(400, gin.H{"error": "email not valid"})
		return
	}

	finalMail := email

	//date validation

	finalFrom, _ := time.Parse("2006-01-02", from)
	finalTo, _ := time.Parse("2006-01-02", to)
	// Save file metadata to database
	finalData := models.Employee{
		Name:     finalName,
		Gender:   finalGender,
		FromDate: finalFrom,
		ToDate:   finalTo,
		Phone:    finalPhone,
		Resume:   fileName,
		Email:    finalMail,
	}
	config.DB.Create(&finalData)
	c.JSON(201, gin.H{"message": "Details uploaded successfully", "Details": finalData})

}
