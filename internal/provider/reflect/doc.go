// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

// Package reflect is a forked version of https://github.com/hashicorp/terraform-plugin-framework/tree/main/internal/reflect
// that has been modified to support speakeasy's terraform generator.
// In particular, behaviour differs in that it is intended to support merging Terraform State and Terraform Plan structures
// into a single data structure, with Known Plan values overriding State values. This allows for code to be written
// that drives API calls from a single point of truth.
// Fork Commit hash is 99f28445b60580b6e39afda88a4bb469461f9bbb
package reflect
