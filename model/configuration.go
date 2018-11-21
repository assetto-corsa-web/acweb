package model

import (
	"errors"

	"github.com/jmoiron/sqlx"
)

type Configuration struct {
	Id                    int64  `json:"id"`
	Name                  string `json:"name"`
	Pwd                   string `json:"pwd"`
	AdminPwd              string `db:"admin_pwd" json:"admin_pwd"`
	PickupMode            bool   `db:"pickup_mode" json:"pickup_mode"`
	LockEntryList         bool   `db:"lock_entry_list" json:"lock_entry_list"`
	RaceOvertime          int    `db:"race_overtime" json:"race_overtime"`
	MaxSlots              int    `db:"max_slots" json:"max_slots"`
	Welcome               string `json:"welcome"`
	Description           string `json:"description"`
	UDP                   int    `json:"udp"`
	TCP                   int    `json:"tcp"`
	HTTP                  int    `json:"http"`
	PacketsHz             int    `db:"packets_hz" json:"packets_hz"`
	LoopMode              bool   `db:"loop_mode" json:"loop_mode"`
	ShowInLobby           bool   `db:"show_in_lobby" json:"show_in_lobby"`
	Threads               int    `json:"threads"`
	AuthPluginAddress     int    `db:"auth_plugin_address" json:"auth_plugin_address"`
	ABS                   int    `json:"abs"`
	TC                    int    `json:"tc"`
	StabilityAid          bool   `db:"stability_aid" json:"stability_aid"`
	AutoClutch            bool   `db:"auto_clutch" json:"auto_clutch"`
	TyreBlankets          bool   `db:"tyre_blankets" json:"tyre_blankets"`
	ForceVirtualMirror    bool   `db:"force_virtual_mirror" json:"force_virtual_mirror"`
	FuelRate              int    `db:"fuel_rate" json:"fuel_rate"`
	DamageRate            int    `db:"damage_rate" json:"damage_rate"`
	TiresWearRate         int    `db:"tires_wear_rate" json:"tires_wear_rate"`
	AllowedTiresOut       int    `db:"allowed_tires_out" json:"allowed_tires_out"`
	MaxBallast            int    `db:"max_ballast" json:"max_ballast"`
	StartRule             int    `db:"start_rule" json:"start_rule"`
	DisableGasCutPenality bool   `db:"disable_gas_cut_penality" json:"disable_gas_cut_penality"`
	TimeOfDayMult         int    `db:"time_of_day_mult" json:"time_of_day_mult"`
	ResultScreenTime      int    `db:"result_screen_time" json:"result_screen_time"`
	DynamicTrack          bool   `db:"dynamic_track" json:"dynamic_track"`
	Condition             string `db:"track_condition" json:"condition"`
	StartValue            int    `db:"start_value" json:"start_value"`
	Randomness            int    `json:"randomness"`
	TransferredGrip       int    `db:"transferred_grip" json:"transferred_grip"`
	LapsToImproveGrip     int    `db:"laps_to_improve_grip" json:"laps_to_improve_grip"`
	KickVoteQuorum        int    `db:"kick_vote_quorum" json:"kick_vote_quorum"`
	SessionVoteQuorum     int    `db:"session_vote_quorum" json:"session_vote_quorum"`
	VoteDuration          int    `db:"vote_duration" json:"vote_duration"`
	Blacklist             int    `json:"blacklist"`
	MaxCollisionsKm       int    `db:"max_collisions_km" json:"max_collisions_km"`
	Booking               bool   `json:"booking"`
	BookingTime           int    `db:"booking_time" json:"booking_time"`
	Practice              bool   `json:"practice"`
	PracticeTime          int    `db:"practice_time" json:"practice_time"`
	CanJoinPractice       bool   `db:"can_join_practice" json:"can_join_practice"`
	Qualify               bool   `json:"qualify"`
	QualifyTime           int    `db:"qualify_time" json:"qualify_time"`
	CanJoinQualify        bool   `db:"can_join_qualify" json:"can_join_qualify"`
	Race                  bool   `json:"race"`
	RaceLaps              int    `db:"race_laps" json:"race_laps"`
	RaceTime              int    `db:"race_time" json:"race_time"`
	RaceWaitTime          int    `db:"race_wait_time" json:"race_wait_time"`
	RaceExtraLap          bool   `db:"race_extra_lap" json:"race_extra_lap"`
	JoinType              int    `db:"join_type" json:"join_type"`
	SunAngle              int    `db:"sun_angle" json:"sun_angle"`
	Track                 string `json:"track"`
	TrackConfig           string `db:"track_config" json:"track_config"`
	LegalTyres            string `db:"legal_tyres" json:"legal_tyres"`
	UdpPluginPort         int    `db:"udp_plugin_local_port" json:"udp_plugin_local_port"`
	UdpPluginAddr         string `db:"udp_plugin_address" json:"udp_plugin_address"`
	RacePitWindowStart    int    `db:"race_pit_window_start" json:"race_pit_window_start"`
	RacePitWindowEnd      int    `db:"race_pit_window_end" json:"race_pit_window_end"`
	ReversedGridRacePos   int    `db:"reversed_grid_race_positions" json:"reversed_grid_race_positions"`
	ServerCfgIni          string `db:"server_cfg_ini" json:"server_cfg_ini"`
	EntryListIni          string `db:"entry_list_ini" json:"entry_list_ini"`

	Weather []Weather `db:"-" json:"weather"`
	Cars    []Car     `db:"-" json:"cars"`
}

