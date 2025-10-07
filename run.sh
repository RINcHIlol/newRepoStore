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
    cat > .env <<'EOF'
DB_PASSWORD=ваш_пароль_от_бд
STRIPE_PUBLISHABLE_KEY=pk_live_вставьте_сюда
STRIPE_SECRET_KEY=sk_live_вставьте_сюда
EOF
    echo "📝 Отредактируйте файл .env: добавьте ключи Stripe и пароль БД"
fi

# Загружаем переменные окружения из .env
set -a
source .env
set +a

# Запускаем приложение
echo "🌐 Запуск приложения на http://localhost:8080"
echo "📱 Откройте браузер и перейдите по адресу: http://localhost:8080"
echo ""
echo "Для остановки нажмите Ctrl+C"
echo ""

go run cmd/main.go
