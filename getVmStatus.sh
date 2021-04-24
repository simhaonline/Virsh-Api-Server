#!/bin/bash
virsh list --all | grep " $1 " | awk '{ print $3}'