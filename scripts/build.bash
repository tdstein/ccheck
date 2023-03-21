#!/usr/bin/env bash

package=$1
if [[ -z "$package" ]]; then
  echo "usage: $0 <package-name>"
  exit 1
fi

package_name=`basename $package`

docker pull ghcr.io/choffmeister/git-describe-semver:latest
version=$(docker run --rm -v $PWD:/workdir ghcr.io/choffmeister/git-describe-semver:latest --fallback "v0.0.0")
	
platforms=($(go tool dist list --json | jq '.[] | select( .FirstClass) | .GOOS + "/" + .GOARCH' | tr -d '"'))

for platform in "${platforms[@]}"
do
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