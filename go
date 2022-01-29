#!/bin/bash
# replaced by unraid driver integration: https://forums.unraid.net/bug-reports/prereleases/unraid-os-version-690-beta35-available-r1125/
#modprobe i915

# Imports module recommended from sensors
modprobe nct6775


# Automate keyfile for EFS
#/boot/config/plugins/user.scripts/scripts/downloadKeyfile/script
cp /boot/config/plugins/user.scripts/scripts/downloadKeyfile/script /tmp
chmod +x /tmp/script
/tmp/script
rm /tmp/script

chmod 777 /dev/dri/*
# Start the Management Utility
/usr/local/sbin/emhttp &
