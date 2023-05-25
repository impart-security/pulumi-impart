
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

    $ go get github.com/impart-security/pulumi-impart

## Configuration

- Login to https://console.impartsecurity.net/. Under manage section click Settings => Access Tokens => New API access token. Select scopes: read:org_api_access_tokens and scopes for resources will be managed by pulumi.
- Set environment variable `IMPART_TOKEN` or spicify `token` argument as part of provider configuration in the code

