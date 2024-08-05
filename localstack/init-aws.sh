#!/bin/bash

awslocal s3api create-bucket \
  --bucket $S3_BUCKET \
  --region ap-northeast-1 \
  --create-bucket-configuration LocationConstraint=ap-northeast-1
