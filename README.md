# MTSAudio
Весь рабочий и не слишком рабочий код располагается в ветке `main`.

## Server start
Для запуска сервера выполните следующие команды:
```bash
cd server
docker-compose up
```

## Mobile start
Следует открыть директорию `Mobile` в AndroidStudio.
Для запуска мобильной версии приложения выполните следующие команды:
```bash
Build -> Clean Project
Build -> Rebuild Project
```
Запускать эмуляторы на ***одном устройстве с сервером***

## Desktop start
Для запуска десктопной версии приложения выполните следующие команды:
```bash
cd desktop
npm install
cd mts-app
ng serve --open
```
