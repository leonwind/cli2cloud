FROM postgres
ENV POSTGRES_DB $POSTGRES_DB
ENV POSTGRES_USER $POSTGRES_USER
ENV POSTGRES_PASSWORD $POSTGRES_PASSWORD
COPY create_schema.sql /create_schema.sql
COPY start_postgres.sh /docker-entrypoint-initdb.d
EXPOSE 5432
