package main

import (
	"archive/zip"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"github.com/gin-gonic/gin"
)


func main() {
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.Static("/", "./public")

	router.POST("/upload", func(c *gin.Context) {
		name := c.PostForm("name")
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		files := form.File["files"]
		if len(files) == 0 {
			c.String(http.StatusBadRequest, "no files uploaded")
			return
		}

		zipName := "./uploads/" + name + ".zip"
		os.MkdirAll("uploads", 0755)

		zipFile, err := os.Create(zipName)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		defer zipFile.Close()

		zipWriter := zip.NewWriter(zipFile)
		defer zipWriter.Close()

		for _, fileHeader := range files {
			src, err := fileHeader.Open()
			if err != nil {
				c.String(http.StatusInternalServerError, err.Error())
				return
			}

			// Prevent Zip Slip
			filename := filepath.Base(fileHeader.Filename)

			dst, err := zipWriter.Create(filename)
			if err != nil {
				src.Close()
				c.String(http.StatusInternalServerError, err.Error())
				return
			}

			_, err = io.Copy(dst, src)
			src.Close()
			if err != nil {
				c.String(http.StatusInternalServerError, err.Error())
				return
			}
		}

		c.String(http.StatusOK, "All files zipped successfully")
	})

	router.Run(":8080")
}
