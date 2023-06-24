# tx-PicUp
Uploading screenshots to cos using pipeline

code just for fun:)

## Usage

Works with [grim](https://sr.ht/~emersion/grim/):

```zsh
grim -g "$(slurp)" - | ./tx-PicUp
```

You also use [wl-clipboard](https://github.com/bugaevc/wl-clipboard) copy the return image path:

```zsh
grim -g "$(slurp)" - | ./tx-PicUp | wl-copy
```

pipeline is a great invention!
