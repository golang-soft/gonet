package mredis

const (
	KEYS_game_global        = "game:global"
	KEYS_game_round         = "game:round:basic:"
	KEYS_user_round_basic   = "user:round:basic:"
	KEYS_user_round_players = "user:round:players:"
)

var PREFIX = "hero"

var REDIS_KEYS map[string]string = map[string]string{
	KEYS_game_global:        PREFIX + ":" + KEYS_game_global,
	KEYS_game_round:         PREFIX + ":" + KEYS_game_round,
	KEYS_user_round_basic:   PREFIX + ":" + KEYS_user_round_basic,
	KEYS_user_round_players: PREFIX + ":" + KEYS_user_round_players,
}
