Service: User Provider
======================

## Для фетчинга graphql модели с удаленного репозитория(пример: [уд. репозиторий](https://github.com/first-debug/lk-graphql-schemas) )
```
go run github.com/first-debug/lk-auth/cmd/schema-fetcher https://raw.githubusercontent.com/first-debug/lk-graphql-schemas/refs/heads/master/schemas/user-provider/schema.graphql graph/schema.graphqls
```

+ https://raw.githubusercontent.com/first-debug/lk-graphql-schemas/refs/heads/master/schemas/user-provider/schema.graphql - сырая ссылка на graphql схему из удаленного репозитория

+ graph/schema.graphqls - путь, в который будет скачиваться модель из удаленного репозитория. (Путь указывается относительно директории, где была запущено)

## После успешного получения graphql модели пользуемся gqlgen. 
### Генерация кода по модели
```
go run github.com/99designs/gqlgen generate
```
