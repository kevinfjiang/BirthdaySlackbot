terraform {
    required_providers {
        aws = {
            source="hashicorp/aws"
            version="3.55.0"
        }
    }
}


variable "aws_region" {
  description = "The AWS region to create into"
  default     = "us-east-2"
}

provider "aws" {
    region = "${var.aws_region}"
    shared_credentials_file = "~/.aws/credentials"
} 
#t2.micro AWS
variable "key_name" { 
    description = " SSH keys to connect to ec2 instance" 
    default     =  "kevins_2018Mac_key" 
}

variable "instance_type" { 
    description = "instance type for ec2" 
    default     =  "t2.micro" 
}

variable "security_group" { 
    description = "Name of security group" 
    default     = "generic-server" 
}

variable "tag_name" { 
    description = "Tag Name of for Ec2 instance" 
    default     = "ec2-instance" 
} 
variable "ami_id" { 
    description = "AMI for Ubuntu Ec2 instance" 
    default     = "ami-00399ec92321828f5" 
}



