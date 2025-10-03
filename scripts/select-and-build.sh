#!/bin/bash

# Script to select and build a project from the projects directory
# Usage: ./select-and-build.sh

echo "=== Project Selector and Builder ==="
echo ""

# Check if projects directory exists
if [ ! -d "projects" ]; then
    echo "Error: projects directory not found."
    exit 1
fi

# Find all project directories
projects=()
while IFS= read -r -d '' project_dir; do
    # Extract just the project name (slug) from the path
    project_name=$(basename "$project_dir")
    projects+=("$project_name")
done < <(find projects -maxdepth 1 -type d -not -path "projects" -print0)

# Check if any projects exist
if [ ${#projects[@]} -eq 0 ]; then
    echo "No projects found in the projects directory."
    echo "Create a new project first with: make new-project"
    exit 1
fi

echo "Available projects:"
echo ""

# Display projects with numbers
for i in "${!projects[@]}"; do
    echo "$((i+1)). ${projects[$i]}"
done

echo ""
read -p "Select a project (1-${#projects[@]}): " selection

# Validate selection
if ! [[ "$selection" =~ ^[0-9]+$ ]] || [ "$selection" -lt 1 ] || [ "$selection" -gt ${#projects[@]} ]; then
    echo "Error: Invalid selection. Please choose a number between 1 and ${#projects[@]}."
    exit 1
fi

# Get selected project slug
selected_slug="${projects[$((selection-1))]}"

echo ""
echo "Selected project: $selected_slug"
echo ""

# Check if the project has the required structure
if [ ! -f "projects/$selected_slug/cmd/main.go" ]; then
    echo "Error: Project '$selected_slug' does not have the required cmd/main.go file."
    exit 1
fi

echo "Building project: $selected_slug"
echo ""

# Create bin directory if it doesn't exist
mkdir -p "projects/$selected_slug/bin"

# Build the project
cd "projects/$selected_slug"
if go build -o bin/"$selected_slug" cmd/main.go; then
    echo ""
    echo "✅ Project '$selected_slug' built successfully!"
    echo "Binary location: projects/$selected_slug/bin/$selected_slug"
    echo ""
    echo "To run the project:"
    echo "  cd projects/$selected_slug"
    echo "  ./bin/$selected_slug"
    echo ""
    echo "Or run directly:"
    echo "  projects/$selected_slug/bin/$selected_slug"
else
    echo ""
    echo "❌ Build failed for project '$selected_slug'"
    exit 1
fi
