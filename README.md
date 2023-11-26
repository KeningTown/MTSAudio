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

# Preview
<video width="320" height="240" controls>
  <source src="Preview.mp4" type="video/mp4">
  Ваш браузер не поддерживает тег video.
</video>