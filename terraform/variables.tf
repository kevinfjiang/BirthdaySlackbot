terraform {
    required_providers {
        aws = {
            source="hashicorp/aws"
            version="3.55.0"
        }
    }
}

variable "aws_region" {
  description = "The AWS region to create things in."
  default     = "eu-west-2"
}

variable "aws_lambda_function" {
  default = "Birthday_Messagers"
}

output "Birthdays" {
  value = "${aws_lambda_function.Birthdays.qualified_arn}"
}

variable "Path"{
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


