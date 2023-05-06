import { makeAutoObservable } from "mobx";

import { CurrentChatProps } from "./models";

class CurrentChat {
  id: string = "";
  companionAvatar: string = "";
  companionName: string = "";

  constructor() {
    makeAutoObservable(this);
  }

  getSelf() {
    return this;
  }

  setCurrentChat({ id, companionAvatar, companionName }: CurrentChatProps) {
    this.id = id;
    this.companionAvatar = companionAvatar;
    this.companionName = companionName;
  }
}

export const currentChat = new CurrentChat();
