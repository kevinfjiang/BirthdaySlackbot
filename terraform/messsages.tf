
# data "archive_file" "saveMSGzip" {
#   type        = "zip"
#   source_file = "save_MSG"
#   output_path = "save_MSG.zip"
# }
# // Here creates function
# resource "aws_lambda_function" "saveMSGlambda" {
#   function_name = "saveMSGlambda"

#   filename         = "${data.archive_file.zip.output_path}"
#   source_code_hash = "${data.archive_file.zip.output_base64sha256}"

#   role    = "${aws_iam_role.iam_for_lambda.arn}"
#   handler = "save_MSG"
#   runtime = "go1.x"

# #   environment {
# #     variables = {
# #       greeting = "Hello"
# #     }
# #   }
# }
