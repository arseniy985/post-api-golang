### Description
самый обычный круд, делал в учбеных целях и для проработки архитектуры  

### Install  
cmd/app/main.go - установить нужные значения в функциях создания структур

### Architecture
api/router.go - роуты (mux routr)  
internal/handlers - функции для роутера, в которых происходит изначальная обработка, форматрование входных данных и далее обращение к сервису  
internal/services - бизнес логика  
helpers.GetBodyData - принимает указатель на структуру необходимых данных и реквест и закидывает в структуру данные из тела запроса  
database/validate.go - валидации входящих данных из запроса  
database/repo.go - функции запросов к базе данных 
