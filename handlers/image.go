package handlers

import (
	"firstpro/usecase"
	"firstpro/utils/response"
	"image"
	"strconv"

	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
)

// @Summary Crop Product Image
// @Description croping of an exsisting image
// @Tags Image Management
// @Accept json
// @Produce json
// @Security Bearer
// @Param product_image_id query string true "Page Count"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/image-crop [post]
func CropImage(c *gin.Context) {
	imageId := c.Query("product_image_id")
	imageID, err := strconv.Atoi(imageId)
	if err != nil {
		errRes := response.ClientResponse(500, "error in string conversion", nil, err)
		c.JSON(500, errRes)
		return
	}
	imageUrl, err := usecase.CropImage(imageID)
	if err != nil {
		errRes := response.ClientResponse(500, "error in cropping", nil, err)
		c.JSON(500, errRes)
		return
	}

	inputImage, err := imaging.Open(imageUrl)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to open image"})
		return
	}

	cropRect := image.Rect(100, 100, 400, 400)

	croppedImage := imaging.Crop(inputImage, cropRect)

	err = imaging.Save(croppedImage, imageUrl)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to save image"})
		return

	}

	c.JSON(200, response.ClientResponse(200, "Image cropped and saved successfully", nil, nil))

}

// func UploadImage(c *gin.Context) {
// 	cfg, _ := config.LoadConfig()

// 	cld, err := cloudinary.NewFromParams(cfg.CLOUD_NAME, cfg.API_KEY_FOR_CLOUDINARY, cfg.API_SECRET_FOR_CLOUDINARY)
// 	if err != nil {
// 		errRes := response.ClientResponse(500, "Failed to intialize Cloudinary", nil, err)
// 		c.JSON(500, errRes)
// 		return
// 	}

// 	ctx := context.Background()

// 	// resp, err := cld.Upload.Upload(ctx, "my_image.jpg", uploader.UploadParams{})
// 	uploadResult, err := cld.Upload.Upload(
//         ctx,
//         "https://cloudinary-res.cloudinary.com/image/upload/cloudinary_logo.png",
//         uploader.UploadParams{PublicID: "logo"})
//     if err != nil {
//         log.Fatalf("Failed to upload file, %v\n", err)
//     }
// ​
//     log.Println(uploadResult.SecureURL)
// }
