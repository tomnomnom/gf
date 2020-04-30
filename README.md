# gf

A wrapper around grep to avoid typing common patterns.

## What? Why?

I use grep a *lot*. When auditing code bases, looking at the output of [meg](https://github.com/tomnomnom/meg),
or just generally dealing with large amounts of data. I often end up using fairly complex patterns like this one:

```
▶ grep -HnrE '(\$_(POST|GET|COOKIE|REQUEST|SERVER|FILES)|php://(input|stdin))' *
```

It's really easy to mess up when typing all of that, and it can be hard to know if you haven't got any
results because there are non to find, or because you screwed up writing the pattern or chose the wrong flags.

I wrote `gf` to give names to the pattern and flag combinations I use all the time. So the above command
becomes simply:

```
▶ gf php-sources
```

### Pattern Files

The pattern definitions are stored in `~/.gf` as little JSON files that can be kept under version control:

```
▶ cat ~/.gf/php-sources.json
{
    "flags": "-HnrE",
    "pattern": "(\\$_(POST|GET|COOKIE|REQUEST|SERVER|FILES)|php://(input|stdin))"
}
```

To help reduce pattern length and complexity a little, you can specify a list of multiple patterns too:

```
▶ cat ~/.gf/php-sources-multiple.json
{
    "flags": "-HnrE",
    "patterns": [
        "\\$_(POST|GET|COOKIE|REQUEST|SERVER|FILES)",
        "php://(input|stdin)"
    ]
}
```

There are some more example pattern files in the `examples` directory.

You can use the `-save` flag to create pattern files from the command line:

```
▶ gf -save php-serialized -HnrE '(a:[0-9]+:{|O:[0-9]+:"|s:[0-9]+:")'
```

### Auto Complete

There's an auto-complete script included, so you can hit 'tab' to show you what your options are:

```
▶ gf <tab>
base64       debug-pages  fw           php-curl     php-errors   php-sinks    php-sources  sec          takeovers    urls
```

#### Bash

To get auto-complete working you need to `source` the `gf-completion.bash` file in your `.bashrc` or similar:

```
source ~/path/to/gf-completion.bash
```

#### Zsh

To get auto-complete working you need to enable autocomplete (not needed if you have oh-my-zsh) using `autoload -U compaudit && compinit` or by putting it into `.zshrc`

Then `source` the `gf-completion.zsh` file in your `.zshrc` or similar:

```
source ~/path/to/gf-completion.zsh
```

Note: if you're using oh-my-zsh or similar you may find that `gf` is an alias for `git fetch`. You can either
alias the gf binary to something else, or `unalias gf` to remove the `git fetch` alias.

### Using custom engines

There are some amazing code searching engines out there that can be a better replacement for grep.
A good example is [the silver searcher](https://github.com/ggreer/the_silver_searcher).
It's faster (like **way faster**) and presents the results in a more visually digestible manner.
In order to utilize a different engine, add `engine: <other tool>` to the relevant pattern file:
```bash
# Using the silver searcher instead of grep for the aws-keys pattern:
# 1. Adding "ag" engine
# 2. Removing the E flag which is irrelevant for ag
{
  "engine": "ag",
  "flags": "-Hanr",
  "pattern": "([^A-Z0-9]|^)(AKIA|A3T|AGPA|AIDA|AROA|AIPA|ANPA|ANVA|ASIA)[A-Z0-9]{12,}"
}
```
* Note: Different engines use different flags, so in the example above, the flag `E` has to be removed from the `aws-keys.json` file in order for ag to successfully run.


## Install

If you've got Go installed and configured you can install `gf` with:

```
▶ go get -u github.com/tomnomnom/gf
```

If you've installed using `go get`, you can enable auto-completion to your `.bashrc` like this:

```
▶ echo 'source $GOPATH/src/github.com/tomnomnom/gf/gf-completion.bash' >> ~/.bashrc
```

Note that you'll have to restart your terminal, or run `source ~/.bashrc` for the changes to
take effect.

To get started quickly, you can copy the example pattern files to `~/.gf` like this:

```
▶ cp -r $GOPATH/src/github.com/tomnomnom/gf/examples ~/.gf
```

My personal patterns that I've included as examples might not be very useful to you, but hopefully
they're still a reasonable point of reference.

## Contributing

I'd actually be most interested in new pattern files! If you've got something you regularly grep for
then feel free to issue a PR to add new pattern files to the examples directory.

Bug fixes are also welcome as always :)
