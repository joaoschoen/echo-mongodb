FROM mongo:latest
# INIT FILES
COPY ./init /docker-entrypoint-initdb.d
RUN chmod u+x /docker-entrypoint-initdb.d/init.js
# PORT CONFIG
EXPOSE 27017:27017