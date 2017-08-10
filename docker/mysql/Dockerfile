FROM mysql:5.7

COPY ./conf.d /etc/mysql/conf.d
COPY ./init.d /docker-entrypoint-initdb.d

ENV MYSQL_ROOT_PASSWORD=root
ENV MYSQL_DATABASE=hiromaily
ENV MYSQL_USER=hiromaily
ENV MYSQL_PASSWORD=12345678

VOLUME /var/lib/mysql

EXPOSE 3306