#!/bin/bash
if [[ "$OSTYPE" == "linux-gnu" ]]; then
        echo "This is Linux"
elif [[ "$OSTYPE" == "darwin"* ]]; then
        echo "This is MacOS"
	brew install go
fi
