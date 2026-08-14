package vars

var (
	MIGRATION_OOONE = `
		CREATE TABLE IF NOT EXISTS videos (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			titulo TEXT NOT NULL,
			descripcion TEXT,
			nombre_archivo TEXT NOT NULL UNIQUE,
			creado_en DATETIME DEFAULT CURRENT_TIMESTAMP
		);

		CREATE TABLE IF NOT EXISTS playlists (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			nombre TEXT NOT NULL,
			descripcion TEXT,
			creado_en DATETIME DEFAULT CURRENT_TIMESTAMP
		);

		CREATE TABLE IF NOT EXISTS playlist_videos (
			playlist_id INTEGER,
			video_id INTEGER,
			orden INTEGER NOT NULL DEFAULT 0,
			agregado_en DATETIME DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (playlist_id, video_id),
			FOREIGN KEY (playlist_id) REFERENCES playlists(id) ON DELETE CASCADE,
			FOREIGN KEY (video_id) REFERENCES videos(id) ON DELETE CASCADE
		);

		CREATE TABLE IF NOT EXISTS procesos_ejecucion (
			pid INTEGER PRIMARY KEY, -- El PID del OS es único por naturaleza mientras corre
			ruta_ejecutable TEXT NOT NULL,
			estado TEXT DEFAULT 'corriendo', -- ejemplo: corriendo, detenido, fallido
			iniciado_en DATETIME DEFAULT CURRENT_TIMESTAMP
		);
	`
	INSERT_PID_STATEMENT      = "INSERT OR REPLACE INTO procesos_ejecucion (pid, ruta_ejecutable, estado) VALUES (?, ?, ?);"
	INSERT_PLAYLIST_STATEMENT = `
		INSERT INTO playlists (nombre, descripcion)
		VALUES (?, ?)
	`

	GET_VIDEOS_PAGES_EXECUTE = `
		SELECT id, titulo, descripcion, nombre_archivo, creado_en 
              FROM videos 
              ORDER BY creado_en DESC 
              LIMIT ? OFFSET ?;
	`
	GET_LAST_TEN_PROCESS_EXECUTE = `
		SELECT pid, ruta_ejecutable, estado, iniciado_en 
			FROM procesos_ejecucion 
			ORDER BY iniciado_en DESC 
			LIMIT 10;
	`
)
