import { makeAutoObservable } from "mobx";

export class ChatBarState {
  isActive: boolean = false;

  constructor() {
    makeAutoObservable(this);
  }

  toggleIsActive() {
    this.isActive = !this.isActive;
  }
}

export const chatBarState = new ChatBarState();
