package nats

// todo: there we can use db method for recover cache
// note: because all data that we give should be from cache.
// Получается, что метод из web будет использовать обращение к бд для сохранения данных в кеше,
// и nats в случае падения также должен восстанавливать данные из кэша
