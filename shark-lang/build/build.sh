#!/bin/bash

# This script builds the Shark compiler with static linking

# Set the desired compiler options
COMPILER_OPTIONS="-static"

# Navigate to the directory where the Shark compiler is located
cd /path/to/shark/compiler

# Build the compiler with static linking
make $COMPILER_OPTIONS

echo "Shark compiler built successfully with static linking!"