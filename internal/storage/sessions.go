package storage

import(
	"database/sql"
	"github.com/Adarshrai24/ttyper/internal/test"
)

type Session struct {
	GWPM float64 
	NWPM float64 
	Accuracy float64
	Duration int
}

func SaveSession(db *sql.DB, result test.Result, duration int) error {
	query := `
	INSERT INTO sessions (gwpm, nwpm, accuracy, duration)
	VALUES(?, ?, ?, ?)
	`
	_, err := db.Exec(
		query, result.GWPM, result.NWPM, result.Accuracy, duration,
	)

	if err != nil {
		return err
	}
	return nil
}

func GetHistory(db *sql.DB) ([]Session, error) {
	query := `
		SELECT gwpm, nwpm, accuracy, duration FROM sessions;
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var history []Session

	for rows.Next() {
		var hist Session
		err = rows.Scan(
			&hist.GWPM, &hist.NWPM, &hist.Accuracy, &hist.Duration,
		)
		if err != nil {
			return nil, err
		}
		history = append(history, hist)
	}

	return history, nil
}
