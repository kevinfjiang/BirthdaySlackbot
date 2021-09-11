terraform {
    required_providers {
        aws = {
            source="hashicorp/aws"
            version="3.55.0"
        }
    }
}

provider "aws" {} // Include more stuff

resource "aws_dynamodb_table" "birthday_message_db" {
    name           = "bday_messages"
    billing_mode   = "PROVISIONED"
    read_capacity  = 20
    write_capacity = 20
    hash_key       = "SlackID_Year"
    range_key      = "Messages"

    attribute {
        name = "SlackID_Year"
        type = "S"
    }

    attribute {
        name = "MessageUser"
        type = "S"
    }

    attribute {
        name = "Message"
        type = "S"
    }

    ttl {
        attribute_name = "TimeToExist"
        enabled        = false
    }

    global_secondary_index {
        name               = "PrivateBdayMSG"
        hash_key           = "MessageUser"
        range_key          = "Message"
        write_capacity     = 10
        read_capacity      = 10
        projection_type    = "KEYS_ONLY"
        non_key_attributes = ["SlackID_Year"]
    }

    tags = {
        Name        = "dynamodb-table-1"
        Environment = "production"
    }
}

// TODO Make sure to set up the lambda stuff, as this is copied frm website
resource "aws_iam_role" "iam_for_lambda" {
  name = "iam_for_lambda"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_lambda_function" "test_lambda" {
  filename      = "lambda_function_payload.zip"
  function_name = "lambda_function_name"
  role          = aws_iam_role.iam_for_lambda.arn
  handler       = "index.test"

  # The filebase64sha256() function is available in Terraform 0.11.12 and later
  # For Terraform 0.11.11 and earlier, use the base64sha256() function and the file() function:
  # source_code_hash = "${base64sha256(file("lambda_function_payload.zip"))}"
  source_code_hash = filebase64sha256("lambda_function_payload.zip")

  runtime = "nodejs12.x"

  environment {
    variables = {
      foo = "bar"
    }
  }
}

