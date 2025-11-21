#!/bin/bash
#
# Setup cron job for LenaLink data synchronization
#
# Usage:
#   ./scripts/setup-cron.sh
#
# This script will:
# 1. Add cron job to sync data every 6 hours
# 2. Create log directory
# 3. Show current crontab

set -e

# Colors for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${BLUE}========================================${NC}"
echo -e "${BLUE}LenaLink Cron Setup${NC}"
echo -e "${BLUE}========================================${NC}\n"

# Get the project directory (parent of scripts/)
PROJECT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
echo -e "${GREEN}✓${NC} Project directory: $PROJECT_DIR"

# Create log directory
LOG_DIR="$PROJECT_DIR/logs"
mkdir -p "$LOG_DIR"
echo -e "${GREEN}✓${NC} Log directory: $LOG_DIR"

# Cron schedule (every 6 hours at minute 0)
CRON_SCHEDULE="0 */6 * * *"

# Cron command
CRON_COMMAND="cd $PROJECT_DIR && docker compose run --rm seed >> $LOG_DIR/sync.log 2>&1"

# Full cron entry
CRON_ENTRY="$CRON_SCHEDULE $CRON_COMMAND"

echo -e "\n${YELLOW}Cron entry to be added:${NC}"
echo -e "  $CRON_ENTRY\n"

# Check if cron entry already exists
if crontab -l 2>/dev/null | grep -F "docker compose run --rm seed" > /dev/null; then
    echo -e "${YELLOW}⚠${NC}  Cron job already exists. Skipping..."
    echo -e "\nTo remove existing cron job:"
    echo -e "  crontab -e"
    echo -e "  # Then delete the line containing 'docker compose run --rm seed'\n"
else
    # Add to crontab
    (crontab -l 2>/dev/null; echo "$CRON_ENTRY") | crontab -
    echo -e "${GREEN}✓${NC} Cron job added successfully!"
fi

# Show current crontab
echo -e "\n${BLUE}Current crontab:${NC}"
crontab -l | grep -v "^#" | grep -v "^$" || echo "  (empty)"

# Create log rotation config
LOGROTATE_CONFIG="/etc/logrotate.d/lenalink-sync"
echo -e "\n${YELLOW}Optional: Log rotation${NC}"
echo -e "To rotate logs automatically, create $LOGROTATE_CONFIG with:"
cat <<EOF

$LOG_DIR/sync.log {
    daily
    rotate 7
    compress
    delaycompress
    missingok
    notifempty
    create 0640 $(whoami) $(whoami)
}
EOF

echo -e "\n${BLUE}========================================${NC}"
echo -e "${GREEN}✓${NC} Setup complete!"
echo -e "${BLUE}========================================${NC}\n"

echo -e "Sync schedule: ${GREEN}Every 6 hours${NC}"
echo -e "Log file:      ${GREEN}$LOG_DIR/sync.log${NC}"
echo -e "\nManual commands:"
echo -e "  Run sync now:        ${GREEN}docker compose run --rm seed${NC}"
echo -e "  View logs:           ${GREEN}tail -f $LOG_DIR/sync.log${NC}"
echo -e "  List cron jobs:      ${GREEN}crontab -l${NC}"
echo -e "  Edit cron jobs:      ${GREEN}crontab -e${NC}"
echo -e ""
