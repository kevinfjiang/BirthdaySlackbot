provider "archive" {}

data "archive_file" "sendMSGzip" {
  type        = "zip"
  source_file = "MSGBot/MSGBot"
  output_path = "MSGBot.zip"
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

resource "aws_iam_role" "iam_for_lambda" {
  name               = "iam_for_lambda"
  assume_role_policy = "${data.aws_iam_policy_document.policy.json}"
}

// Here creates function
resource "aws_lambda_function" "MSGBotlambda" {
  function_name = "MSGBotlambda"

  filename         = "${data.archive_file.zip.output_path}"
  source_code_hash = "${data.archive_file.zip.output_base64sha256}"

  role    = "${aws_iam_role.iam_for_lambda.arn}"
  handler = "MSGBot"
  runtime = "go1.x"

#   environment {
#     variables = {
#       greeting = "Hello"
#     }
#   }
}
