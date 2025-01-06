package core

func RegisterUser(db *sql.DB, user models.User) (int, error) {
    var userID int

    // Проверка уникальности логина
    var existingID int
    err := db.QueryRow("SELECT id FROM users WHERE login = $1", user.Login).Scan(&existingID)
    if err != sql.ErrNoRows {
        return 0, fmt.Errorf("login already exists")
    }

    // Вставка нового пользователя
    err = db.QueryRow(
        "INSERT INTO users (first_name, middle_name, last_name, role_id, group_id, login, password, salt, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, NOW(), NOW()) RETURNING id",
        user.FirstName, user.MiddleName, user.LastName, user.RoleID, user.GroupID, user.Login, user.Password, user.Salt,
    ).Scan(&userID)
    if err != nil {
        return 0, err
    }

    return userID, nil
}

func AuthenticateUser(db *sql.DB, login, password string) (*models.User, error) {
    var user models.User

    // Получение данных пользователя
    err := db.QueryRow(
        "SELECT id, password, salt, role_id FROM users WHERE login = $1",
        login,
    ).Scan(&user.ID, &user.Password, &user.Salt, &user.RoleID)
    if err != nil {
        return nil, fmt.Errorf("invalid login or password")
    }

    // Проверка пароля
    if !CheckPasswordHash(password, user.Password, user.Salt) {
        return nil, fmt.Errorf("invalid login or password")
    }

    return &user, nil
}

func VerifyToken(db *sql.DB, token string) (*models.User, error) {
    claims, err := utils.ValidateToken(token)
    if err != nil {
        return nil, err
    }

    var user models.User
    err = db.QueryRow(
        "SELECT id, role_id FROM users WHERE id = $1",
        claims.UserID,
    ).Scan(&user.ID, &user.RoleID)
    if err != nil {
        return nil, fmt.Errorf("user not found")
    }

    return &user, nil
}

func GetCurrentUser(db *sql.DB, userID int) (*models.User, error) {
    var user models.User

    err := db.QueryRow(
        "SELECT id, first_name, middle_name, last_name, role_id, group_id, login, created_at, updated_at FROM users WHERE id = $1",
        userID,
    ).Scan(&user.ID, &user.FirstName, &user.MiddleName, &user.LastName, &user.RoleID, &user.GroupID, &user.Login, &user.CreatedAt, &user.UpdatedAt)
    if err != nil {
        return nil, err
    }

    return &user, nil
}