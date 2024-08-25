package s3

import (
	"bytes"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/lucs-t/tshell/utils"
)

type S3Client struct {
	cli *s3.S3
	updatePath   string
	bucket 	 string
}

func NewBucket(data map[string]string) (*S3Client,error) {
	ak := data[utils.AK]
	sk := data[utils.SK]
	region := data[utils.Region]
	bucket := data[utils.Bucket]
	updatePath := data[utils.UpdatePath]
	endpoint := data[utils.Endpoint]
	sess, err := session.NewSession(&aws.Config{
		Region:           aws.String(region),
		Endpoint:         aws.String(endpoint),
		Credentials:      credentials.NewStaticCredentials(ak, sk, ""),
		S3ForcePathStyle: aws.Bool(true), // 必须设置为 true
	})
	if err != nil {
		return nil,utils.Errorf("Error: " + err.Error())
	}
	s3Client := s3.New(sess)
	return &S3Client{
		cli: s3Client,
		updatePath: updatePath,
		bucket: bucket,
	},nil
}

func (s *S3Client) Write(data []byte) error {
	_, err := s.cli.PutObject(&s3.PutObjectInput{
		Bucket:        aws.String(s.bucket),
		Key:           aws.String(s.updatePath),
		Body:          bytes.NewReader(data),
		CacheControl:  aws.String("no-cache, no-store, must-revalidate"),
	})
	if err != nil {
		return utils.Errorf(err.Error())
	}
	fmt.Println("Success: write file success")
	return nil
}

func (s *S3Client) Reader() ([]byte, error) {
	resp, err := s.cli.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(s.updatePath),
	})
	if err != nil {
		return nil, utils.Errorf("Error: get file failed by s3\n%w" , err)
	}
	defer resp.Body.Close()
	// 读取文件内容
	buf := new(bytes.Buffer)
	_,err = buf.ReadFrom(resp.Body)
	if err != nil {
		return nil, utils.Errorf("Error: read file failed by s3\n%w", err)
	}
	content := buf.String()
	fmt.Println("File Content:", content)
	return buf.Bytes(), nil
}
