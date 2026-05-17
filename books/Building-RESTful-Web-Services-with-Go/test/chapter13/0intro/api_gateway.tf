#  New API on Amazon API Gateway
resource "aws_api_gateway_rest_api" "test" {
  name = "EC2Example"  # API 名称
  description = "Terraform EC2 REST API Example" # 关于
  endpoint_configuration { # 定义要发布的API的模式：REGIONAL、EDGE、PRIVATE
    types = ["REGIONAL"]
  }
}

# 网关的方法请求 Method request congiguration
resource "aws_api_gateway_method" "test" {
  authorization = "NONE"
  http_method   = "GET"
  resource_id   = aws_api_gateway_rest_api.test.root_resource_id
  rest_api_id   = aws_api_gateway_rest_api.test.id
}

# 网关的方法响应 Method response configuration
resource "aws_api_gateway_method_response" "test" {
  http_method = aws_api_gateway_method.test.http_method
  resource_id = aws_api_gateway_rest_api.test.root_resource_id
  rest_api_id = aws_api_gateway_rest_api.test.id
  status_code = "200"
}

// 网关的集成请求 Integration request configuration
resource "aws_api_gateway_integration" "test" {
  rest_api_id = aws_api_gateway_rest_api.test.id
  resource_id = aws_api_gateway_method.test.resource_id
  http_method = aws_api_gateway_method.test.http_method

  integration_http_method = "GET"
  type                    = "HTTP"
  uri                     = "http://${aws_instance.api_server.public_dns}/api/books"
}


# 网关的集成响应 Integration response configuration
resource "aws_api_gateway_integration_response" "MyDemoIntegrationResponse" {
  rest_api_id = aws_api_gateway_rest_api.test.id
  resource_id = aws_api_gateway_rest_api.test.root_resource_id
  http_method = aws_api_gateway_method.test.http_method

  status_code = aws_api_gateway_method_response.test.status_code
}

// 部署网关 Deploy API on Gateway with test environment
resource "aws_api_gateway_deployment" "test" {
  depends_on = [
    aws_api_gateway_integration.test
  ]

  rest_api_id = aws_api_gateway_rest_api.test.id
  stage_name  = "test"
}
