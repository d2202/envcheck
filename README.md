## Описание
envcheck - cli утилита, которая сравнивает файл эталонный .env.example с актуальным .env и находит расхождения между ними.

## Установка
Требования: go >= 1.21
```bash
go install github.com/d2202/envcheck@latest
```

## Использование
Базовый сценарий:
```bash
envcheck --example .env.example --env .env
```

"Тихий режим" (отдает только os.ExitCode)
```bash
envcheck --example .env.example --env .env --quiet
```

"Строгий режим" - ExitCode = 1 только в случае нехватки ключа в сравниваемом .env, остальные - 0
```bash
envcheck --example .env.example --env .env --strict
```

## Exit codes
0 - Успех - в сравниваемом файле нет отсутствующих ключей
1 - Провал - в сравниваемом файле найдены остутствующие ключи
2 - Предупреждение - есть некритичные расхождения

## Пример вывода
```bash
$ ./envcheck --env .env --example .env.example

Comparing:
        example : .env.example (1 keys)
        env     : .env (1 keys)
⚠️       EXTRA (1)       — keys in env, not in example
SOMEKEY
❌      MISSING (1)     — keys in example, not in env
KEY
Result: 1 missing keys.
```

## Тестирование
```bash
go test ./...
```
