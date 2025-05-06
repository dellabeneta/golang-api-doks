
terraform {
  backend "s3" {
    bucket         = "terraformstate-terraform-golang-api-doks"
    key            = "terraformstate"
    region         = "sa-east-1"
    encrypt        = true
    dynamodb_table = "terraform-lock-golang-api-doks"
  }
}

