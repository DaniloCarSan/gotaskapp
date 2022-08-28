# :runner: GUIA DE INSTALAÇÃO


1 - Configurar as variáveis de ambiente, criando uma copia 
do arquivo **.env_example** e remomeando para **.env**

````
APP_HOST=
APP_PORT=

SENTRY_DNS=

DB_HOST=
DB_PORT=
DB_NAME=
DB_USER=
DB_PASS=

JWT_SECRET=

EMAIL_HOST=
EMAIL_FROM=
EMAIL_PASSWORD=
EMAIL_PORT=
````

2 - Para configurar o banco, copie o template na pasta
````
./docker/mysql/scripts/migrate.sql
````
e rode o mesmo em um SGBD, mas caso possua o docker instado
em sua maquina rode o comando na raiz do projeto
````
docker compose up
````
ele criára seu banco e fará o migrate automatico das tabelas e
dos dados padrão, caso queira editar os dados que são inseridos
nas tabelas por padrão edite o arquivo **migrate.sql**.

3 - Nas rais do projeto baixe as dependencias do projeto de um start com o comando

````
go run .
````
