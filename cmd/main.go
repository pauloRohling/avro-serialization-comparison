package main

import (
	"encoding/json"
	"go-avro-example/internal/user"
	"os"
	"time"

	"github.com/hamba/avro/v2"
	"github.com/klauspost/compress/zstd"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	log.Logger = zerolog.New(consoleWriter).With().Timestamp().Logger()

	userSchema, err := avro.ParseFiles("./schemas/user.avsc")
	if err != nil {
		log.Fatal().Err(err).Msg("Could not read user schema")
	}

	newUser, err := user.New()
	if err != nil {
		log.Fatal().Err(err).Msg("Could not create new user")
	}

	userAvro, err := avro.Marshal(userSchema, newUser)
	if err != nil {
		log.Fatal().Err(err).Msg("Could not marshal user to avro")
	}

	userJson, err := json.Marshal(newUser)
	if err != nil {
		log.Fatal().Err(err).Msg("Could not marshal user to json")
	}

	zstdEncoder, err := zstd.NewWriter(nil, zstd.WithEncoderLevel(zstd.SpeedBetterCompression))
	if err != nil {
		log.Fatal().Err(err).Msg("Could not create zstd encoder")
	}

	userAvroCompressed := zstdEncoder.EncodeAll(userAvro, make([]byte, 0, len(userAvro)))
	userJsonCompressed := zstdEncoder.EncodeAll(userJson, make([]byte, 0, len(userJson)))

	log.Info().
		Int("size", len(userAvro)).
		Int("compressedSize", len(userAvroCompressed)).
		Msg("Generated user in Avro format")

	log.Info().
		Int("size", len(userJson)).
		Int("compressedSize", len(userJsonCompressed)).
		Msg("Generated user in JSON format")
}
