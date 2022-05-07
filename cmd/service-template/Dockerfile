FROM jrei/systemd-centos:7

COPY ProjectInfoManager.viper.yaml /etc/ProjectInfoManager/ProjectInfoManager.viper.yaml
COPY project-info-manager /usr/local/bin/project-info-manager

CMD ["project-info-manager", "run"]
