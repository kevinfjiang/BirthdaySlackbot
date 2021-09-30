#Create security group with firewall rules
resource "aws_security_group" "generic-server" {
  name        = var.security_group
  description = "security group for server"

  ingress {
    from_port   = 8080
    to_port     = 8080
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

 # outbound from server
  egress {
    from_port   = 0
    to_port     = 65535
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags= {
    Name = var.security_group
  }
}

resource "aws_instance" "kevin_server" {
  ami           = var.ami_id
  key_name      = var.key_name
  instance_type = var.instance_type
  tags= {
    Name = var.tag_name
  }
}

# Create Elastic IP address
resource "aws_eip" "kevin_server" {
  vpc      = true
  instance = aws_instance.kevin_server.id
tags= {
    Name = "kevins_elastic_ip"
  }
}