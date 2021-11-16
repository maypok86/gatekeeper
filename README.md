# Gatekeeper
![example workflow](https://github.com/maypok86/gatekeeper/actions/workflows/test.yml/badge.svg)
![Go Report](https://goreportcard.com/badge/github.com/maypok86/gatekeeper)
![License](https://img.shields.io/badge/license-MIT-green)

## Общее описание
Сервис предназначен для борьбы с подбором паролей при авторизации в какой-либо системе.
Сервис вызывается перед авторизацией пользователя и может либо разрешить, либо заблокировать попытку.
Предполагается, что сервис используется только для server-server, т.е. скрыт от конечного пользователя.

## Алгоритм работы
Сервис ограничивает частоту попыток авторизации для различных комбинаций параметров, например:
* не более `RATE_LOGIN = 10` попыток в минуту для данного логина.
* не более `RATE_PASSWORD = 100` попыток в минуту для данного пароля (защита от обратного brute-force).
* не более `RATE_IP = 1000` попыток в минуту для данного IP (число большое, т.к. NAT).

White/black листы содержат списки адресов сетей, которые обрабатываются более простым способом.
Если входящий ip в whitelist - сервис безусловно разрешает авторизацию (ok=true), если в blacklist - отклоняет (ok=false).

Для работы используется пакет [time/rate](https://pkg.go.dev/golang.org/x/time/rate), в котором реализован алгоритм [token bucket](https://en.wikipedia.org/wiki/Token_bucket).

## Описание методов API

### Попытка авторизации
Запрос:
* login
* password
* ip

Ответ:
* ok (true/false) - сервис должен возвращать ok=true, если считает что запрос нормальный
  и ok=false, если считает что происходит bruteforce.

### Сброс bucket по login
* login

Должен очистить bucket-ы соответствующие переданному login.

### Сброс bucket по IP
* ip

Должен очистить bucket-ы соответствующие переданному ip.

### Добавление IP в blacklist
* subnet (ip + mask)

### Удаление IP из blacklist
* subnet (ip + mask)

### Добавление IP в whitelist
* subnet (ip + mask)

### Удаление IP из whitelist
* subnet (ip + mask)

## Command-Line интерфейс
Разработан command-line интерфейс для ручного администрирования сервиса.
Через CLI есть возможность вызвать сброс бакета и управлять whitelist/blacklist-ом.
CLI работает через GRPC интерфейс.
