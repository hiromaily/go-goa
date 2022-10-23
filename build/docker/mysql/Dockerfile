FROM mysql:5.7

COPY ./docker/mysql/conf.d /etc/mysql/conf.d
COPY ./docker/mysql/init.d /docker-entrypoint-initdb.d

ENV MYSQL_ROOT_PASSWORD=root
ENV MYSQL_DATABASE=go-goa
ENV MYSQL_USER=hiromaily
ENV MYSQL_PASSWORD=secret123

VOLUME /var/lib/mysql

EXPOSE 3306

CMD ["mysqld"]