#! /bin/bash

artifact=webpage-analyzer
golangcifile=golangci-report.xml
coverfile=cover.out

# Format The Code To GoLang standards
echo "Formatting Code To GoLang standards..."
go fmt ./...
echo "Formatting Complete."

# Clean the code
echo "Cleaning the code..."
go clean
echo "Code cleaning completed"

# Remove the existing cover.out file if exists
echo "Removing existing $coverfile file if exists..."
if test -f "$coverfile"; then
  rm -f $coverfile
  echo "$coverfile removed"
  else echo "$coverfile not exists"
fi

# Execute unit test cases
echo "Running unit test cases..."
go test ./... -v -cover -coverpkg=./... -coverprofile=$coverfile
echo "Unit test cases execution completed"

# Remove the existing GoLangCI Lint file if exists
echo "Removing existing $golangcifile file if exists..."
if test -f "$golangcifile"; then
  rm -f $golangcifile
  echo "$golangcifile removed"
  else echo "$golangcifile not exists"
fi

# Run GoLangCI Linter
echo "Running GoLangCI Linter..."
golangci-lint run ./...
echo "$golangcifile file generated"

# Run SonarQube
echo "Running SonarQube..."
./sonar-scan.sh
echo "SonarQube executed successfully."

# Build Go Binary
echo "Building Binary..."
go build -o $artifact
if [[ $? != 0 ]]; then
  echo "Build Failed."
  exit 1
fi
echo "Build Passed."
