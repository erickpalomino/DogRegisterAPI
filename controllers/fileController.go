package controllers

import (
	"context"
	"fmt"
	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

func CloudinaryUpload(path string, dni string) (url string, err error) {
	cld, _ := cloudinary.NewFromParams("dziaapbmr", "922581187159196", "sY44Tzpsnok0L-SSYx3JhtbF73I")
	ctx := context.Background()
	resp, err := cld.Upload.Upload(ctx, path, uploader.UploadParams{PublicID: dni,
		Transformation: "c_crop,g_center/q_auto/f_auto", Tags: []string{"dogapp"}})
	if err != nil {
		fmt.Println("error" + err.Error())
	}
	url = resp.SecureURL
	fmt.Println("URLS")
	fmt.Println(url)
	return
}
