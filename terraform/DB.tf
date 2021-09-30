resource "aws_dynamodb_table" "birthday_message_db" {
    name           = "Bday_Messages"
    billing_mode   = "PROVISIONED"
    read_capacity  = 20
    write_capacity = 20
    hash_key       = "SlackID_Year"
    range_key      = "Message"

    attribute {
        name = "SlackID_Year"
        type = "S"
    }

    attribute {
        name = "MessageUser"
        type = "S"
    }

    attribute {
        name = "Message"
        type = "S"
    }

    ttl {
        attribute_name = "TimeToExist"
        enabled        = true
    }

    global_secondary_index {
        name               = "PrivateBdayMSG"
        hash_key           = "MessageUser"
        range_key          = "Message"
        write_capacity     = 10
        read_capacity      = 10
        projection_type    = "INCLUDE"
        non_key_attributes = ["SlackID_Year"]
    }

    tags = {
        Name        = "dynamodb-table-1"
        Environment = "production"
    }
}
