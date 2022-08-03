# ⚠️⚠️ This is not official Terraform provider for Vercel. For the official provider go to: [vercel/terraform-provider-vercel](https://github.com/vercel/terraform-provider-vercel) ⚠️⚠️

> This repository is in **Work in Progress** state. If you need something, create an [issue](https://github.com/ondrejsika/terraform-provider-vercel/issues/new)

# ondrejsika/terraform-provider-vercel

    2019 Ondrej Sika <ondrej@ondrejsika.com>
    https://github.com/ondrejsika/terraform-provider-vercel

![Build](https://github.com/ondrejsika/terraform-provider-vercel/workflows/Build/badge.svg)

## My Related Projects

- [ondrejsika/vercel-go](https://github.com/ondrejsika/vercel-go) - Go client for Vercel API
- [ondrejsika/vercel-api-mock](https://github.com/ondrejsika/vercel-api-mock) - Vercel API Mock

## Buy Domain on Vercel using Terraform

![Buy Domain on Vercel using Terraform](buy-domain-on-zeit-using-terraform.png)

## Example usage

```terraform
provider "vercel" {
  token = "secret-token"
  // Optional
  // api_origin = "https://vercel-api-mock.sikademo.com"
}

resource "vercel_domain" "sikademovercel_com" {
  domain = "sikademovercel.com"
  expected_price = 12
}

resource "vercel_dns" "sikademovercel_com" {
  domain = vercel_domain.sikademovercel_com.domain
  name = ""
  value = "1.2.3.4"
  type = "A"
}

resource "vercel_dns" "www_sikademovercel_com" {
  domain = vercel_domain.sikademovercel_com.domain
  name = "www"
  value = "sikademovercel.com."
  type = "CNAME"
}

resource "vercel_dns" "mail_sikademovercel_com" {
  domain = vercel_domain.sikademovercel_com.domain
  name = "mail"
  value = "5.6.7.8"
  type = "A"
}

resource "vercel_dns" "mx_sikademovercel_com" {
  domain = vercel_domain.sikademovercel_com.domain
  name = ""
  value = "99 mail.sikademovercel.com."
  type = "MX"
}

resource "vercel_project" "demo" {
  name = "sika-demo-vercel"
}
```

## Change Log

### v2.1.0

- Add `vercel_domain` importer

### v2.0.1

- Add Go Realaser config

### v2.0.0

- Change ZEIT to Vercel (terraform-provider-vercel, resource names)

### v1.3.2

- Fix error handing of errors from `ondrejsika/zeit-go` API client
- Handle buy of unavailable domains

### v1.3.1

- Update `ondrejsika/zeit-go` for `/v4/domain/buy` API

### v1.3.0

- Add parameter `remove_domain_on_destroy` with default `false` to `zeit_domain`. When you call `terraform destroy` domain will be kept on Zeit if you not set `remove_domain_on_destroy=true`
- Rewrite for [ondrejsika/zeit-go](https://github.com/ondrejsika/zeit-go)

### v1.2.0

- Add resource `zeit_domain` for buy domains on Zeit

### v1.1.0

- Add `api_origin` configuration for provider

### v1.0.0

- Create provider `zeit`
- Add resource `zeit_dns` with minimum configuration
