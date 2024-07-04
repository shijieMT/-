mysqldump -u root -p notify > db.sql
mysqldump -u root -p --default-character-set=utf8mb4 notify > db.sql
    volumes:
      - ./mysql:/docker-entrypoint-initdb.d