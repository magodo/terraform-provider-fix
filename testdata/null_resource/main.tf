locals {
  foo = "bar"
}

resource "null_resource" "test" {
  triggers = {
    foo = local.foo
  }
}

output "triggers" {
  value = null_resource.test.triggers
}
