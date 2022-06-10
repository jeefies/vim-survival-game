package main

import (
	"github.com/jeefies/vim-survival/input"
	"github.com/hajimehoshi/ebiten/v2"
)

const transitionMaxCount int = 30 // Transition total frames, 60pfs / 30f = 0.5s

var (
	transitionFrom = ebiten.NewImage(screenWidth, screenHeight)
	transitionTo = ebiten.NewImage(screenWidth, screenHeight)
)

type SceneManager struct {
	current Scene
	next Scene
	transitionCount int
}

type Scene interface {
	Update(input * input.Input) error
	Draw(screen * ebiten.Image)
	StartTransition()
	EndTransition()
}

func (sm * SceneManager) Register(scene Scene) {
	if sm.current == nil {
		sm.current = scene
		return
	}

	sm.next = scene
	sm.transitionCount = transitionMaxCount
}

func (sm * SceneManager) Update(ipt * input.Input) error {
	if sm.transitionCount == 0 {
		return sm.current.Update(ipt)
	}

	sm.transitionCount--
	if sm.transitionCount > 0 {
		sm.next.Update(ipt)
		return nil
	}
	// sm.transitionCount == 0
	sm.current = sm.next
	sm.next = nil
	return nil

}

func (sm * SceneManager) Draw(screen * ebiten.Image) {
	if sm.transitionCount == 0 {
		sm.current.Draw(screen)
		return
	}

	transitionFrom.Clear()
	sm.current.Draw(transitionFrom)
	transitionTo.Clear()
	sm.next.Draw(transitionTo)

	screen.DrawImage(transitionFrom, nil)

	opacity := 1 - float64(sm.transitionCount) / float64(transitionMaxCount)
	op := &ebiten.DrawImageOptions{}
	op.ColorM.Scale(1, 1, 1, opacity)
	screen.DrawImage(transitionTo, op)
}
