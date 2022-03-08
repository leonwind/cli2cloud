FROM postgres

ENV POSTGRES_DB cli2cloud 
ENV POSTGRES_USER cli2cloud 
ENV POSTGRES_PASSWORD $POSTGRES_PASSWORD

COPY create_schema.sql /docker-entrypoint-initdb.d 
#COPY start_postgres.sh /docker-entrypoint-initdb.d

EXPOSE 5432
