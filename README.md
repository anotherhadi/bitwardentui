# Bitwarden TUI

A read-only TUI for accessing Bitwarden vault contents from the terminal. Built on top of the bitwarden-cli project.

*WARNING*: This application is experimental and has not been audited. Use at your own risk.

*This project is not associated with the Bitwarden project nor Bitwarden Inc.*

## Todolist

### Functions

- [x] loadVault
- [x] login
- [x] logout
- [x] lock
- [x] unlock
- [x] sync
- [ ] generate password
- [ ] list
- [ ] list folder
- [ ] get
- [ ] get folder
- [ ] search

### UI (V1 is read-only)

- [ ] Login page
- [ ] Locked page
- [ ] Panel
  - Top: search bar
  - Left: folder structure, user/server informations
  - Right: password list
  - Bottom: help
- [ ] Password generation page
  - password rules
  - refresh

### V2

- add
- edit
- move
- delete
- create
- get nerdfont icon for website
- configuration files (show icons, default generation config, colors)
- item page
- save token to env var
- choose copy functions
