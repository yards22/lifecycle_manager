package objectstore

import (
	"errors"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var ErrObjectNotFound = errors.New("object not found")

type ObjectStore interface {
	Put(data io.Reader, name string) (string, error)
	Get(name string) (io.ReadCloser, error)
	Delete(name string) error
}

type S3 struct {
	region   string
	bucket   string
	session  *session.Session
	uploader *s3manager.Uploader
	svc      *s3.S3
}

func New(AWSAccessID, AWSSecret, Region, Bucket string) (*S3, error) {
	s := &S3{
		region: Region,
		bucket: Bucket,
	}
	session, err := session.NewSession(
		&aws.Config{
			Region: aws.String(Region),
			Credentials: credentials.NewStaticCredentials(
				AWSAccessID,
				AWSSecret,
				"",
			),
		})
	if err != nil {
		return s, err
	}
	s.session = session
	s.uploader = s3manager.NewUploader(session)
	s.svc = s3.New(session)
	return s, nil
}

// uploads the object to S3 and returns the location
func (s *S3) Put(data io.Reader, name string) (string, error) {
	up, err := s.uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(name),
		Body:   data,
	})
	if err != nil {
		return "", err
	}
	return up.Location, nil
}

// get the object from S3
func (s *S3) Get(name string) (io.ReadCloser, error) {
	res, err := s.svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(name),
	})
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}

// delete the object from S3
func (s *S3) Delete(name string) error {
	_, err := s.svc.DeleteObject(&s3.DeleteObjectInput{Bucket: aws.String(s.bucket), Key: aws.String(name)})
	if err != nil {
		return err
	}
	return nil
}
