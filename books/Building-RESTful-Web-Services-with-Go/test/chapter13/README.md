# Terraform
> https://developer.hashicorp.com/terraform/language

resource "aws_vpc" "main" {
cidr_block = var.base_cidr_block
}

<BLOCK TYPE> "<BLOCK LABEL>" "<BLOCK LABEL>" {
  # Block body
  <IDENTIFIER> = <EXPRESSION> # Argument
}

Terraform 脚本由四个基本构建块组成：
* 块类型(BLOCK TYPE)：Terraform 预定义的一组块类型——例如，资源和数据。
* 块标签(BLOCK LABEL)：Terraform 脚本中块类型的命名空间+实例。起标识作用，当前块的全局唯一。
* 标识符(IDENTIFIER)：块内的变量。
* 表达式(EXPRESSION)：块内变量的值

```bash
terraform init
terraform plan
terraform apply [-auto-approve]
terraform destroy
terraform output
terraform state list
terraform state show aws_instance.web
terraform state rm aws_instance.web
terraform state pull > terraform.tfstate
terraform state push terraform.tfstate
terraform state mv aws_instance.web aws_instance.web2

# 公钥加密；私钥解密；
ssh-keygen -t rsa -b 4096


```

# AMI
> https://cloud-images.ubuntu.com/locator/ec2/
