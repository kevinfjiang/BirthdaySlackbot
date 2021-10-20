resource "aws_cloudwatch_event_rule" "every_day" {
    name                = "every-day"
    description         = "Fires every day"
    schedule_expression = "cron(0 11 * * ? *)"
}

resource "aws_cloudwatch_event_target" "daily_ping" { # TODO Somehow figure out a way to pass in  more arguments
    rule      = "${aws_cloudwatch_event_rule.every_day.name}"
    target_id = "birthday_lambda"
    arn       = "${aws_lambda_function.birthday_lambda.arn}"

    input = <<JSON
    {
        "Type":  "DailyPing",
        "SendPM": false
    }
    JSON
}