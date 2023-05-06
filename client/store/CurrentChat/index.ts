import { makeAutoObservable } from "mobx";

class CurrentChat {
  id: string = "";

  constructor() {
    makeAutoObservable(this);
  }

  setCurrentChat(id: string) {
    this.id = id;
  }
}

export const currentChat = new CurrentChat();
