#Create security group with firewall rules
resource "aws_cloudwatch_log_group" "log" {
    name              = "/aws/lambda/${aws_lambda_function.birthday_lambda.function_name}"
    retention_in_days = 14
}

resource "aws_lambda_function" "birthday_lambda" {
    function_name = "Birthday_Message_Lambda"

    s3_bucket = aws_s3_bucket.lambda_bucket.id
    s3_key    = aws_s3_bucket_object.birthday_source.key
    source_code_hash = "${data.archive_file.zip.output_base64sha256}"

    role    = "${aws_iam_role.iam_for_lambda.arn}"
    handler = "main"
    runtime = "go1.x"

    environment {
    variables = {
        SLACKBOT_TOKEN   = var.SLACKBOT_TOKEN
        GOOGLE_API_JSON  = var.GOOGLE_API_JSON
        GOOGLE_SHEETS_ID = var.GOOGLE_SHEETS_ID
    }
  }

    tags = {
        Name        = "lambda_birthday"
        Environment = "production"
    }
}

