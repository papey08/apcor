# apcor

*Инструмент для автоматизации сборки, запуска, отладки, 
версионирования микросервисов (и не только)*

## Пример конфигурационного файла

```yaml
version: "1.0"
name: "microservices_project"

services:
  core-service:
    remote-repo-url: 
    - https://github.com/username/microservices_project.git
    path: ./services/core-service
    pre-build:
    - go mod tidy
    build:
    - go build -o core-service
    run: ./core-service
```

## Поддерживаемые команды

* install-deps
* exec-pre-tasks
* exec-post-tasks
* pre-build
* build
* run
* debug
* clone
* checkout
* branch
* pull
