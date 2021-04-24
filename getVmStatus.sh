#!/bin/bash
virsh list --all | grep " Windows10 " | awk '{ print $3}'