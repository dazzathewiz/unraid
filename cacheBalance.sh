#!/bin/bash
btrfs balance start -dusage=80 /mnt/cache
btrfs balance start -dusage=80 /mnt/disks/nvme0