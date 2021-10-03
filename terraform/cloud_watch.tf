resource "aws_cloudwatch_event_rule" "every_day" {
  name                = "every-day"
  description         = "Fires every one minutes"
  schedule_expression = "cron(0 11 * * *)"
}

resource "aws_cloudwatch_event_target" "daily_ping" { # Somehow figure out a way to pass an arguments
  rule      = "${aws_cloudwatch_event_rule.every_day.name}"
  target_id = "lambda"
  arn       = "${aws_lambda_function.Birthday_lambda.arn}"

  input = <<JSON
  {
      "Type":  "DailyPing",
      "SendPM": false,
  }
  JSON
}

resource "aws_lambda_permission" "allow_cloudwatch_to_daily_ping" {
  statement_id  = "AllowExecutionFromCloudWatch"
  action        = "lambda:InvokeFunction"
  function_name = "${aws_lambda_function.Birthday_lambda.function_name}"
  principal     = "events.amazonaws.com"
  source_arn    = "${aws_cloudwatch_event_rule.every_day.arn}"
}