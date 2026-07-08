## Описание
envcheck - cli утилита, которая сравнивает файл эталонный файл настроек с актуальным файлом и находит расхождения между ними.
### Поддерживаемые форматы:
- .env.example -> .env
- .toml.example -> .toml
- .yaml.example -> .yaml


## Установка
Требования: go >= 1.21
```bash
go install github.com/d2202/envcheck@latest
```

## Использование
Базовый сценарий:
```bash
envcheck --actual .env --expected .env.example
```
или
```bash
envcheck --actual test.toml --expected test.toml.example --format toml
```

"Тихий режим" (отдает только os.ExitCode)
```bash
envcheck --actual .env --expected .env.example --format env --quiet
```

"Строгий режим" - ExitCode = 1 только в случае нехватки ключа в сравниваемом файле, остальные - 0
```bash
envcheck --actual .env --expected .env.example --format env --strict
```

JSON - форматирование:
```bash
envcheck --actual .env --expected .env.example --format env --json
```

## Exit codes
0 - Успех - в сравниваемом файле нет отсутствующих ключей
1 - Провал - в сравниваемом файле найдены отсутствующие ключи
2 - Предупреждение - есть некритичные расхождения

## Пример вывода
```bash
$ ./envcheck --actual tests/env/.env --expected tests/env/.env.example

Comparing:
        example : tests/env/.env.example (2 keys)
        env     : tests/env/.env (2 keys)
⚠️       EXTRA (1)       — keys in ACTUAL, not in EXPECTED
EXTRAKEY_NAME
❌      MISSING (1)     — keys in EXPECTED, not in ACTUAL
MISSINGKEY_NAME
Result: 1 missing keys.
```

```bash
$ ./envcheck --actual tests/env/.env --expected tests/env/.env.example --json

{"missing":["MISSINGKEY_NAME"],"extra":["EXTRAKEY_NAME"],"empty":[],"ok":false}
```

## Тестирование
```bash
go test ./...
```
