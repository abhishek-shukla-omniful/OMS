package s3Service

import (
	"github.com/Abhishek-Omniful/OMS/mycontext"
	awsS3 "github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/omniful/go_commons/i18n"
	"github.com/omniful/go_commons/log"
	"github.com/omniful/go_commons/s3"
)

var s3Client *awsS3.Client
var logger = log.DefaultLogger()
var err error

func ConnectS3() {
	ctx := mycontext.GetContext()

	logger.Info(i18n.Translate(ctx, "Connecting to S3..."))

	s3Client, err = s3.NewDefaultAWSS3Client()
	if err != nil {
		logger.Errorf(i18n.Translate(ctx, "Error connecting to S3: %v"), err)
		return
	}
	logger.Info(i18n.Translate(ctx, "Successfully connected to S3"))
}

func GetS3Client() *awsS3.Client {
	return s3Client
}
