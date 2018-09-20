# CSGO Black Bars

![](https://goreportcard.com/badge/github.com/woat/csgo-blackbars)

A simple resolution switcher to change native resolution while playing CSGO.

Requirements:
- Steam
- NVIDIA Server X Settings

If you are using anything else, you can probably use `xrandr` for better compatibility. This tool is made specifically for those who are running into issues with proprietary drivers offered by NVIDIA.

You can check compatibility by running the `./csgo-blackbars` binary yourself.

### Installation

Installation is very simple, just place the `csgo-blackbars` binary anywhere in your filesystem. Then make a `~/.csbb` file and write two lines, first line is `csgo resolution` and second line is `origin resolution`. Then set the launch options to `path/of/binary %command%`. 

Read further for more detail.

### Setting up .csbb

Example (from [example .csbb](https://github.com/woat/csgo-blackbars/blob/master/.csbb)):
```
dvi-i-0: nvidia-auto-select +0+0, dvi-i-3: 1024x768 +1920+0 {viewportin=1024x768, viewportout=814x768+100+0}
dvi-i-0: nvidia-auto-select +0+0, dvi-i-3: 1920x1080 +1920+0
```

Black bars can be achieved by changing the last two `+X+Y` (adjust `+X`) numbers after `ViewPortOut`.

You can get these settings by opening up the NVIDIA X Server Settings and clicking `Save to X Configuration File` and saving it. 

![](https://i.imgur.com/IBGgO9b.png)

You can then navigate to the `.conf` file that you saved and open it up in your favorite text editor. You would simply copy and paste the line required (as shown below) into your `.csbb`.

![](https://i.imgur.com/n3pLqFN.png)

### Setting up launch options

All that's left to do is set your launch options as follows:
```
path/to/binary %command%
```

My launch options look like this:  
![](https://i.imgur.com/hZjyBod.png)

Now you can click OK and start your game.

# Troubleshooting

You can simply run the binary directly by typing `./csgo-blackbars` while in the same directory. This will enable troubleshooting mode and log out any errors that are occuring.

If the troubleshooter fails to find any error and the program still doesn't work, feel free to open up an issue or make a pull request. Please make sure to be as descriptive as possible and post system specs.
