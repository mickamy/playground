package config

import (
	"github.com/caarlos0/env/v11"
)

type AWSConfig struct {
	Region             string `env:"AWS_REGION" envDefault:"ap-northeast-1"`
	S3Bucket           string `env:"S3_BUCKET"`
	AccessKey          string `env:"AWS_ACCESS_KEY_ID"`
	SecretKey          string `env:"AWS_SECRET_ACCESS_KEY"`
	LocalstackEndpoint string `env:"LOCALSTACK_ENDPOINT"`
}

var aws = AWSConfig{}

func init() {
	if err := env.Parse(&aws); err != nil {
		panic(err)
	}
}

func AWS() AWSConfig {
	return aws
}
