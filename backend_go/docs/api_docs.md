# Академическая Система Управления - Backend Implementation Guide

## Общее описание

Система управления учебным процессом с поддержкой ролей (Студент, Преподаватель, Администратор), управлением расписанием, оценками и посещаемостью.

## API Endpoints

### Аутентификация

```
POST /api/auth/login
POST /api/auth/registration
POST /api/auth/verify
GET /api/auth/ - получение информации о текущем пользователе
```

### Пользователи

```
GET /api/user/ - список всех пользователей
GET /api/user/students - список студентов
GET /api/user/teachers - список преподавателей
GET /api/user/:id - информация о пользователе
PUT /api/user/:id - обновление пользователя
DELETE /api/user/:id - удаление пользователя
```

### Группы

```
GET /api/group/ - список всех групп
GET /api/group/:id - информация о группе
POST /api/group/ - создание группы
PUT /api/group/:id - обновление группы
DELETE /api/group/:id - удаление группы
```

### Расписание

```
GET /api/schedule/ - общее расписание
GET /api/schedule/:id - расписание группы/преподавателя
POST /api/schedule/ - создание занятия
PUT /api/schedule/:id - обновление занятия
DELETE /api/schedule/:id - удаление занятия
```

### Оценки

```
GET /api/grades/student/:id - оценки студента
GET /api/grades/group/:id - оценки группы
POST /api/grades/ - выставление оценки
PUT /api/grades/:id - обновление оценки
DELETE /api/grades/:id - удаление оценки
```

### Посещаемость

```
GET /api/attendance/student/:id - посещаемость студента
GET /api/attendance/group/:id - посещаемость группы
POST /api/attendance/ - отметка посещаемости
PUT /api/attendance/:id - обновление посещаемости
```

## Структуры данных

### User

```go
type User struct {
    ID          int       `json:"id"`
    FirstName   string    `json:"first_name"`
    MiddleName  string    `json:"middle_name"`
    LastName    string    `json:"last_name"`
    RoleID      int       `json:"role_id"`
    GroupID     *int      `json:"group_id,omitempty"`
    Login       string    `json:"login"`
    Password    string    `json:"-"`
    Salt        string    `json:"-"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
```

### Group

```go
type Group struct {
    ID        int       `json:"id"`
    Name      string    `json:"name"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
```

### Schedule

```go
type Schedule struct {
    ID        int       `json:"id"`
    GroupID   int       `json:"group_id"`
    SubjectID int       `json:"subject_id"`
    TeacherID int       `json:"teacher_id"`
    DayOfWeek int       `json:"day_of_week"`
    StartTime string    `json:"start_time"`
    EndTime   string    `json:"end_time"`
    Location  string    `json:"location"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
```

### Grade

```go
type Grade struct {
    ID        int       `json:"id"`
    StudentID int       `json:"student_id"`
    SubjectID int       `json:"subject_id"`
    Value     int       `json:"value"`
    Date      time.Time `json:"date"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
```

### Attendance

```go
type Attendance struct {
    ID        int       `json:"id"`
    StudentID int       `json:"student_id"`
    SubjectID int       `json:"subject_id"`
    Date      time.Time `json:"date"`
    Status    string    `json:"status"` // present, absent, excused
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
```

## База данных

### Необходимые таблицы

```sql
CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    value VARCHAR(100) NOT NULL UNIQUE
);

CREATE TABLE groups (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    middle_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100),
    role_id INT REFERENCES roles(id),
    group_id INT REFERENCES groups(id),
    login VARCHAR(100) NOT NULL UNIQUE,
    password TEXT NOT NULL,
    salt TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE subjects (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE schedules (
    id SERIAL PRIMARY KEY,
    group_id INT REFERENCES groups(id),
    subject_id INT REFERENCES subjects(id),
    teacher_id INT REFERENCES users(id),
    day_of_week INT CHECK (day_of_week BETWEEN 1 AND 7),
    start_time TIME NOT NULL,
    end_time TIME NOT NULL,
    location VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE grades (
    id SERIAL PRIMARY KEY,
    student_id INT REFERENCES users(id),
    subject_id INT REFERENCES subjects(id),
    value INT CHECK (value BETWEEN 2 AND 5),
    date DATE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE attendance (
    id SERIAL PRIMARY KEY,
    student_id INT REFERENCES users(id),
    subject_id INT REFERENCES subjects(id),
    date DATE NOT NULL,
    status VARCHAR(20) CHECK (status IN ('present', 'absent', 'excused')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

## Безопасность

1. Использовать JWT для аутентификации
2. Хешировать пароли с солью
3. Реализовать middleware для проверки ролей
4. Добавить rate limiting для API endpoints
5. Валидировать все входные данные

## Дополнительные требования

1. Реализовать кэширование с помощью Redis
2. Добавить логирование всех действий
3. Реализовать обработку ошибок и возврат понятных сообщений
4. Добавить документацию API (например, с помощью Swagger)
5. Реализовать механизм обновления токенов