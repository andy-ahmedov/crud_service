# crud_service

 CRUD приложение для управления списком книг

 Используемый стек на данный момент:
 	golang
  	docker
   	makefile

Фреймворки:
	swagger
 	gin

 Должно быть:

	Книжный репозиторий
	Сервис книг
	Rest Сервер

Book:
	id int
	titile string
	author string
	publishDate time.Time
	rating int

UpdateBookInput:
	titile string
	author sttring
	publishDate time.Time
	rating int

Ошибка: Book not found

<pre>
Book interface {
	Create(ctx context.Context, book domain.Book) error
	GetByID(ctx context.Context, id int64) (domain.Book, error)
	GetAll(ctx context.Context) ([]domain.Book, error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, id int64, inp domain.UpdateBookInput) error
}
</pre>

Handles:
	createBook
	getAllBooks
	getBookByID
	deleteBook
	updateBook

PSQL:
	PostgresConnection

Структура проекта:

	CMD - исполняемый файл
	Internal:
		Domain - описание структур
		Repository/psql - создание структуры и ее методов для взаимодействия с базой данных
		Service - Создание интерфейса и реализация ее методов используя наработки с Repository/psql
		Transport/Rest - Инициализация маршрутизация и интеграция с интерфейсом + middleware
	PKG - подключение к драйверу Базы Данных
