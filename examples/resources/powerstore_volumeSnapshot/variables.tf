variable "username" {
  type=string
  description = "Stores the username of PowerStore volume snapshot."
}

variable "password" {
  type=string
  description = "Stores the password of PowerStore volume snapshot."
}

variable "timeout" {
  type=string
  description = "Stores the timeout of PowerStore volume snapshot."
}

variable "endpoint" {
  type=string
  description = "Stores the endpoint of PowerStore volume snapshot. eg: https://10.1.1.1:443/api/rest/ where 443 is port where API requests are getting accepted"
}
