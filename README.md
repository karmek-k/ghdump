# ghdump

Bulk clone GitHub repos for backup - Go version

The Rust version gave me headaches while writing, so I switched to Go.

I finished the first version in less than 20 commits (1-2 days).

## Features

- Single binary - download it, use it, throw it away!
- Optional authentication via a [personal access token (PAT)](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token)
- Adjustable fork cloning (by default, forks are not cloned)

## Installation

### Binaries

Go to the [releases page](https://github.com/karmek-k/ghdump/releases) and grab a binary for your OS.

### Building from source

[Download the Go toolchain](https://go.dev/dl/). Clone this repository with `git` and navigate to the resulting directory.

Download dependencies:

```
go mod download
```

And then build the binary:

```
go build
```

## Usage

Run `ghdump` to see all commands:

```
ghdump
```

To clone repos of a user or an organization:

```
ghdump clone <username>
```

### Available flags for `clone`

- `--token`/`-t` - allows you to provide a [token (PAT)](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token), necessary for cloning private repositories
- `--output-dir`/`-o` - the directory name where repos are cloned - `dump` by default
- `--visibility`/`-v` - repo visibility (`public`, `private`, `all`) - `all` by default
- `--clone-forks`/`-f` - pass it to also clone repos forked by the user / organization

## Examples

Clone `user321`'s public repos

```
ghdump clone user321
```

Clone all `user321`'s repos with a PAT

```
ghdump clone user321 -t PUT_TOKEN_HERE
```

Clone `user321`'s private repos to `/home/user/repos`

```
ghdump clone user321 -t PUT_TOKEN_HERE -v private -o /home/user/repos
```

## Contributing

Always welcome!

If you want to contribute, please search through other issues first - if there are none that are applicable,
create your own.

### Command scaffolding

I used [`cobra-cli`](https://github.com/spf13/cobra-cli) to generate new commands.

In order to create one, download it:

```
go install github.com/spf13/cobra-cli@latest
```

And then run (for a new command called `mynewcommand` - change to your needs):

```
cobra-cli add mynewcommand
```

**Delete comments at the top of the newly generated file.**

## License

Licensed under the [GNU General Public License v3.0](LICENSE).
