#TODO set up CD CI with s3 buckets
#TODO project is coming together, all I have left is set up the server and linking the aws endpoint
#TODO configure the slack function and ensure better integration in golang code
resource "aws_s3_bucket" "lambda_bucket" {
    bucket = var.BUCKET_NAME
    acl    = "private"

    tags = {
            Name        = "lambda_bucket"
            Environment = "Dev"
        }
    }

data "archive_file" "birthday" {
    type        = "zip"
    source_file = "${var.PATH}/birthday/main"
    output_path = "${var.PATH}/birthday/main.zip"
}

data "archive_file" "receive" {
    type        = "zip"
    source_file = "${var.PATH}/receive/main"
    output_path = "${var.PATH}/receive/main.zip"
}
resource "aws_s3_bucket_object" "birthday_source" {
    bucket = aws_s3_bucket.lambda_bucket.id
    key    = "birthday_source_code"
    acl    = "private"
    source = "${var.PATH}/birthday/main.zip"
    etag = "${var.PATH}/birthday/main.zip"
}

resource "aws_s3_bucket_object" "receive_source" {
    bucket = aws_s3_bucket.lambda_bucket.id
    key    = "receive_source_code"
    acl    = "private"
    source = "${var.PATH}/receive/main.zip"
    etag = "${var.PATH}/receive/main.zip"
}
