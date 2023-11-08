
## API Reference

#### Create new event

```http
  POST /api/event
```

| Header | Description                           |
| :----- | :------------------------------------ |
| `token` | **Required**. Your API key |

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `id` | `uint` | **Required**. IDK Yet |
| `class_id` | `uint` | **Required**. Class ID intended for attendance  |
| `meet` | `uint` | **Required**. Indicates which course meetings |
| `student_id` | `array[uint]` | List of student IDs that have been recorded |

#### Record attendance

```http
  GET /api/event/${id}/record
```

| Header | Description                           |
| :----- | :------------------------------------ |
| `token` | **Required**. Your API key |

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `nim`      | `string` | **Required**. ID of the student whose attendance will be recorded |

#### Record mass attendance

```http
  GET /api/event/${id}/massrecord
```

| Header | Description                           |
| :----- | :------------------------------------ |
| `token` | **Required**. Your API key |

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `nim`      | `array[string]` | **Required**. List of IDs of the student whose attendance will be recorded |

#### Update event

```http
  PATCH/PUT /api/event/${id}
```

| Header | Description                           |
| :----- | :------------------------------------ |
| `token` | **Required**. Your API key |

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `status`      | `uint` | **Required**. Attendance event status |