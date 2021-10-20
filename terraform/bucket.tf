#TODO set up CD CI with s3 buckets
#TODO project is coming together, all I have left is set up the server and linking the aws endpoint
#TODO configure the slack function and ensure better integration in golang code
data "archive_file" "zip" {
    type        = "zip"
    source_file = var.PATH
    output_path = "main.zip"
}

resource "aws_s3_bucket" "lambda_bucket" {
    bucket = var.BUCKET_NAME
    acl    = "private"

    tags = {
            Name        = "lambda_bucket"
            Environment = "Dev"
        }
    }
    resource "aws_s3_bucket_object" "birthday_source" {
    bucket = aws_s3_bucket.lambda_bucket.id
    key    = "birthday_source_code"
    acl    = "private"
    source = "main.zip"
    etag = "main.zip"
}
