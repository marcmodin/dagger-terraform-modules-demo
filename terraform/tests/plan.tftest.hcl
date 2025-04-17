
provider "aws" {
  access_key                  = "test"
  secret_key                  = "test"
  region                      = "us-east-1"
  s3_use_path_style           = true
  skip_credentials_validation = true
  skip_metadata_api_check     = true
  skip_requesting_account_id  = true

  # endpoints {
  #   apigateway     = "http://localstack:4566"
  #   apigatewayv2   = "http://localstack:4566"
  #   cloudformation = "http://localstack:4566"
  #   cloudwatch     = "http://localstack:4566"
  #   dynamodb       = "http://localstack:4566"
  #   ec2            = "http://localstack:4566"
  #   es             = "http://localstack:4566"
  #   elasticache    = "http://localstack:4566"
  #   firehose       = "http://localstack:4566"
  #   iam            = "http://localstack:4566"
  #   kinesis        = "http://localstack:4566"
  #   lambda         = "http://localstack:4566"
  #   rds            = "http://localstack:4566"
  #   redshift       = "http://localstack:4566"
  #   route53        = "http://localstack:4566"
  #   s3             = "http://localstack:4566"
  #   s3control      = "http://localstack:4566"
  #   secretsmanager = "http://localstack:4566"
  #   ses            = "http://localstack:4566"
  #   sns            = "http://localstack:4566"
  #   sqs            = "http://localstack:4566"
  #   ssm            = "http://localstack:4566"
  #   stepfunctions  = "http://localstack:4566"
  #   sts            = "http://localstack:4566"
  # }
}

run "valid_string_concat" {

  command = apply

  variables {
    bucket_name = "test-bucket-23ma0s8fadf8"
    queue_name  = "test-queue-23ma0s8fadf8"
  }

  // assert that the bucket id is the same as the bucket name
  assert {
    condition     = aws_s3_bucket.bucket.bucket == "test-bucket-23ma0s8fadf8"
    error_message = "S3 bucket name did not match expected"
  }

  // assert that the queue id is the same as the queue name
  assert {
    condition     = aws_sqs_queue.queue.name == "test-queue-23ma0s8fadf8"
    error_message = "SQS queue name did not match expected"
  }

}