This is my application for playing the game [ArtifactsMMO](https://artifactsmmo.com/). The `pkg` directory is meant to wrap around [promiseofcake's artifactsmmo-go-client](https://github.com/promiseofcake/artifactsmmo-go-client).
The application is split into two binaries, a CLI client and a combination of DiscordBot/web app.

# CLI Client
The CLI client was made to quickly test different facets of the application in REPL format. There is
a concept of actors which will do a task to completion (or error).

# DiscordBot
The DiscordBot allows to create actors from bot commands. The web app portion is made to login a user
and store the token for use with the given actor.

# Features
- [x] Gathering actor
    - [x] Do the `Gather` action on a map tile until the goal quantity of item is reached
    - [ ] Fight a mob until the goal quantity of item is reached
- [x] Crafting actor
    - [x] Recursively create gathering and intermediate crafting actors to get prerequisite items
    - [ ] Manage inventory for higher quantity crafts
- [x] Simulate fights
    - [x] Basic fight simulation with restore and other utility effects
    - [ ] Simulation with effects from Season 4
- [x] Fighting actor
    - [x] Take the task of the given character and fight a mob until the task is complete
    - [ ] Collect a quantity of a drop from a fight
    - [x] Heal from consumable items in inventory
        - [x] Quit fighting if you would prefer not to rest
    - [x] Rest between fights to recoup HP
- [ ] Inventory management
    - [ ] Deposit all items
    - [ ] Loadouts
