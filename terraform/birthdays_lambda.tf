#Create security group with firewall rules
data "archive_file" "zip" {
  type        = "zip"
  source_file = var.PATH
  output_path = "BirthdayMSG.zip"
}
resource "aws_cloudwatch_log_group" "log" {
  name              = "/aws/lambda/${var.aws_lambda_function}"
  retention_in_days = 14
}

resource "aws_lambda_function" "birthday_lambda" {
  function_name = "${var.aws_lambda_function}"

  s3_bucket = aws_s3_bucket.lambda_bucket.id
  s3_key    = aws_s3_bucket_object.lambda_birthdays.key
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

resource "aws_lambda_permission" "birthday_permission" {
  statement_id  = "AllowAPIGatewayInvoke"
  action        = "lambda:InvokeFunction"
  function_name = "${aws_lambda_function.birthday_lambda.function_name}"
  principal     = "apigateway.amazonaws.com"

  # The /*/* portion grants access from any method on any resource
  # within the API Gateway "REST API".
  source_arn = "${aws_api_gateway_rest_api.BdayGateway.execution_arn}/*/*"
}
