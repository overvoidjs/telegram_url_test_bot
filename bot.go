package main

import (
  "bufio"
  "net/http"
  "os"
  "log"
  "strings"
  "strconv"
  "time"
)

// Читаем env и возвращаем его значение
func getEnv(key string) string {
  file, err := os.Open(".env")
  if err != nil {
      log.Fatal("Не удалось загрузить .env файл")
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
      envMass := strings.Split(scanner.Text(),"=")
      os.Setenv(envMass[0],envMass[1])
  }

  return os.Getenv(key)
}

//Получаем список url
func getList() []string {

  var testList []string

  file, err := os.Open(".testlist")
  if err != nil {
      log.Fatal("Не удалось загрузить .testlist файл")
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    testList = append(testList, scanner.Text())
  }

  return testList
}

// отправляем сообщение в телеграмм
func sendMessage(msg string) {
  botId := getEnv("BOT_ID")
  chatId := getEnv("CHAT_ID")

  resp, err := http.Get("https://api.telegram.org/bot"+botId+"/sendMessage?chat_id="+chatId+"&text="+msg+"")
  if err != nil || resp.StatusCode != http.StatusOK {
    log.Fatal("Не удалось передать сообщение в Telegram")
  }

}

//Запрашиваем код ответа
func testUrl(url string) int {
  resp, err := http.Get(url)
  if err != nil {
    log.Fatal("Не удалось открыть указаный URL: "+url+"")
  }

  return resp.StatusCode
}

//Основная функция
func run(){
  tests := getList()
  timeout, _ := strconv.Atoi(getEnv("TIMEOUT"))

  for _, url := range tests {
    //Получаем код проверки
    respTest := testUrl(url)
    //Если не 200 - сообщаем
    if respTest != 200 {
      sendMessage(url+" отвечает "+strconv.Itoa(respTest))
    }
    //Поспим одну секунду что бы не грузить сервер
    time.Sleep(time.Duration(timeout) * time.Second)
  }

  //Уведомление о проверке
  t := time.Now().Format("2006-01-02 15:04:05")
  sendMessage("Проверка доступности от "+t+" завершена")
}

func main() {
  log.Print("Бот запущен")
  for {
    run()
    looptime, _ := strconv.Atoi(getEnv("LOOP_TIME"))
    time.Sleep(time.Duration(looptime) * time.Minute)
  }

}
