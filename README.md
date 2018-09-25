# Terrabender

This program is intended to use github as a source.

It should handle the merging of pull request with your main branch (default to master)

How does the release process works:

- Have a terraform (usually) git repository configure
- give permission to terrabender to access terraform git repo, merge pull request, post comment.

For every pull request create, terrabender will do a command, ie terraform plan.

Exemple
