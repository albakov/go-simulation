package menu

import (
	"bufio"
	"fmt"
	"github.com/albakov/go-simulation/internal/util"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

type Menu struct {
	reader     *bufio.Reader
	systemStop chan os.Signal
	stop       chan struct{}
	counter    *int
}

func New(counter *int) *Menu {
	m := &Menu{
		reader:     bufio.NewReader(os.Stdin),
		systemStop: make(chan os.Signal, 1),
		stop:       make(chan struct{}),
		counter:    counter,
	}

	signal.Notify(m.systemStop, syscall.SIGINT)

	return m
}

func (m *Menu) OnStart() {
	util.ShowMessage(util.MessageBegin)
	util.ShowMessage(util.MessagePrefix)
	command := m.commandFromInput()

	// Exit
	if command == "в" {
		util.ShowMessage(util.MessageExit, *m.counter)
		os.Exit(0)
	}

	// Start
	if command == "н" {
		return
	}

	util.ShowMessage(util.MessageIncorrectCommand)
	m.OnStart()
}

func (m *Menu) CommandListener(pause <-chan bool, paused *bool) {
	for {
		select {
		case <-m.systemStop:
			util.ShowMessage(util.MessageExit, *m.counter)
			os.Exit(0)
		case <-m.stop:
			util.ShowMessage(util.MessageExit, *m.counter)
			os.Exit(0)
		case p := <-pause:
			*paused = p

			if !p {
				util.ShowMessage(util.MessageContinue)
			}
		}
	}
}

func (m *Menu) KeyboardListener(pause chan<- bool, paused *bool) {
	for {
		command := m.commandFromInput()
		util.ShowMessage(util.MessagePrefix)

		// Pause
		if command == "п" {
			if *paused {
				pause <- false
				pause <- false
			} else {
				pause <- true
				pause <- true
			}

			continue
		}

		// Exit
		if command == "в" {
			m.stop <- struct{}{}

			continue
		}

		util.ShowMessage(util.MessageIncorrectCommand)
	}
}

func (m *Menu) commandFromInput() string {
	var command string

	_, err := fmt.Scanln(&command)
	if err != nil {
		util.ShowMessage(util.MessageIncorrectCommand)

		return ""
	}

	return strings.ToLower(command)
}
