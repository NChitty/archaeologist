This is my application for playing the game [ArtifactsMMO](https://artifactsmmo.com/). The `pkg` directory is meant to wrap around [promiseofcake's artifactsmmo-go-client](https://github.com/promiseofcake/artifactsmmo-go-client).
The application is split into two binaries, a CLI client and a combination of DiscordBot/web app.

# CLI Client
The CLI client was made to quickly test different facets of the application in REPL format. There is
a concept of actors which will do a task to completion (or error).

# DiscordBot
The DiscordBot allows to create actors from bot commands. The web app portion is made to login a user
and store the token for use with the given actor.

# Features
- :white_check_mark: Gathering actor
    - :white_check_mark: Do the `Gather` action on a map tile until the goal quantity of item is reached
    - :white_check_mark: Fight a mob until the goal quantity of item is reached
- :white_check_mark: Crafting actor
    - :white_check_mark: Recursively create gathering and intermediate crafting actors to get prerequisite items
    - :clipboard: Manage inventory for higher quantity crafts
- :white_check_mark: Simulate fights
    - :white_check_mark: Basic fight simulation with restore and other utility effects
    - :clipboard: Simulation with effects from Season 4
- :white_check_mark: Fighting actor
    - :white_check_mark: Take the task of the given character and fight a mob until the task is complete
    - :clipboard: Collect a quantity of a drop from a fight
    - :white_check_mark: Heal from consumable items in inventory
        - :white_check_mark: Quit fighting if you would prefer not to rest
    - :white_check_mark: Rest between fights to recoup HP
- :clipboard: Inventory management
    - :clipboard: Deposit all items
    - :clipboard: Loadouts
