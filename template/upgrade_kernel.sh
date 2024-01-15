#!/bin/bash
set -e
rpm --import https://www.elrepo.org/RPM-GPG-KEY-elrepo.org
rpm -Uvh http://www.elrepo.org/elrepo-release-7.0-3.el7.elrepo.noarch.rpm
yum --disablerepo=\* --enablerepo=elrepo-kernel repolist
yum --disablerepo=\* --enablerepo=elrepo-kernel install  kernel-ml.x86_64  -y
yum remove kernel-tools-libs.x86_64 kernel-tools.x86_64  -y
yum --disablerepo=\* --enablerepo=elrepo-kernel install kernel-ml-tools.x86_64  -y
grub2-set-default 0
reboot