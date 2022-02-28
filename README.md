# brbscenes

A tool for twitch streamers that lets twitch chat select which brbscreen to show.

## Setup

1. [Go here](https://github.com/tlanfer/brbscenes/releases/tag/latest) and download the latest `brbscenes.exe`
2. Run the executable once. It will create an example config file named `config.yaml` that will look roughly like this:
3. 

```yaml
channel: alasdair                # What twitch channel to watch
cooldown: 30s                    # Whats the default cooldown, if a screen doesnt have a custom cooldown
obs:
    brb_scene: brb               # On which scene this tool be active?
                                 # If any other scene is active, we ignore all chat messages 
    password: roflmao            # The password for your obs websocket server
    port: 4444                   # ... and its port
sources:
    - name: tetris               # The name of a source on your brb scene
      keyword: '!tetris'         # The keyword in chat to enable it ( '!'-prefix is optional)
      cooldown: 15s              # Individual cooldown for this scene (optional)
    - name: trover
      keyword: '!trover'
    - name: worms
      keyword: '!worms'
    - name: twiggie
      keyword: '!twiggie'
```

In your OBS, set up a BRB scene with the source

![img.png](docs/obs.png)