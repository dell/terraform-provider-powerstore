variable "username" {
  type = string
}

variable "password" {
  type = string
}

variable "endpoint" {
  type = string
}

provider "powerstore" {
  username = var.username
  password = var.password
  endpoint = var.endpoint
  insecure = true
}