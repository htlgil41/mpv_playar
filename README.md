# 🎥 mpv_playar

Aplicación y API REST escrita en **Go** para el control y gestión de playlists de video en pantallas publicitarias para la corporación/grupo **MaraPlus**.

La aplicación está desarrollada principalmente en Go utilizando el framework web **Gin**. Consta de varios endpoints que manejan directamente la escritura en el archivo socket o pipe del sistema hacia el servidor MPV.

---

### 💡 ¿Cómo funciona la comunicación con MPV?

> 🤖 **El socket (o IPC socket) en MPV es un canal de comunicación que permite controlar el reproductor en tiempo real desde otros programas. Envías comandos en texto (como pausar o cambiar volumen) y el reproductor responde usando formato JSON.**

---

## ⚙️ Inicialización y Configuración

Antes de comenzar, es necesario configurar el entorno para que el servidor HTTP se ejecute correctamente.

En la raíz del binario de la API se debe crear el archivo de configuración `config.yml`, donde se establecen las rutas para el pipe/socket, la ubicación de los videos y la base de datos SQLite.

### 📝 Ejemplo de `config.yml`:

```yaml
paths:
  path_mega: '/home/htgil41/Videos/'
  path_server_mpv: '/tmp/mpvsocket'

server:
  port: 8000
  dblite: './playar.db'
```

### 📌 Descripción de parámetros:

* **`paths.path_mega`**: Directorio donde se encuentran los videos que la API recupera para listar, agregar y quitar de las playlists creadas.
* **`paths.path_server_mpv`**: Ruta donde MPV levanta su socket IPC.
  * **Linux**: `/tmp/mpvsocket`
  * **Windows**: `\\.\pipe\mpvsocket`
* **`server.port`**: Puerto en el que se ejecutará el servidor HTTP.
* **`server.dblite`**: Ruta del archivo de la base de datos SQLite.

> ℹ️ **Nota:** Las rutas (paths) deben colocarse exactamente como las proporciona el sistema operativo (el ejemplo mostrado arriba corresponde a Linux).

---

## 🚀 Documentación de la API (Endpoints)

**Base URL:** `http://localhost:8000` *(Ejemplo)*

### 📊 Sistema y Métricas

| Método | Endpoint     | Descripción                                                       |
| :-----: | :----------- | :----------------------------------------------------------------- |
| `GET` | `/ping`    | Verifica el estado del servidor y la conexión a la base de datos. |
| `GET` | `/pid`     | Obtiene los últimos identificadores de proceso (PIDs) activos.    |
| `GET` | `/metrica` | Devuelve las métricas y estadísticas de los videos reproducidos. |

### 🎬 Gestión de Videos y Rutas

| Método | Endpoint         | Descripción                                                                          |
| :-----: | :--------------- | :------------------------------------------------------------------------------------ |
| `GET` | `/videos-mega` | Lista los archivos de video disponibles en la ruta local configurada (`path_mega`). |

### 🎵 Control de Playlists

|  Método  | Endpoint             | Descripción                                                               |
| :--------: | :------------------- | :------------------------------------------------------------------------- |
|  `GET`  | `/playlist`        | Obtiene la lista de reproducción actual.                                  |
|  `POST`  | `/playlist`        | Crea una nueva lista de reproducción en la base de datos local.           |
|  `POST`  | `/add-vtoplaylist` | Añade un video específico a una lista de reproducción existente.        |
| `DELETE` | `/stop-playlis`    | Vacía por completo la lista de reproducción actual y detiene el flujo.   |
| `DELETE` | `/playlist`        | Elimina una Playlist de la bases de datos local y sus videos relacionados. |

### 🕹️ Reproducción en Servidor (MPV)

| Método | Endpoint              | Descripción                                                                         |
| :------: | :-------------------- | :----------------------------------------------------------------------------------- |
| `POST` | `/play-vtoplaylist` | Envía un video de la playlist directamente al socket Unix de MPV para reproducirlo. |
| `POST` | `/next-video`       | Salta al siguiente video de la cola a través del servidor Unix.                     |
| `POST` | `/playlist-newplay` | Inicializa y reproduce una nueva lista de reproducción cargando la configuración.  |

## 🔨 Build proyecto for Linux/Windows

```
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./app/playar.exe ./cmd/server/
```

```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app/playar ./cmd/server/
```

**CGO_ENABLED** Puro codigo Go - **GOOS** Sistema operativo linux/windows - **GOARCH** Arquitextura del SO
