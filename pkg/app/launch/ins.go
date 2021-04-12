package launch

import (
  "github.com/sirupsen/logrus"
  "os"
  "os/signal"
  "sync"
  "syscall"
  "time"
)

type Cfg struct {
  StartedFn func() // Функция при завершении всех задач
  Timeout   int64  // Таймаут ожидания (time.Millisecond); 0 - без таймаута
  Log       *logrus.Entry
}

type ins struct {
  startedFn  func() // Функция при завершении всех задач
  failFn     func() // Функция при ошибке в одной из задач
  exitIfFail bool   // Выйти в случае ошибки запуска
  timeout    int64  // Таймаут ожидания (time.Millisecond); 0 - без таймаута
  end        bool   // Завершение всех задач
  isErr      bool   // Одна из задач завершилась с ошибкой
  mx         sync.Mutex
  signals    chan os.Signal
  wg         sync.WaitGroup
  log        *logrus.Entry
  timer      *time.Timer
}

func New(c Cfg) *ins {
  o := &ins{
    signals:    make(chan os.Signal, 1),
    startedFn:  c.StartedFn,
    exitIfFail: true,
    timeout:    c.Timeout,
    log:        c.Log,
  }

  signal.Notify(o.signals, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

  if o.timeout > 0 {
    go func() {
      o.timer = time.AfterFunc(time.Millisecond*time.Duration(o.timeout), o.Exit)
    }()
  }

  return o
}

func (i *ins) Exit() {
  i.signals <- nil
}

func (i *ins) Timeout(t int64) {
  i.timeout = t
}

func (i *ins) ExitIfFail(on bool) {
  i.exitIfFail = on
}

func (i *ins) StartedFn(fn func()) {
  i.startedFn = fn
}

func (i *ins) FailFn(fn func()) {
  i.failFn = fn
}

// Task Добавляет задачу на запуск
func (i *ins) Task(fn func() error) {
  if i.end {
    i.log.Panic("Chan cannot be added after Wait")
  }

  go func() {
    i.wg.Add(1)

    if err := fn(); err != nil {
      i.isErr = true
      i.log.Errorf("Ошибка запуска задачи: %v", err)
    }

    i.wg.Done()
  }()
}

// Wait ожидание запуска задач
func (i *ins) Wait() {
  i.end = true

  go func() {
    i.wg.Wait()

    if i.timer != nil {
      i.timer.Stop()
      i.timer = nil
    }

    if i.isErr {
      if i.failFn != nil {
        i.failFn()
      }
      if i.exitIfFail {
        i.Exit()
      }
    } else if !i.isErr && i.startedFn != nil {
      i.startedFn()
    }
  }()

  <-i.signals
}
