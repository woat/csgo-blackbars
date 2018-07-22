# CSGO Black Bars

A simple resolution switcher to change native resolution while playing CSGO.

Requirements:
- Steam
- NVIDIA Server X Settings

If you are using anything else, you can probably use `xrandr` for better compatibility. This tool is made specifically for those who are running into issues with proprietary drivers offered by NVIDIA.

You can check compatibility by running the `./csgo-blackbars` binary yourself.

### Installation

Installation is very simple, just place the `bin/csgo-blackbars` binary anywhere in your filesystem. Then make a `~/.csbb` file and write two lines, first line is `csgo resolution` and second line is `origin resolution`. Then set the launch options to `path/of/binary %command%`. 

Read further for more detail.

### Setting up .csbb

Example (from [example file](https://github.com/woat/csgo-blackbars/blob/master/.csbb)):
```
dvi-i-0: nvidia-auto-select +0+0, dvi-i-3: 1024x768 +1920+0 {viewportin=1024x768, viewportout=814x768+100+0}
dvi-i-0: nvidia-auto-select +0+0, dvi-i-3: 1920x1080 +1920+0
```

Black bars can be achieved by changing the last two `+X+Y` (adjust `+X`) numbers after `ViewPortOut`.

You can get these settings by opening up the NVIDIA X Server Settings and clicking `Save to X Configuration File` and saving it. 

![](https://i.imgur.com/IBGgO9b.png)
![](https://i.imgur.com/n3pLqFN.png)

As you can see, these are my `origin resolution` settings. Just repeat this step again with your desired CSGO resolution.

### Setting up launch options

All that's left to do is set your launch options as follows:
```
path/to/binary %command%
```

My launch options look like this:  
![](https://i.imgur.com/hZjyBod.png)

Now you can click OK and start your game.
