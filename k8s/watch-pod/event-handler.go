package main

import (
	corev1 "k8s.io/api/core/v1"
	"log"
)

type EventHandler struct {
}

func NewEventHandler() *EventHandler {
	return &EventHandler{}
}

func (e *EventHandler) OnAdd(obj interface{}) {
	event := obj.(*corev1.Pod)
	log.Printf("OnAdd: %s", event.ObjectMeta.Name)
}

func (e *EventHandler) OnUpdate(oldObj, newObj interface{}) {
	event := newObj.(*corev1.Pod)
	log.Printf("OnUpdate: %s", event.ObjectMeta.Name)

}

func (e *EventHandler) OnDelete(obj interface{}) {
	event := obj.(*corev1.Pod)
	log.Printf("OnDelete: %s", event.ObjectMeta.Name)

}
