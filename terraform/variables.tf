terraform {
    required_providers {
        aws = {
            source="hashicorp/aws"
            version="3.55.0"
        }
    }
}

variable "aws_region" {
    description = "Home AWS"
    default     = "us-east-2"
}

provider "aws" {
    region = "${var.aws_region}"
    shared_credentials_file = "~/.aws/credentials"
} 

variable "PATH"{
    type = string
} 

# Enviroinment variables
variable "SLACKBOT_TOKEN" {
    type = string
}

variable "GOOGLE_API_JSON" {
    type = string
}

variable "GOOGLE_SHEETS_ID" {
    type = string
}

variable "BUCKET_NAME"{
    type = string
}


