# Development Environment Setup
This repo contain shell script and interactive task runner to help you setup a freshly installed ubuntu to a production ready development environment quickly.

You can fork and add your shell script for you need. The `./dev-env-setup`(main.go) binary will detect any `.sh` in the folder and ask for what you want to install, and generate a `install.sh` of your selected script. Then type `sh install.sh` to start installation.


### Quick Start
```
sh start.sh
```

### Manual Install
```
sh script_name.sh
```

### Software install by software center
You may refer to [software.md](software.md) for sugggest tools.