package psql

// добавить структуру Tokens содержащую в себе поле db
// 
// NewTokens
// 

// метод Create(ctx, token domain.RefreshSession) error
// запрос Exec в бд (INSERT INTO (юзерАйди, токен и время окончания) VALUES())
// возвращаем ошибку

// метод Get(ctx, token string) (domainRefreshSession) error {
// создаем новую переменную типа domain RefreshSession
// запрос QueryRow в бд (SELECT айди, юзерАйди токен и время окончания FROM refresh_tokens WHERE токен равен $1) c последующим сканированием всех полей. Проверка на ошибки
// удаляем строку из таблицы refreshTokens где user_id равно возвращенной userID используя Exec и  DELETE FROM
// возвращаем созданную переменную и ошибку