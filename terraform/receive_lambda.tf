#Create security group with firewall rules
resource "aws_cloudwatch_log_group" "log_receive" {
    name              = "/aws/lambda/${aws_lambda_function.receive_lambda.function_name}"
    retention_in_days = 14
}

resource "aws_lambda_function" "receive_lambda" {
    function_name = "Receive_Message_Lambda"

    s3_bucket = aws_s3_bucket.lambda_bucket.id
    s3_key    = aws_s3_bucket_object.receive_source.key
    source_code_hash = "${data.archive_file.receive.output_base64sha256}"

    role    = "${aws_iam_role.iam_for_lambda.arn}"
    handler = "main"
    runtime = "go1.x"

    environment {
        variables = {
            SLACKBOT_TOKEN   = var.SLACKBOT_TOKEN
        }
  }

    tags = {
        Name        = "lambda_birthday"
        Environment = "production"
    }
}

resource "aws_lambda_permission" "api_gw" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.receive_lambda.function_name
  principal     = "apigateway.amazonaws.com"

  source_arn = "${aws_apigatewayv2_api.lambda.execution_arn}/*/*"
}

