package minio

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
)

type Conf struct {
	Endpoint  string `json:"endpoint,omitempty" yaml:"endpoint"`
	AccessKey string `json:"access_key,omitempty" yaml:"access_key"`
	SecretKey string `json:"secret_key,omitempty" yaml:"secret_key"`
	UseSSL    bool   `json:"use_ssl,omitempty" yaml:"use_ssl"`
}

func InitMinio(c *Conf) *minio.Client {
	client, err := minio.New(c.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(c.AccessKey, c.SecretKey, ""),
		Secure: c.UseSSL,
	})
	if err != nil {
		log.Fatalf("minio connect error: %v\n", err)
	}
	return client
}
