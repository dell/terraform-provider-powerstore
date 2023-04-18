variable "username" {
  type=string
  description = "Stores the username of PowerStore host."
}

variable "password" {
  type=string
  description = "Stores the password of PowerStore host."
}

variable "timeout" {
  type=string
  description = "Stores the timeout of PowerStore host."
}

variable "endpoint" {
  type=string
  description = "Stores the endpoint of PowerStore host. eg: https://10.1.1.1/api/rest/"
}
