При запуске в качестве агрументов передается:
go run cmd/main.go [куда выводить логи(stdout/file)] [источник данных(stdin/file)] [вывод результата(stdout/file)]

Пример работы программы для go run cmd/main.go stdout stdin file:

Логи из stdout:

{"level":"info","msg":"TotalSum(): input - [1 2 3 4], result - 10","time":"2023-11-24 11:01:14"}
{"level":"info","msg":"GetStatusResponse() to https://ya.ru with response 200 OK","time":"2023-11-24 11:01:15"}

Результат из файла data.txt:

Array - [1 2 3 4], Sum - 10, Response Status to https://ya.ru - 200 OK

Для go run cmd/main.go file file file:

Логи из файла logs.txt:

{"level":"info","msg":"GetArrFromJson(): data - [62 60 20 28 16]","time":"2023-11-24 11:03:25"}
{"level":"info","msg":"TotalSum(): input - [62 60 20 28 16], result - 186","time":"2023-11-24 11:03:25"}
{"level":"info","msg":"GetStatusResponse() to https://ya.ru with response 200 OK","time":"2023-11-24 11:03:25"}

Результат из stdout:

Array - [62 60 20 28 16], Sum - 186, Response Status to https://ya.ru - 200 OK
