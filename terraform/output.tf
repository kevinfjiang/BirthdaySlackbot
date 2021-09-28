output "lambda" {
  value = "${aws_lambda_function.MSGBotlambda.qualified_arn}"
}