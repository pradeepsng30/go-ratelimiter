#!/bin/bash

# Configuration
GIT_REPO="https://github.com/pradeepsng30/go-ratelimiter"
MODULE_NAME="github.com/pradeepsng30/go-ratelimiter"

# Function to get the latest version
get_latest_version() {
    # Fetch tags from the repository
    latest_tag=$(git tag -l "v*" | sort -V | tail -n 1)
    if [ -z "$latest_tag" ]; then
        echo "No previous versions found. Starting with v1.0.0."
        echo "v1.0.0"
    else
        echo "Latest version found: $latest_tag"
        echo "$latest_tag"
    fi
}

# Function to increment the minor version
increment_minor_version() {
    version="$1"
    # Remove 'v' prefix
    version="${version#v}"
    
    # Split version into major, minor, and patch
    IFS='.' read -r major minor patch <<< "$version"
    
    # Increment the minor version
    minor=$((minor + 1))
    
    # Reset patch version to 0
    patch=0
    
    # Create new version string
    new_version="v${major}.${minor}.${patch}"
    
    echo "$new_version"
}

# Main script execution
echo "Fetching the latest version..."
latest_version=$(get_latest_version)

# Increment the version
next_version=$(increment_minor_version "$latest_version")
echo "New version: $next_version"

# Check if the user is in the module directory
if [ ! -f "go.mod" ]; then
    echo "Error: No go.mod file found. Make sure you are in the module directory."
    exit 1
fi

# Set Go module path
export GOPROXY="https://proxy.golang.org"

# Check if there are any uncommitted changes
if ! git diff-index --quiet HEAD --; then
    echo "Error: You have uncommitted changes. Please commit or stash them before publishing."
    exit 1
fi

# Create a new tag for the version
git tag "$next_version"

# Push the tag to the remote repository
git push origin "$next_version"

# Run tests to ensure everything is working before publishing
echo "Running tests..."
go test ./...

# Build the module (optional, but ensures that the module builds correctly)
echo "Building module..."
go build ./...

# Publish the module
echo "Publishing module..."
go list -m $MODULE_NAME@$next_version

# Notify the user
echo "Module $MODULE_NAME has been published with version $next_version."

# Optional: Clean up tags (if you want to delete the tag after publishing)
# git tag -d "$next_version"
# git push origin --delete "$next_version"