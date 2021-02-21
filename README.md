![GitHub Workflow Status](https://img.shields.io/github/workflow/status/anschoewe/terraform-provider-curl/release)
![GitHub](https://img.shields.io/github/license/anschoewe/terraform-provider-curl)

# Overview
This is a simple Terraform provider that provides curl-like functionality.  In addition to basic HTTP methods, it also supports OAuth2 tokens.  It's currently hardcoded for Azure AD M2M tokens, but it could be extended in the future to support additional token issuers.  The `examples/` folder shows how to use the provider.

# Make
```
make install
```

# Test
```
cd examples/
rm -fr .terraform/ .terraform.lock.hcl && terraform init && terraform apply --auto-approve
```

# Release
Releases are automatically created when a tag with the semantic versioning format `vX.Y.Z` (ex `v1.0.0`) is pushed to the GitHub repo.
There is a pipeline/action that will create builds for multiple architectures.  Additionally, signing keys have been uploaded and configured with registry.terraform.io.  This means the release will be avaiable in GitHub but also published to the official Terraform Registry.
```
# Create release tag (example)
git tag v0.1.0
# Push release tag to GitHub and start build + release
git push origin v0.1.0
```

# Plugin Authors
Here is how you can generate the documentation
```
go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs
```