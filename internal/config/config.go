package config

import (
	"github.com/BurntSushi/toml"
	"os"
	"path"
)

type Config struct {
	Board
	Herbivore
	Predator
	Grass
	Rock
	Tree
}

type Board struct {
	Rows         int `toml:"board_rows"`
	Columns      int `toml:"board_columns"`
	DelaySeconds int `toml:"board_delay_seconds"`
}

type Herbivore struct {
	Speed                   int     `toml:"herbivore_speed"`
	Hp                      int     `toml:"herbivore_hp"`
	MaxStepsHungryBeforeDie int     `toml:"herbivore_max_steps_hungry_before_die"`
	NutritionValue          int     `toml:"herbivore_nutrition_value"`
	Sign                    string  `toml:"herbivore_sign"`
	SpawnRate               float64 `toml:"herbivore_spawn_rate"`
}

type Predator struct {
	Speed                   int     `toml:"predator_speed"`
	Hp                      int     `toml:"predator_hp"`
	MaxStepsHungryBeforeDie int     `toml:"predator_max_steps_hungry_before_die"`
	AttackPower             int     `toml:"predator_attack_power"`
	Sign                    string  `toml:"predator_sign"`
	SpawnRate               float64 `toml:"predator_spawn_rate"`
}

type Grass struct {
	NutritionValue int     `toml:"grass_nutrition_value"`
	Sign           string  `toml:"grass_sign"`
	SpawnRate      float64 `toml:"grass_spawn_rate"`
}

type Rock struct {
	Sign      string  `toml:"rock_sign"`
	SpawnRate float64 `toml:"rock_spawn_rate"`
}

type Tree struct {
	Sign      string  `toml:"tree_sign"`
	SpawnRate float64 `toml:"tree_spawn_rate"`
}

func MustNew() *Config {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	c := &Config{}

	_, err = toml.DecodeFile(path.Join(wd, "config", "app.toml"), c)
	if err != nil {
		panic(err)
	}

	return c
}
