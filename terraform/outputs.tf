output bucket_name {
  value       = aws_s3_bucket.bucket.id
  description = "The name of the bucket"
}

output sqs_queue_url {
  value       = aws_sqs_queue.queue.url
  description = "The URL of the SQS queue"
}

output sqs_queue_arn {
  value       = aws_sqs_queue.queue.arn
  description = "The ARN of the SQS queue"
}
