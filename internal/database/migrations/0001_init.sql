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
);

CREATE INDEX IF NOT EXISTS idx_playlist_videos_playlist_id ON playlist_videos(playlist_id);
CREATE INDEX IF NOT EXISTS idx_playlist_videos_orden ON playlist_videos(playlist_id, orden);
CREATE INDEX IF NOT EXISTS idx_playlist_videos_video ON playlist_videos(video);
CREATE INDEX IF NOT EXISTS idx_playlists_nombre ON playlists(nombre);
CREATE INDEX IF NOT EXISTS idx_playlists_creado_en ON playlists(creado_en);
CREATE INDEX IF NOT EXISTS idx_procesos_pid ON procesos_ejecucion(pid);
CREATE INDEX IF NOT EXISTS idx_procesos_estado ON procesos_ejecucion(estado);
CREATE INDEX IF NOT EXISTS idx_procesos_iniciado_en ON procesos_ejecucion(iniciado_en);
CREATE INDEX IF NOT EXISTS idx_procesos_estado_iniciado ON procesos_ejecucion(estado, iniciado_en);
CREATE INDEX IF NOT EXISTS idx_stadicticasPlay_fecha_video ON stadicticasPlay (create_to, video);
CREATE INDEX IF NOT EXISTS idx_stadicticasPlay_date_expr ON stadicticasPlay (DATE(create_to), video);