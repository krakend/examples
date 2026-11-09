#!/bin/bash
# Simulates pilots playing Earth vs Aliens so the dashboard has something to show.
# Each game: the pilot inserts a coin, picks a ship, fires power-ups and asks for bomb reloads.
# Quotas are small on purpose: you will see plenty of 429 (game over) responses.
#
# Usage: ./play.sh [GAMES] [KRAKEND_ADDRESS]   (GAMES=0 plays forever, Ctrl+C to stop)

GAMES="${1:-30}"
KRAKEND_ADDRESS="${2:-http://localhost:8080}"

PILOTS=(bart lisa marge homer milhouse nelson ralph)
SHIPS=(ala-x ala-b ala-y)

call() { # pilot ship endpoint
  curl -s -o /dev/null -w '%{http_code}' \
    -H "Authorization: $1" -H "X-Ship: $2" "${KRAKEND_ADDRESS}$3"
}

game=0
while [ "$GAMES" -eq 0 ] || [ "$game" -lt "$GAMES" ]; do
  game=$((game + 1))
  pilot=${PILOTS[$((RANDOM % ${#PILOTS[@]}))]}
  ship=${SHIPS[$((RANDOM % ${#SHIPS[@]}))]}

  coin=$(call "$pilot" "$ship" /consume_credit)
  if [ "$coin" != "200" ]; then
    printf '#%-4s %-9s %-6s INSERT COIN... no credits left (%s)\n' "$game" "$pilot" "$ship" "$coin"
    sleep 0.5
    continue
  fi

  powerups=0 bombs=0 over=""
  for _ in $(seq 1 $((RANDOM % 4 + 1))); do
    [ "$(call "$pilot" "$ship" /request_weapon_power_up)" = "200" ] && powerups=$((powerups + 1)) || over="power-ups depleted"
  done
  for _ in $(seq 1 $((RANDOM % 3))); do
    [ "$(call "$pilot" "$ship" /request_bomb_reload)" = "200" ] && bombs=$((bombs + 1)) || over="${over:+$over, }bomb bay empty"
  done

  printf '#%-4s %-9s %-6s power-ups:%s bomb reloads:%s %s\n' "$game" "$pilot" "$ship" "$powerups" "$bombs" "${over:+[$over]}"
  sleep "0.$((RANDOM % 9 + 1))"
done
