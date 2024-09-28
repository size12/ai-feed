# AI-feed - сервис генерации статей с редактированием личности автора


### Запуск
```shell
git clone https://github.com/size12/ai-feed && cd ai-feed
vim config.yaml # редактирование конфигурации, подробнее ниже
docker-compose up -d
```

### Возможности
+ Генерация статей на актуальные темы
+ Возможность публикация статей в открытый доступ
+ Создание и редактирование личностей авторов
+ Парсинг актуальных новостей с сайтов [vc.ru](https://vc.ru) и [ixbt.com](https://ixbt.com)
+ Возможность генерации на разных языках
+ Генерации избражений с помощью DALL-E
+ Наличие API

### Пример статьи
> Была использована личность Александра Пушкина и стихотворный стиль речи

![article_example.png](images/article_example.png)

### Параметры конфигурации

```shell
go run cmd/feed/main.go -cfg=config.yaml
# Это флаг файла конфигурации, по умолчанию 'config.yaml'
```


```yaml
# config.yaml

app:
  run_port: ":8080" # на каком порту запускается сервис
ai:
  open_ai_endpoint: "" # Эндпоинт OpenAI API, например https://api.proxyapi.ru/openai/v1 или https://api.openai.com/v1
  open_ai_auth_token: "" # Токен авторизациии OpenAI API, вида 'sk-xxx...'
  model_type: "gpt-4o" # LLM модель

  image_prompt_path: "./prompts/image_generation.txt" # Промпт для генерации промпта картинки
  text_prompt_path: "./prompts/text_generation.txt"  # промпт для генерации самой статьи
  title_prompt_path: "./prompts/title_generation.txt"  # промпт для генерации заголовка статьи

storage:
  postgres_endpoint: "" # Эндпоинт для Postgres SQL, например postgresql://username:password@127.0.0.1:5432/ai-feed
  themes_actual_count: 15 # Количество актуальных тем (остальные удаляются как неактуальные)
  worker_actual_update: 10min # Каждое время worker_actual_update оставляется только themes_actual_count тем

feeder:
  feed_update_delay: 1h # Время периода обновления новых тем (каждый раз спустя feed_update_delay парсер собирает актуальные темы)
```