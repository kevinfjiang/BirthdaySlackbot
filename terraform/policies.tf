# lambda
resource "aws_iam_role" "iam_for_lambda" {
    name = "lambda_function_policies"

    assume_role_policy = "${data.aws_iam_policy_document.lambda_policy.json}"
}

data "aws_iam_policy_document" "lambda_policy" {
    statement {
        effect = "Allow"
        actions = ["sts:AssumeRole"]
        sid = ""
        principals {
            type        = "Service"
            identifiers = ["lambda.amazonaws.com"]
        }
        
    }
}

resource "aws_iam_role_policy_attachment" "lambda_policy" {
  role       = aws_iam_role.iam_for_lambda.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}


# AWS ID
data "aws_caller_identity" "current" {}
locals {
    account_id = data.aws_caller_identity.current.account_id
}

#DB access
resource "aws_iam_role_policy" "lambda_policy_extension" {
    name = "DB_Access_lambda_policy"
    role = aws_iam_role.iam_for_lambda.id

    policy = "${data.aws_iam_policy_document.DB_Policy.json}"
}

data "aws_iam_policy_document" "DB_Policy" {
    statement {
        effect = "Allow"
        actions = [
        "dynamodb:BatchGetItem",
        "dynamodb:GetItem",
        "dynamodb:Query",
        "dynamodb:Scan",
        "dynamodb:BatchWriteItem",
        "dynamodb:PutItem",
        "dynamodb:UpdateItem"
        ]

        resources = [
        "arn:aws:dynamodb:${var.aws_region}:${local.account_id}:table/&{aws:username}",
        ]
    }
}


# Cloudwatch
resource "aws_lambda_permission" "allow_cloudwatch_to_daily_ping" { #TODO set up other the json to be sent
    statement_id  = "AllowExecutionFromCloudWatch"
    action        = "lambda:InvokeFunction"
    function_name = "${aws_lambda_function.birthday_lambda.function_name}"
    principal     = "events.amazonaws.com"
    source_arn    = "${aws_cloudwatch_event_rule.every_day.arn}"
}