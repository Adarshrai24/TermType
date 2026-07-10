package storage

import(
	"database/sql"
)

type Stats struct {
	GWPM float64
	NWPM float64
	Accuracy float64
	Duration int
}

func MaximumStats(db *sql.DB, duration int) (Stats, error) {
	query := `
	SELECT MAX(gwpm), MAX(nwpm), MAX(accuracy) 
	FROM sessions
	WHERE duration = ?
	`
	var stats Stats
	stats.Duration = duration
	err := db.QueryRow(query, duration).Scan(
		&stats.GWPM, &stats.NWPM, &stats.Accuracy,
	)

	if err != nil {
		return stats, err
	}
	
	return stats, nil	
}

func AverageStats(db *sql.DB, duration int) (Stats, error) {
	query := `
	SELECT AVG(gwpm), AVG(nwpm), AVG(accuracy) 
	FROM sessions
	WHERE duration = ?
	`
	var stats Stats
	stats.Duration = duration
	err := db.QueryRow(query, duration).Scan(
		&stats.GWPM, &stats.NWPM, &stats.Accuracy,
	)
	

	if err != nil {
		return stats, err
	}
	
	return stats, nil
}
