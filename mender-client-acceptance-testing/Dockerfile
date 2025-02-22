FROM ubuntu:18.04

ENV DEBIAN_FRONTEND=noninteractive

# prepare NFS cache for yocto
RUN mkdir -p /mnt/sstate-cache

# install pyyaml via pip to workardound a later conflict if installed via OS package (awscli depends on python3-yaml)
RUN apt-get update -qq && apt-get install -yyq python3-pip && pip3 install --upgrade pip && pip3 install pyyaml

# Get OS requirements from master
RUN apt-get install -yyq wget && \
    wget https://raw.githubusercontent.com/mendersoftware/meta-mender/master/tests/acceptance/requirements_deb.txt && \
    apt-get install -yyq $(cat requirements_deb.txt) && apt-get install -yq linux-modules-`uname -r`

# Locales
RUN locale-gen --purge en_US.UTF-8
ENV LC_ALL=en_US.UTF-8 LANG=en_US.UTF-8 LANGUAGE=en_US.UTF-8

# Python 3 requirements from master
RUN wget https://raw.githubusercontent.com/mendersoftware/meta-mender/master/tests/acceptance/requirements_py3.txt && \
    pip3 install -r requirements_py3.txt && \
    rm requirements_py3.txt

# Python 2 requirements from warrior
RUN apt-get install -y python-pip && \
    pip2 install --upgrade pip && \
    pip2 install requests --upgrade && \
    pip2 install pytest --upgrade && \
    pip2 install filelock --upgrade && \
    pip2 install pytest-xdist --upgrade && \
    pip2 install pytest-html --upgrade && \
    pip2 install -I fabric==1.14.0 && \
    pip2 install psutil --upgrade && \
    pip2 install boto3 --upgrade && \
    pip2 install pycrypto --upgrade && \
    wget https://raw.githubusercontent.com/mendersoftware/meta-mender/warrior/tests/acceptance/requirements.txt && \
    pip2 install -r requirements.txt && \
    rm requirements.txt

# golang from upstream
RUN wget -q https://storage.googleapis.com/golang/go1.12.linux-amd64.tar.gz && gunzip -c go1.12.linux-amd64.tar.gz | (cd /usr/local && tar x) && ln -sf /usr/local/go/bin/go /usr/local/bin/go &&  ln -sf /usr/local/go/bin/godoc /usr/local/bin/godoc && ln -sf /usr/local/go/bin/gofmt /usr/local/bin/gofmt

# mender user dir
RUN useradd -m -u 1010 mender && usermod -a -G kvm mender && usermod -a -G docker mender

# Prepare mender user
RUN mkdir -p /home/mender/.ssh && echo "mender ALL=(ALL:ALL) NOPASSWD:ALL" >> /etc/sudoers && sed -i -e 's/^\( *Defaults *requiretty *\)$/# \1/' /etc/sudoers && chown -R mender:mender /home/mender

# debugfs
RUN cp /sbin/debugfs /usr/bin/ || echo "debugfs not in /sbin/"

# enable sysstat monitoring suite for Debian/Ubuntu
RUN sed -i 's/false/true/g' /etc/default/sysstat
