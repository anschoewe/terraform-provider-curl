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