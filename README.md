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

| Método | Endpoint          | Descripción                                                       |
| :-----: | :--------------- | :----------------------------------------------------------------- |
| `GET` | `/ping`         | Verifica el estado del servidor y la conexión a la base de datos. |
| `GET` | `/pid`          | Obtiene los últimos identificadores de proceso (PIDs) activos.    |
| `GET` | `/metrica`      | Devuelve las métricas y estadísticas de los videos reproducidos.  |
| `GET` | `/metricas-range` | Devuelve métricas de videos reproduccidos en un rango de fechas específico. |

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
| `DELETE` | `/stop-playlist`   | Detiene la reproducción actual en MPV.                                   |
| `DELETE` | `/playlist`        | Elimina una Playlist de la bases de datos local y sus videos relacionados. |

### 🕹️ Reproducción en Servidor (MPV)

| Método | Endpoint              | Descripción                                                                         |
| :------: | :-------------------- | :----------------------------------------------------------------------------------- |
| `POST` | `/play-vtoplaylist` | Envía un video de la playlist directamente al socket Unix de MPV para reproducirlo. |
| `POST` | `/next-video`       | Salta al siguiente video de la cola a través del servidor Unix.                     |
| `POST` | `/playlist-newplay` | Inicializa y reproduce una nueva lista de reproducción cargando la configuración.  |

### 🔌 WebSocket

| Método | Endpoint | Descripción |
| :------: | :------- | :---------- |
| `GET` (WS) | `/ws` | Endpoint WebSocket para notificaciones en tiempo real. |

---

## 📖 Detalle de Endpoints

### 📈 `GET /metricas-range`

Obtiene métricas de reproducción de videos filtradas por un rango de fechas exacto.

**Parámetros de consulta (query params):**

| Parámetro | Tipo   | Formato      | Requerido | Descripción     |
| :-------- | :----- | :----------- | :--------: | :-------------- |
| `gte`     | `date` | `YYYY-MM-DD` |     Sí     | Fecha de inicio |
| `lte`     | `date` | `YYYY-MM-DD` |     Sí     | Fecha de fin    |

**Ejemplo:**

```
GET /metricas-range?gte=2026-01-01&lte=2026-08-20
```

**Respuesta exitosa (200):**

```json
{
  "querys": {
    "Gte": "2026-01-01T00:00:00Z",
    "Lte": "2026-08-20T00:00:00Z"
  },
  "data": [
    {
      "Fecha": "0001-01-01T00:00:00Z",
      "Video": "video_nombre.mp4",
      "Repoducc": 42
    }
  ]
}
```

**Respuesta error (400):**

```json
{
  "error": "Error al tratar de serealizar los datos",
  "message": "Field Gte is required"
}
```

**Respuesta error (500):**

```json
{
  "error": "error de la base de datos",
  "message": "No se ha podido obtener la informacion de la db"
}
```

---

### 🎵 `POST /playlist` — Crear Playlist

Crea una nueva playlist en la base de datos.

**Body (JSON):**

| Campo        | Tipo     | Requerido | Descripción      |
| :----------- | :------- | :-------: | :--------------- |
| `nombre`     | `string` |    Sí     | Nombre de la playlist |
| `descripcion`| `string` |    Sí     | Descripción de la playlist |

**Ejemplo:**

```json
{
  "nombre": "Mi Playlist",
  "descripcion": "Videos promocionales"
}
```

**Respuesta exitosa (201):**

```json
{
  "error": null,
  "message": "Plyalist creada correctamente"
}
```

**Respuesta error (500):**

```json
{
  "error": "Error al tratar de crear una playlist",
  "details": "UNIQUE constraint failed"
}
```

---

### ➕ `POST /add-vtoplaylist` — Agregar Video a Playlist

Agrega un video a una playlist existente en la base de datos.

**Body (JSON):**

| Campo      | Tipo     | Requerido | Descripción            |
| :--------- | :------- | :-------: | :--------------------- |
| `playlist` | `number` |    Sí     | ID de la playlist      |
| `video`    | `string` |    Sí     | Nombre del archivo de video |
| `orden`    | `number` |    Sí     | Posición del video en la playlist |

**Ejemplo:**

```json
{
  "playlist": 1,
  "video": "video_promo.mp4",
  "orden": 1
}
```

**Respuesta exitosa (200):**

```json
{
  "error": null,
  "message": "Musica agregada correctamente",
  "data": {
    "playlist": 1,
    "video": "video_promo.mp4",
    "orden": 1
  }
}
```

**Respuesta error (500):**

```json
{
  "error": "No se ha podido completar la consula",
  "details": "FOREIGN KEY constraint failed"
}
```

---

### ▶️ `POST /play-vtoplaylist` — Reproducir Video

Envía un video directamente al socket de MPV para reproducirlo.

**Body (JSON):**

