# Doc Bot
A tool to streamline Neovim plugin documentation.

## Installation

### Grab the binary
Head over to the releases page [here](https://github.com/phanorcoll/doc_bot/releases) to download the binary that works on your computer. Easy peasy!


### Using Go Install
Got Go installed, use the following command.
```bash
$ go install github.com/yourusername/doc_bot@latest
```

### Usage
```bash
$ doc_bot [flags]
  -d, --dir string   Directory to scan for plugins
  -o, --output string  Output file path (default "README.md")
  -t, --template string Template file path (default "template.tmpl")
```

### Example
```bash
 $ doc_bot -d plugins -o README.md

# with -t flag
$ doc_bot -d plugins -o README.md -t new_doc.tmpl
```
This will scan the plugins directory, parse plugin headers, and generate a README file.

### Configuration
You can customize the output by creating a **template.tmpl** file. This file uses Go's template syntax to define the structure of the generated README.

- Example **template.tmpl**
```tmpl
# Neovim Plugin Documentation

{{ range . }}
## {{ .Title }}

<img src="assets/goods.png" width="68" />

**Package name:** {{ .Title }}
**Description:** {{ .Desc }}
**URL:** {{ .URL }}

---

{{ end }}
```
In this example, for each plugin, it will render this block, adding an image fro the *assets/* folder.

### Setup metadata block for Neovim plugins

- Plugin configuration example
```lua
-- url: https://github.com/rmagatti/auto-session
-- desc: Auto save and restore the last session
return {
  'rmagatti/auto-session',
  lazy = false,
  ---enables autocomplete for opts
  ---@module "auto-session"
  opts = {
    suppressed_dirs = { '~/', '~/Projects', '~/Downloads', '/' },
    log_level = 'error',
  }
}
```
This plugin utilizes a **metadata block** at the top of the file to define configuration options. Let's break down the different sections:

- url: (https://github.com/rmagatti/auto-session)

    This section specifies the URL of the plugin's repository, providing users with a way to access the source code and additional information.
- desc: (Auto save and restore the last session)

    This section briefly describes the functionality of the plugin.

From this block, **Doc bot** will generate the documentation.

### TODO
- [ ] Scan single file for multiple metadata blocks.
- [ ] When scanning multiple files, get all metadata blocks for each.
