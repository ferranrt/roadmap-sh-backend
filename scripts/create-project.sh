#!/bin/bash

# Script to create a new project structure with the given slug
# Usage: ./create-project.sh

echo "=== Project Structure Creator ==="
echo "This script will create a new project structure in the projects directory."
echo ""

# Prompt user for slug
read -p "Enter project slug (lowercase, hyphens allowed): " slug

# Validate slug format (lowercase, hyphens, alphanumeric)
if [[ ! $slug =~ ^[a-z0-9-]+$ ]]; then
    echo "Error: Slug must contain only lowercase letters, numbers, and hyphens."
    exit 1
fi

# Check if project already exists
if [ -d "projects/$slug" ]; then
    echo "Error: Project '$slug' already exists in projects directory."
    exit 1
fi

echo ""
echo "Creating project structure for: $slug"
echo ""

# Create directory structure
echo "Creating directories..."
mkdir -p "projects/$slug/internal"
mkdir -p "projects/$slug/cmd"

# Create main.go file
echo "Creating main.go file..."
cat > "projects/$slug/cmd/main.go" << EOF
package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Project slug: $slug")
	log.Println("Application started for project: $slug")
}
EOF

# Create go.mod file
echo "Creating go.mod file..."
cat > "projects/$slug/go.mod" << EOF
module ferranrt/roadmap-sh/backend/$slug

go 1.24.0
EOF

echo ""
echo "✅ Project structure created successfully!"
echo ""
echo "Created structure:"
echo "  projects/$slug/"
echo "  ├── internal/"
echo "  └── cmd/"
echo "      └── main.go"
echo ""
echo "You can now navigate to the project directory:"
echo "  cd projects/$slug"
echo ""
echo "To run the project:"
echo "  go run cmd/main.go"