| Campo    | Tipo     | Requerido | Descripción            |
| :------- | :------- | :-------: | :--------------------- |
| `titulo` | `string` |    Sí     | Nombre del archivo de video |

**Ejemplo:**

```json
{
  "titulo": "video_promo.mp4"
}
```

**Respuesta exitosa (200):**

```json
{
  "message": "Se ha podido ejecutar correctamente el comando.",
  "error": null
}
```

**Respuesta error (404):**

```json
{
  "error": "No se ha podido encontrar el archivo dento del directorio configurado",
  "path": "/home/htgil41/Videos/video_promo.mp4"
}
```

**Respuesta error (500):**

```json
{
  "error": "write error",
  "message": "No se ha podido agregar el video"
}
```

---

### ⏭️ `POST /next-video` — Siguiente Video

Salta al siguiente video en la cola de reproducción de MPV.

**Body:** No requiere body.

**Respuesta exitosa (200):**

```json
{
  "message": "Se ha podido ejecutar correctamente el comando"
}
```

**Respuesta error (500):**

```json
{
  "error": "write error",
  "message": "No se ha podido pasar el video"
}
```

---

### 🎬 `POST /playlist-newplay` — Nueva Playlist y Reproducir

Crea una playlist en MPV con todos los videos de una playlist de la DB y comienza la reproducción.

**Body (JSON):**

| Campo      | Tipo     | Requerido | Descripción       |
| :--------- | :------- | :-------: | :---------------- |
| `playlist` | `number` |    Sí     | ID de la playlist |

**Ejemplo:**

```json
{
  "playlist": 1
}
```

**Respuesta exitosa (200):**

```json
{
  "data": ["video1.mp4", "video2.mp4"],
  "message": "Playlist unificada enviada en un solo bloque con éxito",
  "video_notfound": ["missing_video.mp4"]
}
```

**Respuesta error (500) — Playlist vacía:**

```json
{
  "error": "No hay un error en ejecucion sino mas bien la playlist esta vacia",
  "message": "Agregue correctamente la lista de los videos a esta playlist por favor"
}
```

**Respuesta error (500) — Error general:**

```json
{
  "error": "<error message>",
  "message": "No se ha podido correr la playlist"
}
```

---

### 🗑️ `DELETE /playlist` — Eliminar Playlist

Elimina una playlist y todos sus videos asociados de la base de datos (transacción).

**Body (JSON):**

| Campo      | Tipo     | Requerido | Descripción       |
| :--------- | :------- | :-------: | :---------------- |
| `playlist` | `number` |    Sí     | ID de la playlist |

**Ejemplo:**

```json
{
  "playlist": 1
}
```

**Respuesta exitosa (200):**

```json
{
  "error": null,
  "message": "Se ha eliminado la playlist correctamente junto con los videos asociados a ella"
}
```

**Respuesta error (500):**

```json
{
  "error": "<db error>",
  "message": "Error al tratar de eliminar la playlist valide e intente nuevamente"
}
```

---

### ⏹️ `DELETE /stop-playlist` — Detener Reproducción

Detiene la reproducción actual en MPV enviando el comando `stop` al socket.

**Body:** No requiere body.

**Respuesta exitosa (200):**

```json
{
  "status": "Comando correctamente ejecutado"
}
```

**Respuesta error (500):**

```json
{
  "error": "write error",
  "message": "No se ha podido limpiar la playlist"
}
```

---

### 🔌 `GET /ws` — WebSocket

Endpoint WebSocket para recibir notificaciones en tiempo real desde el servidor.

**Protocolo:** WebSocket (HTTP upgrade)

**Configuración:**
- `ReadBufferSize`: 1024
- `WriteBufferSize`: 1024
- `EnableCompression`: true
- `CheckOrigin`: permite todos los orígenes

**Comportamiento:**
1. Se realiza upgrade de HTTP a WebSocket
2. El cliente se agrega al grupo de broadcast del `Hub`
3. Se envía mensaje de bienvenida: `"Alguien mas se ha conectado"`
4. El servidor escucha mensajes continuamente (no procesa entrantes)
5. Al desconectarse, se remueve del hub y se cierra la conexión

**Mensajes del servidor (texto):**

```
Alguien mas se ha conectado
Se esta reproduciendo videos
Playlist cambiada
```

**Ejemplo de conexión (JavaScript):**

```javascript
const ws = new WebSocket('ws://localhost:8000/ws')

ws.onmessage = (event) => {
  console.log('Notificación:', event.data)
}

ws.onclose = () => {
  console.log('Conexión cerrada')
}
```

---

## 🔨 Build proyecto for Linux/Windows

```
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./app/playar.exe ./cmd/server/
```

```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app/playar ./cmd/server/
```

**CGO_ENABLED** Puro codigo Go - **GOOS** Sistema operativo linux/windows - **GOARCH** Arquitextura del SO