type Weather struct {
	Id                     int64  `json:"id"`
	Configuration          int64  `json:"configuration"`
	Weather                string `json:"weather"`
	BaseAmbientTemp        int    `db:"base_ambient_temp" json:"base_ambient_temp"`
	BaseRoadTemp           int    `db:"base_road_temp" json:"base_road_temp"`
	AmbientVariation       int    `db:"ambient_variation" json:"ambient_variation"`
	RoadVariation          int    `db:"road_variation" json:"road_variation"`
	WindBaseSpeedMin       int    `db:"wind_base_speed_min" json:"wind_base_speed_min"`
	WindBaseSpeedMax       int    `db:"wind_base_speed_max" json:"wind_base_speed_max"`
	WindBaseDirection      int    `db:"wind_base_direction" json:"wind_base_direction"`
	WindVariationDirection int    `db:"wind_variation_direction" json:"wind_variation_direction"`
}

type Car struct {
	Id            int64  `json:"id"`
	Configuration int64  `json:"configuration"`
	Car           string `json:"car"`
	Painting      string `json:"painting"`
	Spectator     bool   `json:"spectator"`
	Driver        string `json:"driver"`
	Team          string `json:"team"`
	GUID          string `json:"guid"`
	Position      int    `json:"position"`
	FixedSetup    string `db:"fixed_setup" json:"fixed_setup"`
}

// Joins weather and cars.
func (m *Configuration) Join() error {
	weather, err := GetWeatherByConfiguration(m.Id)

	if err != nil {
		return err
	}

	cars, err := GetCarsByConfiguration(m.Id)

	if err != nil {
		return err
	}

	m.Weather = weather
	m.Cars = cars

	return nil
}

func (m *Configuration) Save() error {
	tx, err := session.Beginx()

	if err != nil {
		return err
	}

	// get an ID first
	if err := m.saveConfiguration(tx); err != nil {
		return err
	}

	// save weather and cars
	if err := m.saveWeather(tx); err != nil {
		return err
	}

	if err := m.saveCars(tx); err != nil {
		return err
	}

	return tx.Commit()
}

