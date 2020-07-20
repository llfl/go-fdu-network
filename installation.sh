#!/bin/bash

LOOP=true

function create_conf(){
    sudo mkdir /etc/go-network && \
    echo "created config file"
}

function install(){
    sudo cp ./go-network /usr/bin/ && \
    sudo cp ./go-network.service /etc/systemd/system && \
    sudo systemctl enable go-network &&\
    sudo systemctl start go-network &&\
    echo "done!"
}

while $LOOP;do

    read -s -p $'Enter Username: \n' USRNAME
    read -s -p $'Enter Password: \n' PASS

    read -s -p $'Make sure you have correct info:[y/N]\n' CONFIRM

    if [ "$CONFIRM" = "y" ];then
        if [ "$1" = 'exist_conf' ] || [ "$1" = 'e' ];then
            sudo cp ./config.json /etc/go-network/
        else 
            create_conf
            cat << EOF | sudo tee -a /etc/go-network/config.json
{
    "username" : "${USRNAME}",
    "password" : "${PASS}"
}
EOF
        fi
        install
        # sudo echo "1 3 * * * systemctl start go-network" >> /var/spool/cron/crontabs/root
        LOOP=false
        exit 0
    else
        echo "IT'S FINE!"
    fi
done




