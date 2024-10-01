# terraform-provider-fix

This provider is a template for [`terrafix`](https://github.com/magodo/terrafix) users, to conduct adhoc fixes on the Terraform configurations. It only defines two provider functions: `terrafix_config_definition` and `terrafix_config_references`. The user are supposed to fork this repository, and implement their own definition/reference config fix logics under *internal/<resource|datasource>/<def|ref>*.

## Example

There is an example that demonstrates how to change the definition and referene of `null_resource` for renaming its attribute `trggers` to `the_triggers`, at *internal/resource/<def|ref>/null_resource.go*.

To run it:

```shell
$ cd <repo root dir>

# Install the provider
$ go install

# Run terrafix 
$ terrafix -provider-path ~/go/bin/terraform-provider-fix -provider-addr registry.terraform.io/hashicorp/null ./testdata/null_resource
Path: testdata/null_resource/main.tf

locals {
  foo = "bar"
}

resource "null_resource" "test" {
  the_triggers = {
    foo = local.foo
  }
}

output "triggers" {
  value = null_resource.test.the_triggers
}
```