func (m *Configuration) saveConfiguration(tx *sqlx.Tx) error {
	if m.Id == 0 {
		if GetDBType() == "mysql" {
			res, err := tx.NamedExec(mysql_configuration_save, m)

			if err != nil {
				tx.Rollback()
				return err
			}

			id, err := res.LastInsertId()

			if err != nil {
				tx.Rollback()
				return err
			}

			m.Id = id
		} else {
			rows, err := tx.NamedQuery(postgres_configuration_save, m)

			if err != nil {
				tx.Rollback()
				return err
			}

			if rows.Next() {
				rows.Scan(&m.Id)
			}

			rows.Close()
		}

		return nil
	}

	var err error

	if GetDBType() == "mysql" {
		_, err = tx.NamedExec(mysql_configuration_update, m)
	} else {
		_, err = tx.NamedExec(postgres_configuration_update, m)
	}

	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (m *Configuration) saveWeather(tx *sqlx.Tx) error {
	for _, weather := range m.Weather {
		if weather.Id == 0 {
			var err error

			if GetDBType() == "mysql" {
				_, err = tx.Exec(mysql_weather_save,
					m.Id,
					weather.Weather,
					weather.BaseAmbientTemp,
					weather.BaseRoadTemp,
					weather.AmbientVariation,
					weather.RoadVariation,
					weather.WindBaseSpeedMin,
					weather.WindBaseSpeedMax,
					weather.WindBaseDirection,
					weather.WindVariationDirection)
			} else {
				_, err = tx.Exec(postgres_weather_save,
					m.Id,
					weather.Weather,
					weather.BaseAmbientTemp,
					weather.BaseRoadTemp,
					weather.AmbientVariation,
					weather.RoadVariation,
					weather.WindBaseSpeedMin,
					weather.WindBaseSpeedMax,
					weather.WindBaseDirection,
					weather.WindVariationDirection)
			}

			if err != nil {
				tx.Rollback()
				return err
			}
		} else {
			var err error

			if GetDBType() == "mysql" {
				_, err = tx.Exec(mysql_weather_update,
					weather.Weather,
					weather.BaseAmbientTemp,
					weather.BaseRoadTemp,
					weather.AmbientVariation,
					weather.RoadVariation,
					weather.WindBaseSpeedMin,
					weather.WindBaseSpeedMax,
					weather.WindBaseDirection,
					weather.WindVariationDirection,
					weather.Id)
			} else {
				_, err = tx.Exec(postgres_weather_update,
					weather.Weather,
					weather.BaseAmbientTemp,
					weather.BaseRoadTemp,
					weather.AmbientVariation,
					weather.RoadVariation,
					weather.WindBaseSpeedMin,
					weather.WindBaseSpeedMax,
					weather.WindBaseDirection,
					weather.WindVariationDirection,
					weather.Id)
			}

			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	return nil
}

func (m *Configuration) saveCars(tx *sqlx.Tx) error {
	for _, car := range m.Cars {
		if car.Id == 0 {
			var err error

			if GetDBType() == "mysql" {
				_, err = tx.Exec(mysql_cars_save,
					m.Id,
					car.Car,
					car.Painting,
					car.Spectator,
					car.Driver,
					car.Team,
					car.GUID,
					car.Position,
					car.FixedSetup)
			} else {
				_, err = tx.Exec(postgres_cars_save,
					m.Id,
					car.Car,
					car.Painting,
					car.Spectator,
					car.Driver,
					car.Team,
					car.GUID,
					car.Position,
					car.FixedSetup)
			}

			if err != nil {
				tx.Rollback()
				return err
			}
		} else {
			var err error

			if GetDBType() == "mysql" {
				_, err = tx.Exec(mysql_cars_update,
					car.Car,
					car.Painting,
					car.Spectator,
					car.Driver,
					car.Team,
					car.GUID,
					car.Position,
					car.FixedSetup,
					car.Id)
			} else {
				_, err = tx.Exec(postgres_cars_update,
					car.Car,
					car.Painting,
					car.Spectator,
					car.Driver,
					car.Team,
					car.GUID,
					car.Position,
					car.FixedSetup,
					car.Id)
			}

			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	return nil
}

func (m *Configuration) Remove() error {
	if m.Id == 0 {
		return errors.New("ID must be set")
	}

	tx, err := session.Beginx()

	if err != nil {
		return err
	}

	if GetDBType() == "mysql" {
		_, err = tx.NamedExec(mysql_weather_delete_configuration, m)

		if err != nil {
			tx.Rollback()
			return err
		}

		_, err = tx.NamedExec(mysql_cars_delete_configuration, m)

		if err != nil {
			tx.Rollback()
			return err
		}

		_, err = tx.NamedExec(mysql_configuration_delete, m)

		if err != nil {
			tx.Rollback()
			return err
		}
	} else {
		_, err = tx.NamedExec(postgres_weather_delete_configuration, m)

		if err != nil {
			tx.Rollback()
			return err
		}

		_, err = tx.NamedExec(postgres_cars_delete_configuration, m)

		if err != nil {
			tx.Rollback()
			return err
		}

		_, err = tx.NamedExec(postgres_configuration_delete, m)

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (m *Weather) Remove() error {
	var err error

	if GetDBType() == "mysql" {
		_, err = session.NamedExec(mysql_weather_delete, m)
	} else {
		_, err = session.NamedExec(postgres_weather_delete, m)
	}

	return err
}

func (m *Car) Remove() error {
	var err error

	if GetDBType() == "mysql" {
		_, err = session.NamedExec(mysql_cars_delete, m)
	} else {
		_, err = session.NamedExec(postgres_cars_delete, m)
	}

	return err
}

func GetWeatherByConfiguration(id int64) ([]Weather, error) {
	weather := make([]Weather, 0)

	if GetDBType() == "mysql" {
		if err := session.Select(&weather, mysql_weather_get_configuration, id); err != nil {
			return nil, err
		}
	} else {
		if err := session.Select(&weather, postgres_weather_get_configuration, id); err != nil {
			return nil, err
		}
	}

	return weather, nil
}

func GetCarsByConfiguration(id int64) ([]Car, error) {
	cars := make([]Car, 0)

	if GetDBType() == "mysql" {
		if err := session.Select(&cars, mysql_cars_get_configuration, id); err != nil {
			return nil, err
		}
	} else {
		if err := session.Select(&cars, postgres_cars_get_configuration, id); err != nil {
			return nil, err
		}
	}

	return cars, nil
}

func GetAllConfigurations() ([]Configuration, error) {
	configs := make([]Configuration, 0)

	if GetDBType() == "mysql" {
		if err := session.Select(&configs, mysql_configuration_get_all); err != nil {
			return nil, err
		}
	} else {
		if err := session.Select(&configs, postgres_configuration_get_all); err != nil {
			return nil, err
		}
	}

	return configs, nil
}

func GetConfigurationById(id int64) (*Configuration, error) {
	config := new(Configuration)

	if GetDBType() == "mysql" {
		if err := session.Get(config, mysql_configuration_get_id, id); err != nil {
			return nil, err
		}
	} else {
		if err := session.Get(config, postgres_configuration_get_id, id); err != nil {
			return nil, err
		}
	}

	return config, nil
}
