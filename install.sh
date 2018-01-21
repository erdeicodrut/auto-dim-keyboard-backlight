#!/bin/bash

sudo cp kb /usr/bin
sudo cp keyboard.service /etc/systemd/system/

sudo systemctl enable keyboard.service
sudo systemctl start keyboard.service
