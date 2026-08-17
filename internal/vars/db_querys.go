package vars

var (
	MIGRATION_OOONE = `
		CREATE TABLE IF NOT EXISTS playlists (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			nombre TEXT NOT NULL,
			descripcion TEXT,
			creado_en DATETIME DEFAULT CURRENT_TIMESTAMP
		);

		CREATE TABLE IF NOT EXISTS playlist_videos (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			playlist_id INTEGER,
			video TEXT,
			orden INTEGER NOT NULL DEFAULT 0,
			agregado_en DATETIME DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (playlist_id) REFERENCES playlists(id) ON DELETE CASCADE
		);

		CREATE TABLE IF NOT EXISTS procesos_ejecucion (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			pid INTEGER NOT NULL,
			ruta_ejecutable TEXT NOT NULL,
			estado TEXT DEFAULT 'corriendo',
			iniciado_en DATETIME DEFAULT CURRENT_TIMESTAMP
		);

		CREATE TABLE IF NOT EXISTS stadicticasPlay (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			video TEXT NOT NULL,
			create_to DATETIME DEFAULT CURRENT_TIMESTAMP
		)

		CREATE INDEX IF NOT EXISTS idx_playlist_videos_playlist_id ON playlist_videos(playlist_id);
		CREATE INDEX IF NOT EXISTS idx_playlist_videos_orden ON playlist_videos(playlist_id, orden);
		CREATE INDEX IF NOT EXISTS idx_playlist_videos_video ON playlist_videos(video);
		CREATE INDEX IF NOT EXISTS idx_playlists_nombre ON playlists(nombre);
		CREATE INDEX IF NOT EXISTS idx_playlists_creado_en ON playlists(creado_en);
		CREATE INDEX IF NOT EXISTS idx_procesos_pid ON procesos_ejecucion(pid);
		CREATE INDEX IF NOT EXISTS idx_procesos_estado ON procesos_ejecucion(estado);
		CREATE INDEX IF NOT EXISTS idx_procesos_iniciado_en ON procesos_ejecucion(iniciado_en);
		CREATE INDEX IF NOT EXISTS idx_procesos_estado_iniciado ON procesos_ejecucion(estado, iniciado_en);
		CREATE INDEX idx_stadicticasPlay_fecha_video ON stadicticasPlay (create_to, video);
		CREATE INDEX idx_stadicticasPlay_date_expr ON stadicticasPlay (DATE(create_to), video);
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
	GET_LIST_PLAYIST      = `SELECT id, nombre, descripcion FROM playlists order by creado_en DESC`
	GET_VIDEO_BY_PLAYLIST = `SELECT video FROM playlist_videos WHERE playlist_id = ? ORDER BY orden`

	INSERT_PID_STATEMENT      = "INSERT OR REPLACE INTO procesos_ejecucion (pid, ruta_ejecutable, estado) VALUES (?, ?, ?);"
	INSERT_PLAYLIST_STATEMENT = `
		INSERT INTO playlists (nombre, descripcion)
		VALUES (?, ?)
	`
	INSERT_VIDEOS_ON_PLAYLIST = `
		INSERT INTO playlist_videos (playlist_id, video, orden)
		VALUES (?, ?, ?)
	`
)
