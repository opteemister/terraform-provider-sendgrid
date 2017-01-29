# terraform-provider-sendgrid

Sendgrid provider plugin for Terraform.

## Installation

### Requirements

1. [Terraform](http://terraform.io). Make sure you have it installed and it's accessible from your `$PATH`.
2. Sendgrid

### From Source (only method right now)

* Install Go and [configure](https://golang.org/doc/code.html) your workspace.

* Download this repo:

```shell
$ go get github.com/opteemister/terraform-provider-sendgrid
```

* Install the dependencies:

```shell
$ cd $GOPATH/src/github.com/jtopjian/terraform-provider-sendgrid
$ go get
```

* Compile it:

```shell
$ go build -o terraform-provider-sendgrid
```

* Copy it to a directory:

```shell
$ sudo cp terraform-provider-sendgrid ~/.terraform/providers/
```

* Define new provider in terraform configuration:

```Edit file
~/.terraform/terraformrc
```
```Add lines
providers {
    sendgrid = "~/.terraform/providers/terraform-provider-sendgrid"
}
```

## Usage

Here's a simple Terraform file to get you started:

```ruby
provider "sendgrid" {
}

resource "sendgrid_template" "first_template" {
  name = "name"
}

resource "sendgrid_template_version" "first_template_version" {
  name = "version name"
	template_id = "${sendgrid_template.first_template.id}"
	subject = "email subject"
	html_content_file = "./resources/test_template.html"
	plain_content_file = "./resources/test_template_plain.html"
  active = true
}
```

For either example, save it to a `.tf` file and run:

```shell
$ terraform plan
$ terraform apply
$ terraform show
```

## Reference

### Provider

#### Example

```ruby
provider "sendgrid" {
  apiKey = "sendgrid_key"
}
```

#### Parameters

* `apiKey`: Optional. Set the key for sendgrid account. You can set it through the ENV_VARS

### sendgrid_template

#### Example

```ruby
resource "sendgrid_template" "my_template" {
  name = "my_template"
}
```

#### Parameters

* `name`: Required. The name of the template.

#### Exported Parameters

* `id`: The Id of the new template.

### sendgrid_template_version

#### Example

```ruby
resource "sendgrid_template_version" "my_template_version" {
  name = "version name"
	template_id = "${sendgrid_template.first_template.id}"
	subject = "email subject"
	html_content_file = "./resources/test_template.html"
	plain_content_file = "./resources/test_template_plain.html"
  active = true
}
```

#### Parameters

* `name`: Required. The name of the template_version.
* `template_id`: Required. The id of the template.
* `subject`: Required. The subject for the email template.
* `html_content_file`: Required. Html content file path.
* `plain_content_file`: Required. Plain text file path.
* `active`: Optional. Boolean option if template version is active. Default true.