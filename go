#!/bin/bash
modprobe i915

# Imports module for MB sensors (such as temp and PWM fans)
modprobe it87

# Automate keyfile for EFS
#/boot/config/plugins/user.scripts/scripts/downloadKeyfile/script
cp /boot/config/plugins/user.scripts/scripts/downloadKeyfile/script /tmp
chmod +x /tmp/script
/tmp/script
rm /tmp/script

chmod 777 /dev/dri/*
# Start the Management Utility
/usr/local/sbin/emhttp &
