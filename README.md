# ccheck

CCheck is a copyright linter.

It checks each project file for the presence of the copyright notice defined in the .ccheck configuration file.

Use the `.ccheckignore` file to ignore certain files or directories. It uses the same syntax as `.gitignore`.

## Usage

```shell
just build
``

This will build the ccheck binary in the `./bin` directory.

```shell
./bin/ccheck <your_project_directory>
```
