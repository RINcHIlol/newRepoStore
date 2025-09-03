#!/bin/bash

echo "🚀 Запуск Store API..."

# Проверяем, что Go установлен
if ! command -v go &> /dev/null; then
    echo "❌ Go не установлен. Установите Go 1.19+"
    exit 1
fi

# Проверяем версию Go
GO_VERSION=$(go version | awk '{print $3}' | sed 's/go//')
echo "✅ Go версия: $GO_VERSION"

# Устанавливаем зависимости
echo "📦 Установка зависимостей..."
go mod download

# Проверяем наличие .env файла
if [ ! -f .env ]; then
    echo "⚠️  Файл .env не найден. Создаю пример..."
    echo "DB_PASSWORD=ваш_пароль_от_бд" > .env
    echo "📝 Отредактируйте файл .env с вашим паролем от базы данных"
fi

# Запускаем приложение
echo "🌐 Запуск приложения на http://localhost:8080"
echo "📱 Откройте браузер и перейдите по адресу: http://localhost:8080"
echo ""
echo "Для остановки нажмите Ctrl+C"
echo ""

go run cmd/main.go
