# 定义了要使用的云提供商类型，并配置了安全凭据和区域
provider "aws" {
  profile = "default"
  region = "eu-central-1"
}

# 定义要提供的资源类型及其属性
resource "aws_instance" "api_server" {
  # Amazon Machine Image
  ami = "ami-0c55b159cbfafe1f4"
  instance_type = "t2.micro"
  key_name = aws_key_pair.api_server_key.key_name # 引用之前创建的密钥对
}

# 在EC2上添加公钥：~/.ssh/authorized_keys
resource "aws_key_pair" "api_server_key" {
  key_name = "api-server-key"
  public_key = "ssh-rsa ABCD...XYZ naren@Narens-MacBook-Air.local"
}

