# Roadmap SH Backend by FerranRT

## Scripts

### Upgrade Go version of all packages

```
# Interactive mode
./scripts/upgrade-go-version.sh

# Command-line mode
./scripts/upgrade-go-version.sh 1.25.0
```

And outputs:

```
=== Go Version Upgrader ===
Upgrading Go version to: 1.25.0

Found 2 go.mod file(s):
  - projects/github-user-activity/go.mod
  - projects/my-api/go.mod

Processing: projects/github-user-activity/go.mod
  Current version: 1.24.0
  ✅ Updated to version 1.25.0

Processing: projects/my-api/go.mod
  Current version: 1.25.0
  ✅ Already at version 1.25.0

=== Summary ===
Total go.mod files processed: 2
Files updated: 1

✅ Go version upgrade completed successfully!

Next steps:
1. Review the changes in the go.mod files
2. Run 'go mod tidy' in each project directory to update dependencies
3. Test your projects to ensure compatibility
```

## Projects

| Project              | URL                                              |
| -------------------- | ------------------------------------------------ |
| Github User Activity | https://roadmap.sh/projects/github-user-activity |
|                      |                                                  |
