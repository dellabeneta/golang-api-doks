
variable "do_token" {
  description = "DigitalOcean API token"
  type        = string
}

variable "doks_region" {
  description = "Region for the DOKS cluster"
  type        = string
  default     = "sfo2"
}

variable "doks_version" {
  description = "Version of the DOKS cluster"
  type        = string
  default     = "1.32.2-do.0"
}