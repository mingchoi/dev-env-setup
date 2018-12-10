# Development Environment Setup
This repo contain shell script and interactive task runner to help you setup a freshly installed ubuntu to a production ready development environment quickly.

You can fork and add your shell script for you need. The `./dev-env-setup`(main.go) binary will detect any `.sh` in the folder and ask for what you want to install, and generate a `install.sh` of your selected script. Then type `sh install.sh` to start installation.


## Quick Start
```
sh start.sh
```

### Included Install Script
* nodejs - Node.js
* git - Version Control
* docker - Docker-CE container platform
* golang - Go language
* golang-lib - Useful go library like gin, viper and more...
* mariadb-in-docker - SQL Database (install in docker)
* gnome - Vanilla Gnome (for who don't like Ubuntu Dock)
* elementary-desktop - Column view finder

### Manual Install
```
sh script_name.sh
```

### Others: Software install by software center
You may refer to [software.md](software.md) for sugggest tools.