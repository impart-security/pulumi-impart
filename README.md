
# Impart Resource Provider

The Impart Resource Provider lets you configure and manage Impart resources.

## Installing

This package is available in many languages in the standard packaging formats.

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

    $ npm install @impart-security/pulumi-impart

or `yarn`:

    $ yarn add  @impart-security/pulumi-impart

### Go

To use from Go, use `go get` to fetch the latest version of the library

    $ go get github.com/impart-security/pulumi-impart/sdk/go/impart

## Configuration

- Login to [https://console.impartsecurity.net](https://console.impartsecurity.net). Under the `Manage` section click `Settings => Access Tokens => New API access token`. Select the scope `read:org_api_access_tokens` along with scopes for resources that will be managed by Pulumi.
- Set environment variable `IMPART_TOKEN` or specify a `token` argument as part of provider configuration in the code

