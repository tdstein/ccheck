#!/usr/bin/env bash

# MIT License
#
# Copyright (c) 2023 Taylor Steinberg
#
# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to deal
# in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:
#
# The above copyright notice and this permission notice shall be included in all
# copies or substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
# OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
# SOFTWARE.

package=$1
if [[ -z "$package" ]]; then
	echo "usage: $0 <package-name>"
	exit 1
fi

package_name=$(basename $package)

docker pull ghcr.io/choffmeister/git-describe-semver:latest
version=$(docker run --rm -v $PWD:/workdir ghcr.io/choffmeister/git-describe-semver:latest --fallback "v0.0.0")

platforms=($(go tool dist list --json | jq '.[] | select( .FirstClass) | .GOOS + "/" + .GOARCH' | tr -d '"'))

for platform in "${platforms[@]}"; do
	echo "Building $platform"
	platform_split=(${platform//\// })
	GOOS=${platform_split[0]}
	GOARCH=${platform_split[1]}
	output_name='./bin/'$package_name'_'$version'_'$GOOS'_'$GOARCH
	if [ $GOOS = "windows" ]; then
		output_name+='.exe'
	fi

	env GOOS=$GOOS GOARCH=$GOARCH go build -o $output_name -ldflags "-X main.Version=$version" $package
	if [ $? -ne 0 ]; then
		echo 'An error has occurred! Aborting the script execution...'
		exit 1
	fi
done

# Create alias for current OS
current_os=$(go env GOOS)
current_arch=$(go env GOARCH)
current_binary="./bin/${package_name}_${version}_${current_os}_${current_arch}"
if [ "$current_os" = "windows" ]; then
	current_binary+='.exe'
fi

alias_path="./bin/$package_name"
if [ -f "$current_binary" ]; then
	echo "Aliasing $alias_path -> $current_binary"
	rm -f "$alias_path"
	ln -s "$(basename "$current_binary")" "$alias_path"
else
	echo "Warning: Binary for current platform ($current_os/$current_arch) not found"
fi
