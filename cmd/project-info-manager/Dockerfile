FROM jrei/systemd-centos:7

WORKDIR /app

COPY project-info-manager /app
COPY ProjectInfoManager.viper.yaml /app

ENTRYPOINT ["./project-info-manager"]

EXPOSE 50310 50311 50312

CMD ["run"]
