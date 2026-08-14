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
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    pid INTEGER NOT NULL,
    ruta_ejecutable TEXT NOT NULL,
    estado TEXT DEFAULT 'corriendo',
    iniciado_en DATETIME DEFAULT CURRENT_TIMESTAMP
);
