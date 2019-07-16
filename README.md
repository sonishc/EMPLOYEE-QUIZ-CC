# quiz-vue

## Project setup
```
yarn install
go get github.com/gorilla/mux
go get github.com/jinzhu/gorm
sudo -u postgres psql postgres
CREATE ROLE quiz WITH SUPERUSER CREATEDB CREATEROLE LOGIN ENCRYPTED PASSWORD 'quiz';
create database "quiz";
```

### Compiles and hot-reloads for development
```
yarn run serve
```

### Compiles and minifies for production
```
yarn run build
```

### Run your tests
```
yarn run test
```

### Lints and fixes files
```
yarn run lint
```

### Customize configuration
See [Configuration Reference](https://cli.vuejs.org/config/).
