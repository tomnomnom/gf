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

I wrote `gf` to gives names to the pattern and flag combinations I use all the time. So the above command
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

There are some example pattern files in the `examples` directory.

### Auto Complete

There's an auto-complete script included, so you can hit 'tab' to show you what your options are:

```
▶ gf <tab>
base64       debug-pages  fw           php-curl     php-errors   php-sinks    php-sources  sec          takeovers    urls
```

To get auto-complete working you need to `source` the `gf-completion.bash` file in your `.bashrc` or similar:

```
▶ echo 'source ~/path/to/gf-completion.bash' >> ~/.bashrc
```

## Install

If you've got Go installed and configured you can install `gf` with:

```
▶ go get -u github.com/tomnomnom/gf
```


