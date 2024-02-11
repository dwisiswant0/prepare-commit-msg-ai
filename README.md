# Prepare Git Commit Message with AI

`prepare-commit-msg-ai` helps you to write commit message based on code changes with AI.

* [Prerequisites](#prerequisites)
* [Installation](#installation)
  * [Usage](#usage)
* [Getting Started](#getting-started)
  * [Install](#install)
  * [Uninstall](#uninstall)
  * [Edit](#edit)
* [License](#license)

---

## Prerequisites

* `go`
* `git`
* [`mods`](https://github.com/charmbracelet/mods)

## Installation

```
go install github.com/dwisiswant0/prepare-commit-msg-ai@latest
```

### Usage

```console
$ prepare-commit-msg-ai -h

  Prepare Git Commit Message with AI v0.0.1
  --
  made with ♥ by @dwisiswant0

Usage:
  prepare-commit-msg-ai [COMMAND]

Commands:
  install           Install hook in current Git project (alias: isntall, i)
  uninstall         Uninstall generated hook file (alias: unisntall, u)
  edit              Edit hook file (alias: ed, e)

Examples:
  prepare-commit-msg-ai install
```

## Getting Started

Explore the basic functionalities of the tool with the following actions:

### Install

To set up this hook, you must individually install it in each Git project, and the installation should take place within the respective project directory. <sup>Ain't no way to install this globally unless you make an alias for `git commit`.</sup>

```bash
prepare-commit-msg-ai install
```

> [!TIP]
> Alias for install: **isntall**, **i**.

> [!IMPORTANT]
> Don't forget to set up `mods` with [`--settings`](https://github.com/charmbracelet/mods#settings).

#### How it works?

```console
$ cd your-git-project/
$ # make changes
$ git add .
$ git commit
$ # generating commit message for ya,
$ # and it'll pop your `EDITOR` to confirm the commit message.
$ # satifying? save it and leave it.
```

### Uninstall

To remove the `prepare-commit-msg` hook in your current Git project:

```bash
prepare-commit-msg-ai uninstall
```

> [!TIP]
> Alias for uninstall: **unisntall**, **u**.

### Edit

Want to change model or prompt?

```bash
prepare-commit-msg-ai edit
```

> [!TIP]
> Alias for edit: **ed**, **e**.

> [!NOTE]
> Currently, the hook works correctly when tested with OpenAI's models, using default settings from `mods`. To clarify, I have NOT tested its compatibility with `GPT4All-J *` models or models from Azure OpenAI at this time.<br>
> <sup>I mean, I gave the `llama2`, `codellama`, `openchat`, etc. models a shot with different prompts on Ollama, but goddamn the output keeps throwing off the system prompt and always goes off-script, making it a pain for me to create an extractor.</sup>

## License

Contributions are welcome! Comments? Feedbacks? Thoughts? Just [open an issue](https://github.com/dwisiswant0/prepare-commit-msg-ai/issues/new).

`prepare-commit-msg-ai` is made with ♥ by [**@dwisiswant0**](https://github.com/dwisiswant0) under the MIT license. See [LICENSE](/LICENSE).