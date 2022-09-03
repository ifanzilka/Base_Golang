#!/bin/bash
if [[ "$OSTYPE" == "linux-gnu" ]]; then
        echo "This is Linux"
	sudo apt update -y
	sudo apt upgrade -y
	sudo apt install -y golang
elif [[ "$OSTYPE" == "darwin"* ]]; then
        echo "This is MacOS"
	brew install go
fi
