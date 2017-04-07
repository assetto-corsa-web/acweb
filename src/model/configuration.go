package model

import (
	"database/sql"
	"db"
	"errors"
)

type Configuration struct {
	Id                 int64     `json:"id"`
	Name               string    `json:"name"`
	Pwd                string    `json:"pwd"`
	AdminPwd           string    `json:"admin_pwd"`
	PickupMode         bool      `json:"pickup_mode"`
	RaceOvertime       int       `json:"race_overtime"`
	MaxSlots           int       `json:"max_slots"`
	Description        string    `json:"description"`
	UDP                int       `json:"udp"`
	TCP                int       `json:"tcp"`
	HTTP               int       `json:"http"`
	PacketsHz          int       `json:"packets_hz"`
	LoopMode           bool      `json:"loop_mode"`
	ShowInLobby        bool      `json:"show_in_lobby"`
	ABS                string    `json:"abs"`
	TC                 string    `json:"tc"`
	StabilityAid       bool      `json:"stability_aid"`
	AutoClutch         bool      `json:"auto_clutch"`
	TyreBlankets       bool      `json:"tyre_blankets"`
	ForceVirtualMirror bool      `json:"force_virtual_mirror"`
	FuelRate           int       `json:"fuel_rate"`
	DamageRate         int       `json:"damage_rate"`
	TiresWearRate      int       `json:"tires_wear_rate"`
	AllowedTiresOut    int       `json:"allowed_tires_out"`
	MaxBallast         int       `json:"max_ballast"`
	DynamicTrack       bool      `json:"dynamic_track"`
	Condition          string    `json:"condition"`
	StartValue         int       `json:"start_value"`
	Randomness         int       `json:"randomness"`
	TransferredGrip    int       `json:"transferred_grip"`
	LapsToImproveGrip  int       `json:"laps_to_improve_grip"`
	KickVoteQuorum     int       `json:"kick_vote_quorum"`
	SessionVoteQuorum  int       `json:"session_vote_quorum"`
	VoteDuration       int       `json:"vote_duration"`
	Blacklist          string    `json:"blacklist"`
	Booking            bool      `json:"booking"`
	BookingTime        int       `json:"booking_time"`
	Practice           bool      `json:"practice"`
	PracticeTime       int       `json:"practice_time"`
	CanJoinPractice    bool      `json:"can_join_practice"`
	Qualify            bool      `json:"qualify"`
	QualifyTime        int       `json:"qualify_time"`
	CanJoinQualify     bool      `json:"can_join_qualify"`
	Race               bool      `json:"race"`
	RaceTime           int       `json:"race_time"`
	RaceWaitTime       int       `json:"race_wait_time"`
	JoinType           string    `json:"join_type"`
	Time               string    `json:"time"`
	Weather            []Weather `json:"weather"`
	Track              string    `json:"track"`
	Cars               []Car     `json:"cars"`
}

type Weather struct {
	Id                int64  `json:"id"`
	Configuration     int64  `json:"configuration"`
	Weather           string `json:"weather"`
	BaseAmbientTemp   int    `json:"base_ambient_temp"`
	RealisticRoadTemp int    `json:"realistic_road_temp"`
	BaseRoadTemp      int    `json:"base_road_temp"`
	AmbientVariation  int    `json:"ambient_variation"`
	RoadVariation     int    `json:"road_variation"`
}

type Car struct {
	Id            int64  `json:"id"`
	Configuration int64  `json:"configuration"`
	Car           string `json:"car"`
	Painting      string `json:"painting"`
	Position      int    `json:"position"`
}

