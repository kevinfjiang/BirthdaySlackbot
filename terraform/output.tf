output "base_url" {
  description = "Base URL for API Gateway stage."

  value = aws_api_gateway_deployment.gateway_deploy.invoke_url
}