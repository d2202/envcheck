# Changelog

## v2.0.0 
### Breaking changes
- Флаги `--env`/`--example` переименованы в `--actual`/`--expected`
- Добавлен обязательный флаг `--format` (env/toml/yaml) default - `env`
### Added
- Поддержка TOML файлов
- Поддержка YAML файлов
- JSON вывод теперь возвращает `[]` вместо `null` для пустых полей

## v1.1.0
### Added
- Вывод результата в JSON

## v1.0.0
### Added
- Тесты Reporter
- README.md
- GitHub CI

## v0.3.0
### Added
- Логика сравнения итоговых мап в Checker
- Тесты Checker
- Reporter логика + вывод

## v0.2.0
### Added
- Тесты на parseLine, Parse
### Fixed
- Обработка ошибочного сценария scanner.Err() не учитывала путь

## v0.1.1
### Added
- Поддержка кавычек в значениях .env
### Fixed
- Корректная обработка пустых строк
- Корректная обработка закомментированных строк и inline-комментариев
- Корректная обработка строк формата "export KEY=VALUE"

## v0.1.0
### Added
- Базовый парсинг .env файлов