func (m *Configuration) Save() error {
	tx, err := db.Get().Begin()

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

func (m *Configuration) saveConfiguration(tx *sql.Tx) error {
	if m.Id == 0 {
		res, err := tx.Exec(`INSERT INTO configurations (name,
			pwd,
			admin_pwd,
			pickup_mode,
			race_overtime,
			max_slots,
			description,
			udp,
			tcp,
			http,
			packets_hz,
			loop_mode,
			show_in_lobby,
			abs,
			tc,
			stability_aid,
			auto_clutch,
			tyre_blankets,
			force_virtual_mirror,
			fuel_rate,
			damage_rate,
			tires_wear_rate,
			allowed_tires_out,
			max_ballast,
			dynamic_track,
			track_condition,
			start_value,
			randomness,
			transferred_grip,
			laps_to_improve_grip,
			kick_vote_quorum,
			session_vote_quorum,
			vote_duration,
			blacklist,
			booking,
			booking_time,
			practice,
			practice_time,
			can_join_practice,
			qualify,
			qualify_time,
			can_join_qualify,
			race,
			race_time,
			race_wait_time,
			join_type,
			time,
			track) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
			?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
			?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
			?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
			?, ?, ?, ?, ?, ?, ?, ?)`, m.Name,
			m.Pwd,
			m.AdminPwd,
			m.PickupMode,
			m.RaceOvertime,
			m.MaxSlots,
			m.Description,
			m.UDP,
			m.TCP,
			m.HTTP,
			m.PacketsHz,
			m.LoopMode,
			m.ShowInLobby,
			m.ABS,
			m.TC,
			m.StabilityAid,
			m.AutoClutch,
			m.TyreBlankets,
			m.ForceVirtualMirror,
			m.FuelRate,
			m.DamageRate,
			m.TiresWearRate,
			m.AllowedTiresOut,
			m.MaxBallast,
			m.Condition,
			m.StartValue,
			m.Randomness,
			m.TransferredGrip,
			m.LapsToImproveGrip,
			m.KickVoteQuorum,
			m.SessionVoteQuorum,
			m.VoteDuration,
			m.Blacklist,
			m.Booking,
			m.BookingTime,
			m.Practice,
			m.PracticeTime,
			m.CanJoinPractice,
			m.Qualify,
			m.QualifyTime,
			m.CanJoinQualify,
			m.Race,
			m.RaceTime,
			m.RaceWaitTime,
			m.JoinType,
			m.Time,
			m.Track)

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
		return nil
	}

	_, err := tx.Exec(`UPDATE configurations SET name = ?,
		pwd = ?,
		admin_pwd = ?,
		pickup_mode = ?,
		race_overtime = ?,
		max_slots = ?,
		description = ?,
		udp = ?,
		tcp = ?,
		http = ?,
		packets_hz = ?,
		loop_mode = ?,
		show_in_lobby = ?,
		abs = ?,
		tc = ?,
		stability_aid = ?,
		auto_clutch = ?,
		tyre_blankets = ?,
		force_virtual_mirror = ?,
		fuel_rate = ?,
		damage_rate = ?,
		tires_wear_rate = ?,
		allowed_tires_out = ?,
		max_ballast = ?,
		dynamic_track = ?,
		track_condition = ?,
		start_value = ?,
		randomness = ?,
		transferred_grip = ?,
		laps_to_improve_grip = ?,
		kick_vote_quorum = ?,
		session_vote_quorum = ?,
		vote_duration = ?,
		blacklist = ?,
		booking = ?,
		booking_time = ?,
		practice = ?,
		practice_time = ?,
		can_join_practice = ?,
		qualify = ?,
		qualify_time = ?,
		can_join_qualify = ?,
		race = ?,
		race_time = ?,
		race_wait_time = ?,
		join_type = ?,
		time = ?,
		track = ? WHERE id = ?`, m.Name,
		m.Pwd,
		m.AdminPwd,
		m.PickupMode,
		m.RaceOvertime,
		m.MaxSlots,
		m.Description,
		m.UDP,
		m.TCP,
		m.HTTP,
		m.PacketsHz,
		m.LoopMode,
		m.ShowInLobby,
		m.ABS,
		m.TC,
		m.StabilityAid,
		m.AutoClutch,
		m.TyreBlankets,
		m.ForceVirtualMirror,
		m.FuelRate,
		m.DamageRate,
		m.TiresWearRate,
		m.AllowedTiresOut,
		m.MaxBallast,
		m.Condition,
		m.StartValue,
		m.Randomness,
		m.TransferredGrip,
		m.LapsToImproveGrip,
		m.KickVoteQuorum,
		m.SessionVoteQuorum,
		m.VoteDuration,
		m.Blacklist,
		m.Booking,
		m.BookingTime,
		m.Practice,
		m.PracticeTime,
		m.CanJoinPractice,
		m.Qualify,
		m.QualifyTime,
		m.CanJoinQualify,
		m.Race,
		m.RaceTime,
		m.RaceWaitTime,
		m.JoinType,
		m.Time,
		m.Track,
		m.Id)

	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (m *Configuration) saveWeather(tx *sql.Tx) error {
	for _, weather := range m.Weather {
		if weather.Id == 0 {
			_, err := tx.Exec(`INSERT INTO weather (configuration,
				weather,
				base_ambient_temp,
				realistic_road_temp,
				base_road_temp,
				ambient_variation,
				road_variation) VALUES (?, ?, ?, ?, ?, ?, ?)`, m.Id,
				weather.Weather,
				weather.BaseAmbientTemp,
				weather.RealisticRoadTemp,
				weather.BaseRoadTemp,
				weather.AmbientVariation,
				weather.RoadVariation)

			if err != nil {
				tx.Rollback()
				return err
			}
		} else {
			_, err := tx.Exec(`UPDATE weather SET weather = ?,
				base_ambient_temp = ?,
				realistic_road_temp = ?,
				base_road_temp = ?,
				ambient_variation = ?,
				road_variation = ? WHERE id = ?`, weather.Weather,
				weather.BaseAmbientTemp,
				weather.RealisticRoadTemp,
				weather.BaseRoadTemp,
				weather.AmbientVariation,
				weather.RoadVariation,
				m.Id)

			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	return nil
}

func (m *Configuration) saveCars(tx *sql.Tx) error {
	for _, car := range m.Cars {
		if car.Id == 0 {
			_, err := tx.Exec(`INSERT INTO cars (configuration,
				car,
				painting,
				position) VALUES (?, ?, ?, ?)`, m.Id,
				car.Car,
				car.Painting,
				car.Position)

			if err != nil {
				tx.Rollback()
				return err
			}
		} else {
			_, err := tx.Exec(`UPDATE cars SET car = ?,
				painting = ?,
				position = ? WHERE id = ?`, car.Car,
				car.Painting,
				car.Position,
				m.Id)

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

	tx, err := db.Get().Begin()

	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM weather WHERE configuration = ?", m.Id)

	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("DELETE FROM cars WHERE configuration = ?", m.Id)

	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("DELETE FROM configuration WHERE id = ?", m.Id)

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
