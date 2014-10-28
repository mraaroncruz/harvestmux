## Harvestmux - harvestapp plugin for tmux

I just hacked this out. It is not `go get`able yet. It's pretty raw

If anyone is interested in binaries or better instructions, contact me via twitter [@mraaroncruz](http://twitter.com/mraaroncruz)

### Screenshot

![wow](https://www.evernote.com/shard/s25/sh/1e4d48b8-1300-4bfa-af56-d55aa3e99d50/e97c6dcbdc869aeb6d34b72f61bb5bee/deep/0/1.-tmux-(tmux)-and-Contributors-to-go-yaml-yaml.png)

### Usage
In your `~/.tmux.conf`  
`set-option -g status-right "#[fg=magenta]hrs #(harvestmux -config ~/.harvest/config.yml) #[fg=$TMUX_SHELL_COLOR]$sHost#[default]#[fg=cyan] %d %b %R:%S"`

the `-config` part is the absolute path to your `config.yml` file. Just make a copy of the `config.example.yml` in the repo

### Thank you

[niemeyer](https://github.com/niemeyer) for [go-yaml](https://github.com/go-yaml/yaml)

### License

[go-yaml license](https://github.com/go-yaml/yaml/blob/v2/LICENSE)

[So you don't have to scroll up](https://github.com/pferdefleisch/harvestmux/blob/master/LICENSE)


__Enjoy__
