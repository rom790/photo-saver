# Makefile

# Имя исполняемого файла
BINARY_NAME=app

# Путь к основному файлу
MAIN_FILE=main.go

# Цель сборки
build:
	go build -o $(BINARY_NAME).exe $(MAIN_FILE)

# Цель запуска
run: build
	./$(BINARY_NAME)

# Цель очистки
clean:
	go clean
	rm -f $(BINARY_NAME)

.PHONY: build run clean
