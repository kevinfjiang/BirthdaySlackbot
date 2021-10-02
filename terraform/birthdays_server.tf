#Create security group with firewall rules
provider "aws" {
    region = "${var.aws_region}"
    shared_credentials_file = "~/.aws/credentials"
} 

data "archive_file" "zip" {
  type        = "zip"
  source_file = var.Path
  output_path = "BirthdayMSG.zip"
}
resource "aws_cloudwatch_log_group" "log" {
  name              = "/aws/lambda/${var.aws_lambda_function}"
  retention_in_days = 14
}

resource "aws_iam_role" "iam_for_lambda" {
  name               = "iam_for_lambda"
  assume_role_policy = "${data.aws_iam_policy_document.policy.json}"
}

data "aws_iam_policy_document" "policy" {
  statement {
    sid    = ""
    effect = "Allow"

    principals {
      identifiers = ["lambda.amazonaws.com"]
      type        = "Service"
    }

    actions = ["sts:AssumeRole"]
  }
}

resource "aws_lambda_function" "Birthdays" {
  function_name = "${var.aws_lambda_function}"

  filename         = "${data.archive_file.zip.output_path}"
  source_code_hash = "${data.archive_file.zip.output_base64sha256}"

  role    = "${aws_iam_role.iam_for_lambda.arn}"
  handler = "BirthdayFinder.lambda_handler"
  runtime = "go1.x"

  environment {
    variables = {
      SLACKBOT_TOKEN   = var.SLACKBOT_TOKEN
      GOOGLE_API_JSON  = var.GOOGLE_API_JSON
      GOOGLE_SHEETS_ID = var.GOOGLE_SHEETS_ID
    }
  }

  tags = {
        Name        = "lambda-birthday"
        Environment = "production"
    }
}