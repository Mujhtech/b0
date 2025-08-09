#!/bin/bash

# Setup script for B0 with Traefik

set -e

echo "🚀 Setting up B0 with Traefik..."

# Create external network if it doesn't exist
if ! docker network ls | grep -q "web"; then
    echo "📡 Creating external network 'web'..."
    docker network create web
else
    echo "✅ Network 'web' already exists"
fi

# Create necessary directories
echo "📁 Creating directories..."
mkdir -p traefik/letsencrypt traefik/config

# Set proper permissions for Let's Encrypt directory
chmod 600 traefik/letsencrypt

echo "🔧 Starting Traefik..."
docker-compose -f traefik-compose.yaml up -d

echo "⏳ Waiting for Traefik to be ready..."
sleep 5

echo "🏗️  Starting main services..."
docker-compose -f docker-compose.dev.yaml up -d

echo "✅ Setup complete!"
echo ""
echo "🌐 Available services:"
echo "  - Traefik Dashboard: http://traefik.localhost:8080"
echo "  - Frontend: http://b0-frontend.local"
echo "  - Backend API: http://b0-backend.local"
echo ""
echo "📝 Make sure to add these entries to your /etc/hosts file:"
echo "127.0.0.1 traefik.localhost"
echo "127.0.0.1 b0-frontend.local"
echo "127.0.0.1 b0-backend.local"
