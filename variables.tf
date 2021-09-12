terraform {
    required_providers {
        aws = {
            source="hashicorp/aws"
            version="3.55.0"
        }
    }
}

provider "aws" {
    region = "${var.aws_region}"
    shared_credentials_file = "~/.aws/credentials"
} 


output "lambda" {
  value = "${aws_lambda_function.lambda.qualified_arn}"
}

variable "aws_region" {
  description = "The AWS region to create into"
  default     = "us-east-1"
}