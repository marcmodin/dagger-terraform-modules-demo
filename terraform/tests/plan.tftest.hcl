
provider "aws" {
  access_key                  = "dummy-key"
  secret_key                  = "dummy-secret"
  region                      = "us-east-1"
  s3_use_path_style           = true
  skip_credentials_validation = true
  skip_metadata_api_check     = true
  skip_requesting_account_id  = true
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