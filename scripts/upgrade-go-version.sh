#!/bin/bash

# Script to upgrade Go version in all project go.mod files
# Usage: ./upgrade-go-version.sh [new_version]
# Example: ./upgrade-go-version.sh 1.25.0

echo "=== Go Version Upgrader ==="
echo ""

# Check if projects directory exists
if [ ! -d "projects" ]; then
    echo "Error: projects directory not found."
    exit 1
fi

# Get new version from command line argument or prompt user
if [ -n "$1" ]; then
    new_version="$1"
else
    read -p "Enter new Go version (e.g., 1.25.0): " new_version
fi

# Validate version format (basic check)
if [[ ! $new_version =~ ^[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
    echo "Error: Invalid version format. Please use format like 1.25.0"
    exit 1
fi

echo "Upgrading Go version to: $new_version"
echo ""

# Find all go.mod files in projects
go_mod_files=()
while IFS= read -r -d '' file; do
    go_mod_files+=("$file")
done < <(find projects -name "go.mod" -print0)

# Check if any go.mod files exist
if [ ${#go_mod_files[@]} -eq 0 ]; then
    echo "No go.mod files found in projects directory."
    exit 1
fi

echo "Found ${#go_mod_files[@]} go.mod file(s):"
for file in "${go_mod_files[@]}"; do
    echo "  - $file"
done
echo ""

# Process each go.mod file
updated_count=0
for go_mod_file in "${go_mod_files[@]}"; do
    echo "Processing: $go_mod_file"
    
    # Get current version
    current_version=$(grep "^go " "$go_mod_file" | awk '{print $2}')
    
    if [ -z "$current_version" ]; then
        echo "  ⚠️  Warning: No 'go' directive found in $go_mod_file"
        continue
    fi
    
    echo "  Current version: $current_version"
    
    if [ "$current_version" = "$new_version" ]; then
        echo "  ✅ Already at version $new_version"
    else
        # Create backup
        cp "$go_mod_file" "$go_mod_file.backup"
        
        # Update the version
        sed -i.tmp "s/^go [0-9]\+\.[0-9]\+\.[0-9]\+/go $new_version/" "$go_mod_file"
        
        # Remove temporary file
        rm -f "$go_mod_file.tmp"
        
        # Verify the change
        updated_version=$(grep "^go " "$go_mod_file" | awk '{print $2}')
        if [ "$updated_version" = "$new_version" ]; then
            echo "  ✅ Updated to version $new_version"
            ((updated_count++))
        else
            echo "  ❌ Failed to update version"
            # Restore backup
            mv "$go_mod_file.backup" "$go_mod_file"
        fi
    fi
    echo ""
done

echo "=== Summary ==="
echo "Total go.mod files processed: ${#go_mod_files[@]}"
echo "Files updated: $updated_count"
echo ""

if [ $updated_count -gt 0 ]; then
    echo "✅ Go version upgrade completed successfully!"
    echo ""
    echo "Next steps:"
    echo "1. Review the changes in the go.mod files"
    echo "2. Run 'go mod tidy' in each project directory to update dependencies"
    echo "3. Test your projects to ensure compatibility"
    echo ""
    echo "To revert changes, restore from .backup files:"
    for go_mod_file in "${go_mod_files[@]}"; do
        if [ -f "$go_mod_file.backup" ]; then
            echo "  mv $go_mod_file.backup $go_mod_file"
        fi
    done
else
    echo "ℹ️  No files needed updating (all already at version $new_version)"
fi
