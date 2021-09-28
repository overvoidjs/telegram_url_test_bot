# Telegram URL-Test Bot
- данный бот реализует уведосления о недоступности URL адреса из списка
------------

## Установка
### UBUNTU:
Код писался под убунту, по этому если будете запускать на другой ОС - проверьте что все работает ок.
- Переименовываем файлы **demo.env** и **demo.testlist** в **.env** и **.testlist**
- В **.env** заполняем поля с:  и ID-чата в который добавлен бот, должно получиться что то вроде, только с вашими ключами:
-- **ID-бота**( _0000000000:ХХХХХХХХХХХХХ_ )
-- **ID-чата** в который добавлен бот и куда будет присылать уведомления
-- **TIMEOUT** - время между запросами в **секундах**
-- **LOOP_TIME** - время между запусками проверки в **минутах**
 ```
BOT_ID=0000000000:ХХХХХХХХХХХХХ
CHAT_ID=-0000000000000
TIMEOUT=1
LOOP_TIME=60
 ```
- В **.testlist** заполняем URL адреса которые хотим тестировать с переносом строки, должно получиться что-то вроде:
```
https://www.google.com/
https://yandex.ru/
```
- Запускать можно через ```./bot &``` что бы сразу отпустило консоль и бот спокойно работал в режиме демона
