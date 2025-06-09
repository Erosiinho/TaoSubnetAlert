# Bittensor Subnet Alert

Ce dépôt contient un monitor d'alertes pour les subnets de Bittensor écrit en Go. Il surveille automatiquement les variations de prix des subnets spécifiés et envoie des notifications via Discord ou Twitter lorsque les seuils de variation sont atteints.

## 🚀 Fonctionnalités

- **Surveillance en temps réel** des prix des subnets Bittensor
- **Alertes personnalisables** basées sur des seuils de pourcentage de variation
- **Notifications multi-canaux** : Discord et/ou Twitter
- **Configuration flexible** des intervalles de vérification
- **Support multi-subnets** : surveillez plusieurs subnets simultanément

## 📋 Prérequis

### API Key Tao.app (OBLIGATOIRE)
Vous devez impérativement posséder une clé API sur [tao.app](https://tao.app) :

1. Connectez votre wallet sur tao.app
2. Générez une clé API dans votre dashboard
3. **Limitation** : 1000 appels/mois avec la version gratuite

### Services de notification (optionnels)
- **Discord** : Webhook URL de votre canal Discord
- **Twitter** : URL de votre service Twitter configuré

## ⚙️ Configuration

Créez un fichier `.env` à la racine du projet avec les variables suivantes :

```env
# Service principal (requis)
SERVICE=twitter or discord
API_KEY=votre_cle_api_tao_app

# URLs des services de notification (au moins un requis)
TWITTER_SERVICE_URL=https://votre-service-twitter.com/webhook
DISCORD_SERVICE_URL=https://discord.com/api/webhooks/your/webhook/url

# Configuration de surveillance
SUBNET_IDS=1,5,18,32
PERCENT_THRESHOLD=5.0
CHECK_INTERVAL_MINUTES=15
```

## 🚀 Installation et lancement

### Clone

```bash
git clone https://github.com/Erosiinho/TaoSubnetAlert.git
```

```bash
cd TaoSubnetAlert/
```

### Docker compose

Fill .env variables, then run :

```bash
docker-compose up -d
```