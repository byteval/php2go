**Структура каталогов**:
```
notification/
├── cmd/
│   └── notification-server/
│       └── main.go                     # Инициализация конфига, логгера, DI-контейнера
├── config/                             # Конфигурация
│   └── config.go                       # Структура конфига проекта (параметры базы данных, SMTP, IMAP, Websocket, Webhook)
├── docs/                               # Swagger
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── migrations/                         # Миграции
├── internal/
│    ├── app/                           # Операции / Use Cases
│    │   ├── notification/              # Уведомления
│    │   │   ├── create/                # Создание уведомления
│    │   │   │   ├── creator.go         # Создание уведомления
│    │   │   │   ├── notifier.go        # Отправка уведомления
│    │   │   │   ├── mapper.go          # Маппинг ToDomain/ToResponse
│    │   │   │   └── dto.go             # DTO Request/Response
│    │   │   └── get/                   # Получение уведомления
│    │   │       ├── getter.go          
│    │   │       ├── mapper.go          
│    │   │       └── dto.go             
│    │   └── layout/                    # Шаблоны уведомлений
│    │       ├── create/
│    │       │   ├── creator.go
│    │       │   ├── validator.go
│    │       │   └── dto.go
│    │       ├── update/
│    │       │   ├── updater.go
│    │       │   └── dto.go
│    │       ├── get/
│    │       │   ├── getter.go     
│    │       │   └── dto.go
│    │       ├── list/
│    │       │   ├── lister.go
│    │       │   └── dto.go
│    │       └── delete/
│    │           └── deleter.go
│    ├── container/            
│    │   └── container.go               # Контейнер DI
│    ├── domain/     
│    │   ├── notification/              
│    │   │     ├── errors.go            # Структура с описанием ошибок
│    │   │     ├── notification.go      # Модель (Notification)
│    │   │     └── ports/               # Интерфейсы сервисов
│    │   │         ├── repository.go    # Интерфейс репозитория Notification
│    │   │         ├── smtp.go          # Отправка уведомления SMTP
│    │   │         ├── webhook.go       # Отправка уведомления webhook
│    │   │         └── websocket.go     # Отправка уведомления websocket
│    │   └── layout/                    # Шаблоны
│    │       ├── layout.go              # Модель (Layout)
│    │       ├── errors.go
│    │       └── ports/
│    │           ├── repository.go
│    │           └── renderer.go        # Интерфейс для рендеринга
│    ├── transport/
│    │   └── http/
│    │       ├── handlers/              # Обработчики HTTP запросов
│    │       │   ├── healthcheck.go
│    │       │   ├── notifications/
│    │       │   │   ├── create.go      # POST /notifications
│    │       │   │   └── get.go         # GET /notifications/{id}    
│    │       │   └── layouts/           
│    │       │       ├── create.go      # POST /layouts
│    │       │       ├── update.go      # PUT /layouts/{id}
│    │       │       ├── get.go         # GET /layouts/{id}
│    │       │       ├── list.go        # GET /layouts
│    │       │       ├── delete.go      # DELETE /layouts/{id}
│    │       │       └── render.go      # POST /layouts/{id}/render
│    │       ├── middleware/
│    │       │   ├── middleware.go   
│    │       │   └── validator.go   
│    │       └── router.go
│    ├── errors/      
│    ├── infrastructure/                # Реализации (БД, SMTP и т.д.) technical details
│    │   ├── postgres/
│    │   │   ├── database.go            # Подключение к СУБД POSTGRES
│    │   │   ├── notification_repository.go
│    │   │   └── layout_repository.go   
│    │   └── notifiers/                 # Реализации сервисов
│    │       ├── smtp/                  # Сервис отправки почты по SMTP
│    │       │   ├── client.go
│    │       │   └── notifier.go
│    │       └── webhook/               # Сервис отправки webhook
│    │           ├── client.go
│    │           └── notifier.go
├── pkg/                                # Дополнительные пакеты
│    └── logger 
├── test/ 
```
