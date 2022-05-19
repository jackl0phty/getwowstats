[![ko-fi](https://ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/M4M3CQPLY)

# GETWOWSTATS
Rust CLI utility that hits Blizzard's pubic API for WoW character stats.

# Supported Platforms
I will unoffically only be supporting Linux Mint since that's the platform I run this tool from.
However it _should_ work on most *NIX systems that can run Rust.

# Install Prerequisites

## Install Rustup ( As non-priv user)
```
$ cd /tmp
$ curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs -o rust.sh
$ vim rust.sh
$ chmod u+x rust.sh
info: downloading installer

Welcome to Rust!

This will download and install the official compiler for the Rust
programming language, and its package manager, Cargo.

Rustup metadata and toolchains will be installed into the Rustup
home directory, located at:

  /home/yoda/.rustup

This can be modified with the RUSTUP_HOME environment variable.

The Cargo home directory located at:

  /home/yoda/.cargo

This can be modified with the CARGO_HOME environment variable.

The cargo, rustc, rustup and other commands will be added to
Cargo's bin directory, located at:

  /home/yoda/.cargo/bin

This path will then be added to your PATH environment variable by
modifying the profile files located at:

  /home/yoda/.profile
  /home/yoda/.bash_profile
  /home/yoda/.bashrc

You can uninstall at any time with rustup self uninstall and
these changes will be reverted.

Current installation options:


   default host triple: x86_64-unknown-linux-gnu
     default toolchain: stable (default)
               profile: default
  modify PATH variable: yes

1) Proceed with installation (default)
2) Customize installation
3) Cancel installation
>

info: profile set to 'default'
info: default host triple is x86_64-unknown-linux-gnu
info: syncing channel updates for 'stable-x86_64-unknown-linux-gnu'
info: latest update on 2022-02-24, rust version 1.59.0 (9d1b2106e 2022-02-23)
info: downloading component 'cargo'
info: downloading component 'clippy'
info: downloading component 'rust-docs'
 19.3 MiB /  19.3 MiB (100 %)  13.4 MiB/s in  1s ETA:  0s
info: downloading component 'rust-std'
 25.7 MiB /  25.7 MiB (100 %)  13.9 MiB/s in  2s ETA:  0s
info: downloading component 'rustc'
 53.8 MiB /  53.8 MiB (100 %)  13.4 MiB/s in  4s ETA:  0s
info: downloading component 'rustfmt'
info: installing component 'cargo'
info: installing component 'clippy'
info: installing component 'rust-docs'
 19.3 MiB /  19.3 MiB (100 %)  13.3 MiB/s in  1s ETA:  0s
info: installing component 'rust-std'
 25.7 MiB /  25.7 MiB (100 %)  15.1 MiB/s in  1s ETA:  0s
info: installing component 'rustc'
 53.8 MiB /  53.8 MiB (100 %)  18.0 MiB/s in  3s ETA:  0s
info: installing component 'rustfmt'
info: default toolchain set to 'stable-x86_64-unknown-linux-gnu'

  stable-x86_64-unknown-linux-gnu installed - rustc 1.59.0 (9d1b2106e 2022-02-23)


Rust is installed now. Great!

To get started you may need to restart your current shell.
This would reload your PATH environment variable to include
Cargo's bin directory ($HOME/.cargo/bin).

To configure your current shell, run:
source $HOME/.cargo/env
$ source $HOME/.cargo/env
```

## Confirm Rust version & that we're up-to-date
```
$ rustup self update
  info: checking for self-updates
  rustup unchanged - 1.24.3
```
## Follow Rust nightly build
```
$ rustup default nightly
info: syncing channel updates for 'nightly-x86_64-unknown-linux-gnu'
info: latest update on 2022-03-24, rust version 1.61.0-nightly (9f4dc0b4d 2022-03-23)
info: downloading component 'cargo'
info: downloading component 'clippy'
info: downloading component 'rust-docs'
 19.6 MiB /  19.6 MiB (100 %)  14.3 MiB/s in  1s ETA:  0s
info: downloading component 'rust-std'
 26.6 MiB /  26.6 MiB (100 %)  14.0 MiB/s in  2s ETA:  0s
info: downloading component 'rustc'
 55.3 MiB /  55.3 MiB (100 %)  13.3 MiB/s in  4s ETA:  0s
info: downloading component 'rustfmt'
info: installing component 'cargo'
info: installing component 'clippy'
info: installing component 'rust-docs'
 19.6 MiB /  19.6 MiB (100 %)  13.0 MiB/s in  1s ETA:  0s
info: installing component 'rust-std'
 26.6 MiB /  26.6 MiB (100 %)  15.0 MiB/s in  1s ETA:  0s
info: installing component 'rustc'
 55.3 MiB /  55.3 MiB (100 %)  17.9 MiB/s in  3s ETA:  0s
info: installing component 'rustfmt'
info: default toolchain set to 'nightly-x86_64-unknown-linux-gnu'

  nightly-x86_64-unknown-linux-gnu installed - rustc 1.61.0-nightly (9f4dc0b4d 2022-03-23)
```
## Initial Project Setup
```
$ cargo init --lib
mkdir src/lib
mv src/lib.rs src/lib
mkdir src/bin
touch src/bin/getwowstats.rs
```
## 

# HARD REQUIREMENTS YOU MUST DO TO USE GETWOWSTATS
You MUST have a valid WoW token for Blizzard's API or you can't use this tool. No worries it's free.

# How/Where do I get a WoW API Token?
Great question and Blizz has your back. Read their doc [here](https://develop.battle.net/documentation/guides/getting-started).

# OK I Have a WoW API Token What next?
You MUST export ENV VAR WOW_API_TOKEN with the value of a valid token.
```
export WOW_API_TOKEN=your_secret_api_wow_token_goes_here
```
# Sample Usage & Output

## Get profile summary stats for a given character
```
./getwowstats -r aggramar -l us -c pownage -m GetProfileSummary
Name: Pownage
Gender: Male
Faction: Horde
Race: Orc
Class: Warlock
Realm: Aggramar
Level: 54
Active Spec: Affliction
Experience: 246258
Achieveent Points: 1310
Last Login: 1.622774087e+12
Avg. Item Lvl.: 94
Equipped Item Lvl.: 88
```
## Get version of getwowstats
```
./getwowstats -veresion
getwowstats: Version 1.0
```
# TESTING

## Install the following in order to test getwowstats
Run the formatting linter gofmt
```
gofmt -e getwowstats.go
```

goimport adds missing imports and removes unreferenced imports.
```
 go get golang.org/x/tools/cmd/goimports
```

Just say NO to global variables
```
go get 4d63.com/gochecknoglobals```
